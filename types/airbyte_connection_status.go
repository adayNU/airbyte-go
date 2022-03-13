package types

type Status int

const (
	Succeeded Status = iota
	Failed
)

type AirbyteConnectionStatus struct {
	Status  Status
	Message string
}
