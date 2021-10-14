package types

import "errors"

type AirbyteCatalog struct {
	Streams []AirbyteStream
}

type AirbyteStream struct {
	Name                    string
	JSONSchema              JSONData
	SyncModes               []SyncMode
	SourceDefinedCursor     bool
	DefaultCursorEndField   []string
	SourceDefinedPrimaryKey [][]string
	Namespace               string
}

type SyncMode int

const (
	FullRefresh SyncMode = iota
	Incremental
)

type DestinationSyncMode int

const (
	Append DestinationSyncMode = iota
	Overwrite
	AppendDedup
)

type ConfiguredAirbyteCatalog struct {
	Streams []ConfiguredAirbyteStream
}

type ConfiguredAirbyteStream struct {
	Stream              AirbyteStream
	SyncMode            SyncMode
	CursorField         []string
	DestinationSyncMode DestinationSyncMode
	PrimaryKey          [][]string
}

// Validate returns an error if |c| is invalid according to the
// Airbyte Protocol.
// https://github.com/airbytehq/airbyte/blob/62826f82fdb198785ec788b3da71c771d140d645/airbyte-protocol/models/src/main/resources/airbyte_protocol/airbyte_protocol.yaml#L173-L200
func (c *ConfiguredAirbyteStream) Validate() error {
	if c.SyncMode == Incremental && (c.CursorField == nil || len(c.CursorField) == 0) {
		return errors.New("CursorField is required for sync mode INCREMENTAL")
	}
	if c.DestinationSyncMode == AppendDedup && (c.PrimaryKey == nil || len(c.PrimaryKey) == 0) {
		return errors.New("DestinationSyncMode is required for destination sync mode APPEND_DEDUP")
	}

	return nil
}
