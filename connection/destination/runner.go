package destination

import (
	"bufio"
	"context"
	"encoding/json"
	"os"

	"github.com/adayNU/airbyte-go/connection/protocol"
	"github.com/adayNU/airbyte-go/types"
	"github.com/jessevdk/go-flags"
)

// Run will give the appropriate response for the
// given command for Destination |d|.
func Run(d Destination) error {
	var opts = &protocol.Options{}

	var _, err = flags.Parse(opts)
	if err != nil {
		return err
	}

	var out = &types.AirbyteMessage{}

	switch os.Args[1] {
	case protocol.Spec:
		out.Type = types.Spec
		out.Spec = d.Spec()
	case protocol.Check:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			return err
		}

		out.Type = types.ConnectionStatus
		out.ConnectionStatus = d.Check(cfg)
	case protocol.Write:
		var cfg, err = opts.ParsedConfig()
		if err != nil {
			return err
		}

		var catalog *types.ConfiguredAirbyteCatalog
		catalog, err = opts.ParsedCatalog()
		if err != nil {
			return err
		}

		var messages = make(chan *types.AirbyteMessage)
		var done = make(chan bool, 1)

		var ctx = context.Background()

		go d.Write(ctx, cfg, catalog, messages, done)

		var scanner = bufio.NewScanner(os.Stdin)
		for {
			var msg = &types.AirbyteMessage{}

			var ok = scanner.Scan()
			if !ok {
				close(messages)
				// TODO(Andy): Might want some sort of timeout for Write
				// gracefully shutting down (i.e. a select on the done channel
				// as well as a ticker).
				<-done

				if err = scanner.Err(); err != nil {
					return err
				}

				break
			}

			err = json.Unmarshal(scanner.Bytes(), msg)
			if err != nil {
				return err
			}

			messages <- msg
		}

		return nil
	default:
		panic("unknown command")
	}

	var b []byte
	b, err = json.Marshal(out)
	if err != nil {
		return err
	}

	_, err = os.Stdout.WriteString(string(b) + "\n")
	return err
}
