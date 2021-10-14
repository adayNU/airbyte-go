package source

import (
	"context"

	"github.com/adayNU/airbyte-go/types"
)

type Source interface {
	Spec() *types.ConnectorSpecification
	Check(types.JSONData) *types.AirbyteConnectionStatus
	Discover(types.JSONData) *types.AirbyteCatalog
	// Read reads data from the underlying data source and converts
	// it into AirbyteRecordMessage (wrapped in an AirbyteMessage).
	// It can optionally return AirbyteStateMessages, which is used to
	// track how much of the data source has been synced.
	//
	// Per the Airbyte documentation: The connector ideally will only pull
	// the data described in the catalog argument. It is permissible for
	// the connector, however, to ignore the catalog and pull data from any
	// stream it can find. If it follows this second behavior, the extra
	// data will be pruned in the worker. We prefer the former behavior
	// because it reduces the amount of data that is transferred and allows
	// control over not sending sensitive data. There are some sources for
	// which this is simply not possible.
	// 	https://docs.airbyte.io/understanding-airbyte/airbyte-specification#source
	//
	// It is Read's responsibility to close the returned channel to signal to
	// there is no more data to read. Read should close the channel and return
	// on a canceled context.
	Read(context.Context, types.JSONData, *types.ConfiguredAirbyteCatalog, types.JSONData) <-chan types.AirbyteMessage
}
