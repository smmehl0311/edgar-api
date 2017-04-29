package jsons

import (
	"fmt"
	"io"
	"encoding/xml"
	"utils"
	"golang.org/x/text/encoding/charmap"
)

func DecodeOwnershipDoc(body io.ReadCloser) OwnershipDocument {
	decoder := xml.NewDecoder(body)
	decoder.CharsetReader = makeCharsetReader
	ownershipDocument := OwnershipDocument{}
	err := decoder.Decode(&ownershipDocument)
	utils.HandleError(err)
	return ownershipDocument
}

func DecodeRecentFilings(body io.ReadCloser) RecentFilings {
	decoder := xml.NewDecoder(body)
	decoder.CharsetReader = makeCharsetReader
	recentFilings := RecentFilings{}
	err := decoder.Decode(&recentFilings)
	utils.HandleError(err)
	return recentFilings
}

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "ISO-8859-1" {
		// Windows-1252 is a superset of ISO-8859-1, so should do here
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("Unknown charset: %s", charset)
}