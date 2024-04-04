package account

var (
	// represents when the document searching was not found in memory
	ErrDocumentNotFound Error = "document not found"
)

// implements error interface and must be used to represents a business violation.
type Error string

// Error returns the violation code as string
func (e Error) Error() string {
	return string(e)
}

type AccountError struct {
	Description string `json:"description"`
}
