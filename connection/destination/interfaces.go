package destination

import "github.com/adayNU/airbyte-go/types"

type Destination interface {
	Spec() *types.ConnectorSpecification
	Check(types.JSONData) *types.AirbyteConnectionStatus
	Write(types.JSONData, *types.ConfiguredAirbyteCatalog, <-chan *types.AirbyteMessage, chan<- bool)
}
