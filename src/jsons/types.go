package jsons

import (
	"encoding/json"
	"time"
	"encoding/xml"
)

type Response struct {
	Result Result `json:"result,omitempty"`
}

type Result struct {
	TotalRows int `json:"totalrows,omitempty"`
	RowList   []Rows `json:"rows,omitempty"`
}

type Rows struct {
	RowNum int `json:"rownum,omitempty"`
	Values []Values `json:"values,omitempty"`
}

type Values struct {
	Field string `json:"field,omitempty"`
	Value json.Number `json:"value,omitempty"`
}

type Transaction struct {
	FilerName string `json:"filerName,omitempty"`
	TransactionDate time.Time `json:"transactionDate,omitempty"`
	TransactionType string `json:"transactionType,omitempty"`
	OwnershipType string `json:"ownershipType,omitempty"`
	Relationship string `json:"relationship, omitempty"`
	IssueTicker string `json:"issueTicker,omitempty"`
	IssueType string `json:"issueType,omitempty"`
	IssueId int `json:"issueId,omitempty"`
	FilerId int `json:"filerId,omitempty"`
	FilingId int `json:"filingId,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type RssEnclosure struct {
	//RSS 2.0 <enclosure url="http://example.com/file.mp3" length="123456789" type="audio/mpeg" />
	XMLName xml.Name `xml:"enclosure"`
	Url     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

type RssFeed struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`       // required
	Link           string   `xml:"link"`        // required
	Description    string   `xml:"description"` // required
}