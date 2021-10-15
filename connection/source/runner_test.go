package source

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/adayNU/airbyte-go/types"
	"gopkg.in/check.v1"
)

type RunnerSuite struct {
	r, w *os.File
	old  *os.File
}

func (r *RunnerSuite) SetUpTest(c *check.C) {
	r.old = os.Stdout

	var err error
	r.r, r.w, err = os.Pipe()
	c.Assert(err, check.IsNil)

	os.Stdout = r.w
}

func (r *RunnerSuite) TearDownTest(_ *check.C) {
	r.ResetStdout()
}

func (r *RunnerSuite) ResetStdout() {
	os.Stdout = r.old
}

func (r *RunnerSuite) TestSpec(c *check.C) {
	os.Args = []string{"cmd", "spec"}

	var err = Run(&mockSource{}, &mockProtocol{})
	c.Check(err, check.IsNil)

	_ = r.w.Close()

	var got []byte
	got, err = ioutil.ReadAll(r.r)
	r.ResetStdout()

	c.Check(err, check.IsNil)
	c.Check(string(got), check.Equals, `{"Type":3,"spec":{"DocumentationURL":"","ChangelogURL":"","ConnectionSpecification":null,"SupportsIncremental":false,"SupportsNormalization":false,"SupportsDBT":false,"SupportedDestinationSyncModes":null,"AuthSpecification":null}}`+"\n")
}

func (r *RunnerSuite) TestCheck(c *check.C) {
	os.Args = []string{"cmd", "check", "--config", "test"}

	var err = Run(&mockSource{}, &mockProtocol{})
	c.Check(err, check.IsNil)

	_ = r.w.Close()

	var got []byte
	got, err = ioutil.ReadAll(r.r)
	r.ResetStdout()

	c.Check(err, check.IsNil)
	c.Check(string(got), check.Equals, `{"Type":4,"connectionStatus":{"Status":0,"Message":""}}`+"\n")
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

type mockProtocol struct{}

func (m *mockProtocol) ParsedConfig() (types.JSONData, error) { return nil, nil }
func (m *mockProtocol) ParsedCatalog() (*types.ConfiguredAirbyteCatalog, error) {
	return &types.ConfiguredAirbyteCatalog{}, nil
}
func (m *mockProtocol) ParsedState() (types.JSONData, error) { return nil, nil }

var _ = check.Suite(&RunnerSuite{})
