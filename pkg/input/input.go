package input

const (
	TypeFile = "file"
)

// Input represents a secret input
type Input struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}
