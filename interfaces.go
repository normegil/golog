package log

type SimpleLogger interface {
	Print(string)
}

type StructuredLogger interface {
	SimpleLogger
	WithField(key string, object interface{}) StructuredLogger
}
