package protocol

import (
	"encoding/json"
	"os"

	"github.com/adayNU/airbyte-go/types"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Config  string `long:"config" description:"Path to configuration file"`
	Catalog string `long:"catalog" description:"Path to catalog file"`
	State   string `long:"state" description:"Path to state file"`
}

// Protocol ... (I don't love this name, will brainstorm a better one)
type Protocol interface {
	ParsedConfig() (types.JSONData, error)
	ParsedCatalog() (*types.ConfiguredAirbyteCatalog, error)
	ParsedState() (types.JSONData, error)
}

// Init populates the config fields based on the passed arguments.
func (o *Options) Init() error {
	var _, err = flags.Parse(o)

	if err != nil {
		return err
	}

	return nil
}

func (o *Options) ParsedConfig() (types.JSONData, error) {
	var b, err = os.ReadFile(o.Config)
	if err != nil {
		return nil, err
	}

	var out types.JSONData
	err = json.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (o *Options) ParsedCatalog() (*types.ConfiguredAirbyteCatalog, error) {
	var b, err = os.ReadFile(o.Catalog)
	if err != nil {
		return nil, err
	}

	var out = &types.ConfiguredAirbyteCatalog{}
	err = json.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (o *Options) ParsedState() (types.JSONData, error) {
	var b, err = os.ReadFile(o.State)
	if err != nil {
		return nil, err
	}

	var out types.JSONData
	err = json.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

const (
	Spec     = "spec"
	Check    = "check"
	Discover = "discover"
	Read     = "read"
	Write    = "write"
)
