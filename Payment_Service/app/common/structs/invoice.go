package commonStructs

type InvoiceStatus string

const (
	Success InvoiceStatus = "SUCCESS"
	Ongoing InvoiceStatus = "ONGOING"
	Failed  InvoiceStatus = "FAILED"
)
