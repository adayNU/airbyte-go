package types

type LogLevel int

const (
	Fatal LogLevel = iota
	Critical
	Error
	Warn
	Warning
	Info
	Debug
	Trace
)

type AirbyteLogMessage struct {
	Level   LogLevel
	Message string
}
