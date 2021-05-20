package sapoci

type CompleteTransmissionIndicator string
type SustainableCompliance string
type SustainableIndicatorTypeCode string
type DeletionIndicator string

const (
	DeltaCatalog CompleteTransmissionIndicator = ""
	FullCatalog  CompleteTransmissionIndicator = "X"

	FullyCompliant     SustainableCompliance = "01"
	PartiallyCompliant SustainableCompliance = "02"
	NotCompliant       SustainableCompliance = "02"

	DirectEmission    SustainableIndicatorTypeCode = "01"
	IndirectEmission  SustainableIndicatorTypeCode = "02"
	TemporaryEmission SustainableIndicatorTypeCode = "03"
	RecycleEmission   SustainableIndicatorTypeCode = "04"
	WaterConsumption  SustainableIndicatorTypeCode = "05"
	EnergyConsumption SustainableIndicatorTypeCode = "06"
	WasteGenerated    SustainableIndicatorTypeCode = "07"

	ItemIsDeleted DeletionIndicator = "X"
)
