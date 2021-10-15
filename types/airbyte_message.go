package types

type MessageType int

const (
	Record MessageType = iota
	State
	Log
	Spec
	ConnectionStatus
	Catalog
)

type AirbyteMessage struct {
	Type             MessageType
	Spec             *ConnectorSpecification  `json:"spec,omitempty"`
	ConnectionStatus *AirbyteConnectionStatus `json:"connectionStatus,omitempty"`
	Catalog          *AirbyteCatalog          `json:"catalog,omitempty"`
	Record           *AirbyteRecordMessage    `json:"record,omitempty"`
	State            *AirbyteStateMessage     `json:"state,omitempty"`
}
