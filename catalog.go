package sapoci

type Catalog struct {
	CurrentPage                   int    `json:"CURRENTPAGE"`
	TotalPages                    int    `json:"TOTALPAGES"`
	TotalItems                    int    `json:"TOTALITEMS"`
	TransactionID                 int    `json:"TRANSACTIONID"`
	CompleteTransmissionIndicator string `json:"CTI,omitempty"`
	Items                         []Item `json:"Items"`
}
