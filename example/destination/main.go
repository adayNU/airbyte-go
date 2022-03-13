package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adayNU/airbyte-go/connection/destination"
	"github.com/adayNU/airbyte-go/types"
)

type mockDestination struct{}

func (m *mockDestination) Spec() *types.ConnectorSpecification {
	return &types.ConnectorSpecification{
		DocumentationURL:              "foo",
		ChangelogURL:                  "bar",
		ConnectionSpecification:       map[string]string{"foo": "bar"},
		SupportsIncremental:           false,
		SupportsNormalization:         false,
		SupportsDBT:                   false,
		SupportedDestinationSyncModes: []types.DestinationSyncMode{types.Append},
	}
}

func (m *mockDestination) Check(types.JSONData) *types.AirbyteConnectionStatus {
	return &types.AirbyteConnectionStatus{
		Status:  types.Succeeded,
		Message: "connected",
	}
}

func (m *mockDestination) Write(_ context.Context, _ types.JSONData, _ *types.ConfiguredAirbyteCatalog, messages <-chan *types.AirbyteMessage, done chan<- error) {
	for {
		select {
		case msg, ok := <-messages:
			if ok {
				fmt.Println("test:", msg)
			} else {
				done <- nil
				return
			}
		}
	}
}

func main() {
	var err = destination.Run(&mockDestination{})
	if err != nil {
		log.Fatal(err)
	}
}
