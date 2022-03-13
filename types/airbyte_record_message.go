package types

type AirbyteRecordMessage struct {
	Stream string
	Data JSONData
	EmittedAt int64
	Namespace string
}
