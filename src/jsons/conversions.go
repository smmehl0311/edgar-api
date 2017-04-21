package jsons

import (
	"time"
	"fmt"
)

func ToTransactions(rows []Rows) []Transaction {
	transactions := make([]Transaction, len(rows))
	for i, row := range rows {
		transactions[i] = BuildTransaction(row)
	}
	return transactions
}

func BuildTransaction(row Rows) Transaction {
	filerName := ""
	const layout = "1/2/2006"
	var transactionDate time.Time
	transactionType := ""
	ownershipType := ""
	relationship := ""
	issueTicker := ""
	issueType := ""
	var issueId int64 = 0
	var filerId int64 = 0
	var filingId int64 = 0
	price := 0.0
	for _, values := range row.Values {
		switch values.Field {
		case "filername": filerName = string(values.Value)
		case "transactiondate": transactionDate, _ = time.Parse(layout, string(values.Value))
		case "transactiontype": transactionType = string(values.Value)
		case "ownershiptype": ownershipType = string(values.Value)
		case "relationship": relationship = string(values.Value)
		case "issueticker": issueTicker = string(values.Value)
		case "issuetype": issueType = string(values.Value)
		case "issueid": issueId, _ = values.Value.Int64()
		case "filerid": filerId, _ = values.Value.Int64()
		case "filingid": filingId, _ = values.Value.Int64()
		case "price": price, _ = values.Value.Float64()
		}
	}
	fmt.Println(transactionDate)
	return Transaction{filerName, transactionDate,
				 transactionType, ownershipType, relationship,
				 issueTicker, issueType, int(issueId), int(filerId), int(filingId), float64(price)}
}
