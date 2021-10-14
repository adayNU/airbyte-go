package main

import (
	"fmt"

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

func (m *mockDestination) Write(config types.JSONData, catalog *types.ConfiguredAirbyteCatalog, messages <-chan *types.AirbyteMessage) {
	for {
		select {
		case foo := <-messages:
			fmt.Println("test", foo)
		}
	}
}

func main() {
	destination.Run(&mockDestination{})
}
