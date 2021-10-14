package source

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/adayNU/airbyte-go/types"
	"gopkg.in/check.v1"
)

type RunnerSuite struct{}

func (r *RunnerSuite) TestSpec(c *check.C) {
	os.Args = []string{"cmd", "spec"}

	var std = os.Stdout
	defer func() {
		os.Stdout = std
	}()

	var read, write, err = os.Pipe()
	c.Assert(err, check.IsNil)

	os.Stdout = write

	err = Run(&mockSource{})
	c.Check(err, check.IsNil)

	_ = write.Close()

	var got []byte
	got, err = ioutil.ReadAll(read)
	os.Stdout = std
	c.Check(err, check.IsNil)

	c.Check(string(got), check.Equals, `{"Type":3,"Spec":{"DocumentationURL":"","ChangelogURL":"","ConnectionSpecification":null,"SupportsIncremental":false,"SupportsNormalization":false,"SupportsDBT":false,"SupportedDestinationSyncModes":null,"AuthSpecification":null},"ConnectionStatus":null,"Catalog":null,"Record":null,"State":null}`+"\n")
}

type mockSource struct{}

func (m *mockSource) Spec() *types.ConnectorSpecification {
	return &types.ConnectorSpecification{}
}

var badConfig = "i'm bad"

func (m *mockSource) Check(config types.JSONData) *types.AirbyteConnectionStatus {
	if config == badConfig {
		return &types.AirbyteConnectionStatus{Status: types.Failed}
	}

	return &types.AirbyteConnectionStatus{Status: types.Succeeded}
}
func (m *mockSource) Discover(config types.JSONData) *types.AirbyteCatalog {
	return &types.AirbyteCatalog{}
}

func (m *mockSource) Read(ctx context.Context, config types.JSONData, catalog *types.ConfiguredAirbyteCatalog, state types.JSONData) <-chan types.AirbyteMessage {
	// Need to do better here.
	return make(<-chan types.AirbyteMessage)
}

var _ = check.Suite(&RunnerSuite{})
