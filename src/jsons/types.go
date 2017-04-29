package jsons

import (
	"time"
	"encoding/xml"
)

type RecentFilings struct {
	XMLName xml.Name	`xml:"feed"`
	Title   string		`xml:"title"`
	Link    Link		`xml:"link"`
	Id      string		`xml:"id"`
	Author  Author		`xml:"author"`
	Updated string		`xml:"updated"`
	Entries []Entry		`xml:"entry"`
}

type LastUpdated struct {
	Title 		string		`xml,json:"title"`
	LastUpdated	time.Time	`xml,json:"updated"`
}

type Link struct {
	Href		string		`xml:"href,attr"`
}

type Author struct {
	Name		string		`xml:"name"`
	Email		string		`xml:"email"`
}

type Entry struct {
	Title		string		`xml:"title"`
	Link		Link		`xml:"link"`
	Summary		string		`xml:"summary"`
	Updated		string		`xml:"updated"`
	Category	string		`xml:"category"`
	Id			string		`xml:"id"`
}

type OwnershipDocument struct {
	Version			string		`xml:"schemaVersion"`
	Type			string		`xml:"documentType"`
	Period			string		`xml:"periodOfReport"`
	Something 		string 		`xml:"notSubjectToSection16"`
	Issuer 			Issuer 		`xml:"issuer"`
	Owner			[]Owner		`xml:"reportingOwner"`
	NonDerivativeTable NonDerivativeTable	`xml:"nonDerivativeTable"`
	DerivativeTable	DerivativeTable	`xml:"derivativeTable"`
	Footnotes 		Footnotes	`xml:"footnotes"`
	Signature		Signature	`xml:"ownerSignature"`
}

type Issuer struct {
	CIK		int		`xml:"issuerCik"`
	Name	string	`xml:"issuerName"`
	Ticker	string	`xml:"issuerTradingSymbol"`
}

type Owner struct {
	OwnerId 			OwnerId				`xml:"reportingOwnerId"`
	OwnerRelationship	OwnerRelationship	`xml:"reportingOwnerRelationship"`
}

type OwnerId struct {
	CIK		int		`xml:"rptOwnerCik"`
	Name	string	`xml:"rptOwnerName"`
}

type OwnerRelationship struct {
	IsDirector			string 		`xml:"isDirector"`
	IsOfficer			string		`xml:"isOfficer"`
	IsTenPercentOwner	string		`xml:"isTenPercentOwner"`
	IsOther				string		`xml:"isOther"`
	Title				string		`xml:"officerTitle"`
}

type NonDerivativeTable struct {
	NonDerivativeTransactions	[]NonDerivativeTransaction 		`xml:"nonDerivativeTransaction"`
}

type NonDerivativeTransaction struct {
	Title  KeyValue	`xml:"securityTitle"`
	Date   KeyValue	`xml:"transcationDate"`
	Coding Coding	`xml:"transactionCoding"`
	Amounts	Amounts	`xml:"transactionAmounts"`
	PostAmounts PostAmounts `xml:"postTransactionAmounts"`
	OwnershipNature	OwnershipNature	`xml:"ownershipNature"`
}

type DerivativeTable struct {
	DerivativeTransactions	[]DerivativeTransaction		`xml:"derivativeTransaction"`
}

type DerivativeTransaction struct {
	Title				KeyValue 				`xml:"securityTitle"`
	Price				KeyValue				`xml:"conversionOrExercisePrice"`
	ExerciseDate		KeyValue 				`xml:"exerciseDate"`
	TransactionDate		KeyValue				`xml:"transactionDate"`
	ExpirationDate		KeyValue				`xml:"expirationDate"`
	Coding				Coding					`xml:"transactionCoding"`
	Amounts				Amounts					`xml:"transactionShares"`
	UnderlyingSecurity	UnderlyingSecurity		`xml:"underlyingSecurity"`
	PostTransactionAmts PostTransactionAmts		`xml:"postTransactionAmounts"`
	OwnershipNature		OwnershipNature			`xml:"ownershipNature"`
	FootNotes			[]Footnotes				`xml:"footnotes"`
}

type Coding struct {
	Type 		int		`xml:"transactionFormType"`
	Code		string	`xml:"transactionCode"`
	EquitySwap 	string	`xml:"equitySwapInvolved"`
}

type Amounts struct {
	Shares 		KeyValue 	`xml:"transactionShares"`
	PPS			KeyValue	`xml:"transactionPricePerShare"`
	Code		KeyValue	`xml:"transactionAcquiredDisposedCode"`
}

type PostAmounts struct {
	SharesOwned		KeyValue	`xml:"sharesOwnedFollowingTransaction"`
}

type OwnershipNature struct {
	DirectOrIndirect	KeyValue	`xml:"directOrIndirectOwnership"`
}

type Footnotes struct {
	Footnote	string	`xml:"footnote"`
}

type UnderlyingSecurity struct {
	Title	KeyValue		`xml:"underlyingSecurityTitle"`
	Shares	KeyValue		`xml:"underlyingSecurityShares"`
}

type PostTransactionAmts struct {
	SharesOwned		KeyValue	`xml:"sharesOwnedFollowingTransaction"`
}

type Signature	struct {
	Name	string	`xml:"signatureName"`
	Date	string	`xml:"signatureDate"`
}

type KeyValue struct {
	Value	string	`xml:"value"`
}