package service

// Logger the logger definition.
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
}

// StandardLogger godoc.
type StandardLogger interface {
	Printf(string, ...interface{})
}

// NopLogger godoc.
func NopLogger() Logger { return &nopLogger{} }

type nopLogger struct{}

func (n *nopLogger) Debugf(string, ...interface{}) {}
func (n *nopLogger) Infof(string, ...interface{})  {}
func (n *nopLogger) Warnf(string, ...interface{})  {}
func (n *nopLogger) Errorf(string, ...interface{}) {}
