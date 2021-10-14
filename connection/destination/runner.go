package destination

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/adayNU/airbyte-go/types"
	"github.com/jessevdk/go-flags"
)

type opts struct {
	Config string `long:"config" description:"Path to configuration file"`

	Catalog string `long:"catalog" description:"Path to catalog file"`
}

func (o *opts) ParsedConfig() (types.JSONData, error) {
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

func (o *opts) ParsedCatalog() (*types.ConfiguredAirbyteCatalog, error) {
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

const (
	Spec  = "spec"
	Check = "check"
	Write = "write"
)

func Run(d Destination) {
	var opts = &opts{}

	var _, err = flags.Parse(opts)
	if err != nil {
		panic(err)
	}

	var out = &types.AirbyteMessage{}

	switch os.Args[1] {
	case Spec:
		out.Type = types.Spec
		out.Spec = d.Spec()
	case Check:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			panic(err)
		}

		out.Type = types.ConnectionStatus
		out.ConnectionStatus = d.Check(cfg)
	case Write:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			panic(err)
		}

		var catalog *types.ConfiguredAirbyteCatalog
		catalog, err = opts.ParsedCatalog()
		if err != nil {
			panic(err)
		}

		var messages = make(chan *types.AirbyteMessage)
		defer close(messages)

		go d.Write(cfg, catalog, messages)

		var scanner = bufio.NewScanner(os.Stdin)
		for {
			var msg = &types.AirbyteMessage{}

			var ok = scanner.Scan()
			if !ok {
				if err = scanner.Err(); err != nil {
					panic(err)
				}
				break
			}

			err = json.Unmarshal(scanner.Bytes(), msg)
			if err != nil {
				panic(err)
			}

			messages <- msg
		}

		return
	default:
		panic("not jug")
	}

	var b []byte
	b, err = json.Marshal(out)
	if err != nil {
		panic(err)
	}

	var w = bufio.NewWriter(os.Stdout)
	_, _ = w.Write(b)
	_ = w.Flush()
}
