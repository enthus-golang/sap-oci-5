package sapoci

type Item struct {
	NewItem *NewItem `json:"NEWITEM"`
}

type NewItem struct {
	ExternalProductID      string        `json:"EXT_PRODUCT_ID" validate:"required,max=40"`
	CatalogName            string        `json:"CATALOG_NAME,omitempty" validate:"max=255"`
	VendorName             string        `json:"VENDOR_NAME,omitempty" validate:"max=255"`
	Vendor                 string        `json:"VENDOR,omitempty" validate:"max=10"`
	VendorMaterialNumber   string        `json:"VENDOR_MAT,omitempty" validate:"max=40"`
	ItemType               string        `json:"ITEM_TYPE,omitempty" validate:"max=1"`
	Languages              []Language    `json:"LANGUAGES" validate:"max=required,min=1"`
	Description            []Description `json:"DESCRIPTION,omitempty"`
	Service                string        `json:"SERVICE,omitempty" validate:"max=1"`
	PriceValidFrom         *Time         `json:"PRICE_VALID_FROM,omitempty"`
	PriceValidTo           *Time         `json:"PRICE_VALID_TO,omitempty"`
	Price                  float64       `json:"PRICE,omitempty"`
	PriceQuantity          string        `json:"PRICE_QUANTITY,omitempty" validate:"max=15"`
	Currency               string        `json:"CURRENCY,omitempty" validate:"max=5"`
	PriceScales            []PriceScale  `json:"PRICE_SCALES,omitempty"`
	Tax                    float64       `json:"TAX,omitempty"`
	Longtext               []Longtext    `json:"LONGTEXT,omitempty"`
	SearchTerm             []SearchTerm  `json:"SEARCH_TERM,omitempty"`
	MaterialNumber         string        `json:"MATNR,omitempty" validate:"max=40"`
	MaterialGroup          string        `json:"MATGROUP,omitempty" validate:"max=20"`
	ManufacturerName       string        `json:"MANUFACTNAME,omitempty" validate:"max=255"`
	ManufacturerPartNumber string        `json:"MANUFACTMAT,omitempty" validate:"max=255"`
	ManufacturerCode       string        `json:"MANUFACTCODE,omitempty" validate:"max=10"`
	LeadTime               string        `json:"LEADTIME,omitempty" validate:"max=5"`
	Unit                   string        `json:"UNIT,omitempty" validate:"max=3"`
	ContentUnit            string        `json:"CONTENT_UNIT,omitempty" validate:"max=3"`
	PackagingQuantity      string        `json:"PACKAGING_QTY,omitempty" validate:"max=15"`
	MinimumOrderQuantity   string        `json:"MINORDER_QTY,omitempty" validate:"max=15"`
	LOTSize                string        `json:"LOTSIZE,omitempty" validate:"max=15"`
	Contract               string        `json:"CONTRACT,omitempty" validate:"max=10"`
	ContractItem           string        `json:"CONTRACT_ITEM,omitempty" validate:"max=19"`
	ExternalQuoteID        string        `json:"EXT_QUOTE_ID,omitempty" validate:"max=35"`
	ExternalQuoteItem      string        `json:"EXT_QUOTE_ITEM,omitempty" validate:"max=10"`
	ExternalSchemaType     string        `json:"EXT_SCHEMA_TYPE,omitempty" validate:"max=10"`
	ExternalCategoryID     string        `json:"EXT_CATEGORY_ID,omitempty" validate:"max=60"`
	ExternalCategory       string        `json:"EXT_CATEGORY,omitempty" validate:"max=40"`
	//Attributes     []Attribute        `json:"ATTRIBUTES"`
	AverageRating         float64                `json:"AVERAGE_RATING,omitempty"`
	AssortmentIndicator   string                 `json:"ASSORTMENT_IND,omitempty" validate:"max=2"`
	SustainableCompliance SustainableCompliance  `json:"SUSTAINABLE_CMPL,omitempty" validate:"max=2"`
	SustainableIndicators []SustainableIndicator `json:"SUSTAINABLE_IND,omitempty"`
	ParentID              string                 `json:"PARENT_ID,omitempty" validate:"max=255"`
	//Attachments            []Attachment  `json:"ATTACHMENTS"`
	DeletionIndicator DeletionIndicator `json:"DELETION_IND,omitempty" validate:"max=1"`
	CatalogManaged    string            `json:"CATALOG_MANAGED,omitempty" validate:"max=1"`
	ChangedAt         *Time             `json:"CHANGED_AT,omitempty"`
	SLDSystemName     string            `json:"SLD_SYS_NAME,omitempty" validate:"max=60"`
	CustomerFields    []CustomerField   `json:"CUSTOMER_FIELDS,omitempty"`
}

type Language struct {
	LanguageCode string `json:"languageCode" validate:"len=2"`
}

type Description struct {
	LanguageCode string `json:"languageCode" validate:"len=2"`
	Description  string `json:"description" validate:"required,max=255"`
}

type PriceScale struct {
	ScaleType string  `json:"scaleType" validate:"len=2"`
	Low       string  `json:"low" validate:"max=15"`
	Price     float64 `json:"price"`
}

type Longtext struct {
	LanguageCode string `json:"languageCode" validate:"len=2"`
	Longtext     string `json:"longtext"`
}

type SearchTerm struct {
	LanguageCode string `json:"languageCode" validate:"len=2"`
	Keyword      string `json:"keyword" validate:"required,max=255"`
}

type CustomerField struct {
	FieldName    string `json:"fieldName" validate:"max=100"`
	FieldValue   string `json:"fieldValue" validate:"max=100"`
	LanguageCode string `json:"languageCode" validate:"len=2"`
}

type SustainableIndicator struct {
	TypeCode SustainableIndicatorTypeCode `json:"typeCode"`
	Value    string                       `json:"value"`
	Unit     string                       `json:"unit"`
}
