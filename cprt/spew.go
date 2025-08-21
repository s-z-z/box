package cprt

import "github.com/davecgh/go-spew/spew"

func Spew(o ...interface{}) string {
	s := spew.ConfigState{
		Indent:                  "  ",
		DisableCapacities:       true,
		DisableMethods:          true,
		DisablePointerAddresses: true,
		DisablePointerMethods:   true,
	}
	return s.Sdump(o...)
}

func SpewInfo(o ...interface{}) {
	Info(Spew(o...))
}

func SpewWarning(o ...interface{}) {
	Warning(Spew(o...))
}
func SpewDebug(o ...interface{}) {
	Debug(Spew(o...))
}

func SpewError(o ...interface{}) {
	Error(Spew(o...))
}
