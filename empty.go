package golog

type VoidLogger struct{}

func (v VoidLogger) Print(string) {}

func (v VoidLogger) WithField(key string, object interface{}) StructuredLogger {
	return v
}
