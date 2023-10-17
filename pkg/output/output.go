package output

const (
	TypeApply = "apply"
	TypeFile  = "file"
)

// Output represents a sealed secret output
type Output struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}
