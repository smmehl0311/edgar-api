package constants

import (
	"net/url"
)

const MongoUrl = "mongodb://localhost"
const MongoDb = "edgar"
const TransactionsCollection = "transactions"

const debug = "false"
const isDeleted = "false"
const sortBy = "transactiondate+asc"
const baseUrl = "http://edgaronline.api.mashery.com/v2"
const insiderApiKey = "wjyrpr284a4ue8ajy4qq6zj7"

var insiderSummaryFields = url.PathEscape("issueid,transactiondate,transactionpricefrom,ownershiptype,transactiontype,sharesbought,numbuys,sharessole,numsells,grosssharesm,netshares,numtransactions")
var InsiderSummaryUrl = baseUrl+"/insiders/summary?appkey="+insiderApiKey+"&debug="+debug+"&fields="+insiderSummaryFields

var insiderTransactionFields = url.PathEscape("issueid,filerid,filername,transactiondate,transactionpricefrom,transactiontype,ownershiptype,insiderformtype,relationship,issuename,issueticker,pricedate,price,issuetype,issueid,filingid")
var InsiderTransactionsUrl = baseUrl+"/insiders/transactions?appkey="+insiderApiKey+"&debug="+debug+"&deleted="+isDeleted+"&sortby="+sortBy+"&fields="+insiderTransactionFields