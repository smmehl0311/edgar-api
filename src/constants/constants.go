package constants

import (
	//"net/url"
)

const TimeLayout = "2006-01-02T15:04:05-4:00"

const MongoUrl = "mongodb://localhost"
const MongoDb = "edgar"
const TransactionsCollection = "transactions"
const LastUpdated = "last_updated"
const Form4 = "recent_form_4"

const BaseUrl = "https://www.sec.gov"
const RecentFilingsUrl = BaseUrl + "/cgi-bin/browse-edgar?action=getcurrent&CIK=&type=4&company=&dateb=&owner=include&output=atom"
