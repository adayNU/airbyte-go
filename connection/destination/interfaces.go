package destination

import (
	"context"

	"github.com/adayNU/airbyte-go/types"
)

type Destination interface {
	Spec() *types.ConnectorSpecification
	Check(config types.JSONData) *types.AirbyteConnectionStatus
	// Write should read in the AirbyteMessages and write any that are of
	// type AirbyteRecordMessage to the underlying data store. The destination
	// should return an error if any of the messages it receives do not match
	// the structure described in the catalog.
	//
	// Write should continue to write until the passed channel is closed.
	// Once the channel is closed, and all passed messages, have been written to
	// the underlying data store, it should send on the done channel to signify
	// to the caller that all work is done.
	//
	// On a cancelled context, Write should finish writing any messages already
	// read from the message channel and then send on the done channel (and
	// return context.Canceled).
	Write(ctx context.Context, config types.JSONData, catalog *types.ConfiguredAirbyteCatalog, messages <-chan *types.AirbyteMessage, done chan<- bool) error
}
