package source

import "github.com/adayNU/airbyte-go/types"

type Source interface {
	Spec() *types.ConnectorSpecification
	Check(types.JSONData) *types.AirbyteConnectionStatus
	Discover(types.JSONData) *types.AirbyteCatalog
	Read(types.JSONData, *types.ConfiguredAirbyteCatalog, types.JSONData) <-chan types.AirbyteMessage
}
