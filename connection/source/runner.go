package source

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/adayNU/airbyte-go/connection/protocol"
	"github.com/adayNU/airbyte-go/types"
	"github.com/jessevdk/go-flags"
)

func Run(s Source) {
	var opts = &protocol.Options{}

	var _, err = flags.Parse(opts)
	if err != nil {
		panic(err)
	}

	var out = &types.AirbyteMessage{}
	var w = bufio.NewWriter(os.Stdout)

	switch os.Args[1] {
	case protocol.Spec:
		out.Type = types.Spec
		out.Spec = s.Spec()
	case protocol.Check:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			panic(err)
		}

		out.Type = types.ConnectionStatus
		out.ConnectionStatus = s.Check(cfg)
	case protocol.Discover:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			panic(err)
		}

		out.Type = types.Catalog
		out.Catalog = s.Discover(cfg)
	case protocol.Read:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			panic(err)
		}

		var catalog *types.ConfiguredAirbyteCatalog
		catalog, err = opts.ParsedCatalog()
		if err != nil {
			panic(err)
		}

		var state types.JSONData
		state, err = opts.ParsedState()
		if err != nil {
			panic(err)
		}

		var messages = s.Read(cfg, catalog, state)

		for {
			select {
			case msg, ok := <-messages:
				if ok {
					var b []byte
					b, err = json.Marshal(msg)
					if err != nil {
						panic(err)
					}

					_, err = w.Write(b)
					if err != nil {
						panic(err)
					}
				} else {
					break
				}
			}
		}

		return
	default:
		panic("unknown command")
	}

	var b []byte
	b, err = json.Marshal(out)
	if err != nil {
		panic(err)
	}

	_, _ = w.Write(b)
	_ = w.Flush()
}
