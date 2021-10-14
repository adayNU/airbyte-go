package protocol

import (
	"encoding/json"
	"os"

	"github.com/adayNU/airbyte-go/types"
)

type Options struct {
	Config  string `long:"config" description:"Path to configuration file"`
	Catalog string `long:"catalog" description:"Path to catalog file"`
	State   string `long:"state" description:"Path to state file"`
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
