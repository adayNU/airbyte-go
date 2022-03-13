package source

import (
	"context"
	"encoding/json"
	"os"

	"github.com/adayNU/airbyte-go/connection/protocol"
	"github.com/adayNU/airbyte-go/types"
)

// Run will give the appropriate response for the
// given command for Source |s|.
func Run(s Source, opts protocol.Protocol) error {
	var out = &types.AirbyteMessage{}

	switch os.Args[1] {
	case protocol.Spec:
		out.Type = types.Spec
		out.Spec = s.Spec()
	case protocol.Check:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			return err
		}

		out.Type = types.ConnectionStatus
		out.ConnectionStatus = s.Check(cfg)
	case protocol.Discover:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			return err
		}

		out.Type = types.Catalog
		out.Catalog = s.Discover(cfg)
	case protocol.Read:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			return err
		}

		var catalog *types.ConfiguredAirbyteCatalog
		catalog, err = opts.ParsedCatalog()
		if err != nil {
			return err
		}

		var state types.JSONData
		state, err = opts.ParsedState()
		if err != nil {
			return err
		}

		var ctx = context.Background()

		var reader = s.Read(ctx, cfg, catalog, state)

		for {
			var wrapped = &types.AirbyteMessage{}

			select {
			case msg := <-reader.Records():
				wrapped = &types.AirbyteMessage{
					Type:   types.Record,
					Record: msg,
				}
			case msg := <-reader.States():
				wrapped = &types.AirbyteMessage{
					Type:  types.State,
					State: msg,
				}
			case <-reader.Done():
				return reader.Err()
			}

			var b []byte
			b, err = json.Marshal(wrapped)
			if err != nil {
				return err
			}

			_, err = os.Stdout.WriteString(string(b) + "\n")
			if err != nil {
				return err
			}
		}

		return nil
	default:
		panic("unknown command")
	}

	var b, err = json.Marshal(out)
	if err != nil {
		return err
	}

	_, err = os.Stdout.WriteString(string(b) + "\n")
	return err
}
