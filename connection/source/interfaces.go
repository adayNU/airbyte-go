package source

import (
	"context"

	"github.com/adayNU/airbyte-go/types"
)

// readResponse receives messages from the Read operation.
type readResponse struct {
	recordMessages <-chan *types.AirbyteRecordMessage
	stateMessages  <-chan *types.AirbyteStateMessage
	done           <-chan struct{}
	err            error
}

func (r *readResponse) Records() <-chan *types.AirbyteRecordMessage {
	return r.recordMessages
}

func (r *readResponse) States() <-chan *types.AirbyteStateMessage {
	return r.stateMessages
}

func (r *readResponse) Done() <-chan struct{} {
	return r.done
}

func (r *readResponse) Err() error {
	return r.err
}

type ReadResponse interface {
	Records() <-chan *types.AirbyteRecordMessage
	States() <-chan *types.AirbyteStateMessage
	Done() <-chan struct{}
	Err() error
}

type Source interface {
	Spec() *types.ConnectorSpecification
	Check(config types.JSONData) *types.AirbyteConnectionStatus
	Discover(config types.JSONData) *types.AirbyteCatalog
	// Read reads data from the underlying data source and converts
	// it into AirbyteRecordMessage. It can optionally return
	// AirbyteStateMessages, which is used to track how much of the
	// data source has been synced.
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
	// It is Read's responsibility to send a message on the done channel to signify
	// it has completed it's work. If Read encounters an error, it should set the
	// err field and send on the done channel. A done message should also be sent
	// on a canceled context / exceeded deadline (with err set to context.Canceled or
	// context.DeadlineExceeded).
	Read(ctx context.Context, config types.JSONData, catalog *types.ConfiguredAirbyteCatalog, state types.JSONData) ReadResponse
}
