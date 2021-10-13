package types

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

type ConfiguredAirbyteCatalog struct{}

type ConfiguredAirbyteStream struct {
	Stream              AirbyteStream
	SyncMode            SyncMode
	CursorField         []string
	DestinationSyncMode DestinationSyncMode
	PrimaryKey          [][]string
}
