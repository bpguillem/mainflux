package senml

const (
	// JSON represents SenML in JSON format content type.
	JSON = "application/senml+json"

	// CBOR represents SenML in CBOR format content type.
	CBOR = "application/senml+cbor"
)

// Message represents a resolved (normalized) SenML record.
type Message struct {
	Channel     string   `json:"channel,omitempty"`
	Topic	    string   `json:"topic,omitempty"`
	Subtopic    string   `json:"subtopic,omitempty"`
	Publisher   string   `json:"publisher,omitempty"`
	Protocol    string   `json:"protocol,omitempty"`
	Name        string   `json:"name,omitempty"`
	Unit        string   `json:"unit,omitempty"`
	Time        float64  `json:"time,omitempty"`
	UpdateTime  float64  `json:"update_time,omitempty"`
	Value       *float64 `json:"value,omitempty"`
	StringValue *string  `json:"string_value,omitempty"`
	DataValue   *string  `json:"data_value,omitempty"`
	BoolValue   *bool    `json:"bool_value,omitempty"`
	Sum         *float64 `json:"sum,omitempty"`
}
