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
	Spec             *ConnectorSpecification
	ConnectionStatus *AirbyteConnectionStatus
	Catalog          *AirbyteCatalog
	Record           *AirbyteRecordMessage
	State            *AirbyteStateMessage
}
