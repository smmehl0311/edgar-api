package main

import (
	"net/http"
	"jsons"
	"fmt"
	"encoding/xml"
	"io"
	"os"
	"strings"
	"golang.org/x/net/html"
	"db"
	"time"
	"github.com/gorilla/mux"
	"routes"
	"utils"
	"constants"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/token", routes.GetToken).Methods("GET")
	router.HandleFunc("/last_updated", routes.GetLastUpdated).Methods("GET")
	router.HandleFunc("/recent", routes.GetRecentFilings).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))

	resp := getResponse(constants.RecentFilingsUrl + "&start=0&count=100")
	fmt.Println("Got response from: %v", constants.RecentFilingsUrl)
	defer resp.Body.Close()
	recentFilings := jsons.DecodeRecentFilings(resp.Body)

	ownershipDocs := getOwnershipDocs(recentFilings)
	db.InsertOwnershipDocuments(db.GetSession(), constants.Form4, ownershipDocs)

	lastUpdated := createLastUpdated(recentFilings)
	db.InsertLastUpdated(db.GetSession(), "last_updated", lastUpdated)
}

func getOwnershipDocs(recentFilings jsons.RecentFilings) []jsons.OwnershipDocument {
	var ownershipDocuments []jsons.OwnershipDocument
	start := 0
	count := 100
	lastUpdated := db.GetLastUpdated(db.GetSession())
	fmt.Println(lastUpdated)
	for i, entry := range recentFilings.Entries {
		if isNewEntry(entry, lastUpdated) {
			newList := append(ownershipDocuments, getOwnershipDoc(entry))
			ownershipDocuments = newList
		} else {
			fmt.Println("No more new entries")
			break
		}
		if i == 99 {
			fmt.Println("NEW PAGE")
			start += 100
			count += 100
			params := fmt.Sprintf("&start=%s&count=%s", start, count)
			resp := getResponse(constants.RecentFilingsUrl + params)
			recentFilings = jsons.DecodeRecentFilings(resp.Body)
			newList := append(ownershipDocuments, getOwnershipDocs(recentFilings)...)
			ownershipDocuments = newList
			resp.Body.Close()
		}
	}
	return ownershipDocuments
}

func isNewEntry(entry jsons.Entry, lastUpdated time.Time) bool {
	entryTime, err := time.Parse(constants.TimeLayout, entry.Updated)
	utils.HandleError(err)
	fmt.Println(entryTime.Sub(lastUpdated))
	sub := entryTime.Sub(lastUpdated)
	oneDayAgo := time.Now().Sub(time.Now().Add(-24*time.Hour))
	return sub > 0 && sub <= oneDayAgo
}

func getOwnershipDoc(entry jsons.Entry) jsons.OwnershipDocument {
	resp1 := getResponse(entry.Link.Href)
	ownershipDocumentUrl := constants.BaseUrl + getXmlLink(resp1.Body)
	resp2 := getResponse(ownershipDocumentUrl)
	ownershipDocument := jsons.DecodeOwnershipDoc(resp2.Body)
	//logOwnershipDoc(ownershipDocument)
	defer resp1.Body.Close()
	defer resp2.Body.Close()
	return ownershipDocument
}

func logOwnershipDoc(doc jsons.OwnershipDocument) {
	pprint, err := xml.MarshalIndent(doc, "  ", "  ")
	utils.HandleError(err)
	os.Stdout.Write(pprint)
}

func getResponse(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	utils.HandleError(err)
	return resp
}

func getXmlLink(body io.ReadCloser) string {
	xmlLink := ""
	z := html.NewTokenizer(body)
	for i := 0; i < 2; {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		} else if tt == html.StartTagToken {
			_, v, _ := z.TagAttr()
			if strings.Contains(string(v), ".xml") {
				fmt.Println(string(v))
				if i == 1 {
					xmlLink = string(v)
				}
				i++
			}
		}
	}
	fmt.Printf("xmlLink: %v", xmlLink)
	return xmlLink
}

func createLastUpdated(recentFilings jsons.RecentFilings) jsons.LastUpdated {
	var lastUpdated jsons.LastUpdated
	lastUpdated.Title = recentFilings.Title
	t, err := time.Parse(constants.TimeLayout, recentFilings.Updated) //timezone is EDT
	utils.HandleError(err)
	lastUpdated.LastUpdated = t
	return lastUpdated
}