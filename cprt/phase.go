package cprt

import (
	"fmt"

	"k8s.io/klog/v2"

	"github.com/s-z-z/box/symbol"
)

func PhaseTitle(title string, a ...interface{}) {
	prefix := fmt.Sprintf("%s", symbol.PHASE)
	title = fmt.Sprintf("[Phase] %s", title)
	title = fmt.Sprintf(title, a...)
	klog.Infof("%s %s\n", prefix, Blue(title))
}

func PhaseResult(err error, format string, a ...interface{}) {
	result := fmt.Sprintf(format, a...)
	if err == nil {
		PhaseOKStr(result)
	} else {
		PhaseErrorStr(result)
	}
}

func PhaseOK() {
	Ok("%s  ", symbol.OK)
}
func PhaseOKStr(s string) {
	Ok("%s   %s", symbol.OK, s)
}

func PhaseOKFormat(format string, a ...interface{}) {
	Ok("%s   %s", symbol.OK, fmt.Sprintf(format, a...))
}

func PhaseWarning() {
	Warning("%s  ", symbol.WARN)
}

func PhaseWarningStr(s string) {
	Warning("%s   %s", symbol.WARN, s)
}

func PhaseError() {
	Error("%s  ", symbol.Error)
}

func PhaseErrorStr(s string) {
	Error("%s   %s", symbol.Error, s)
}

func PhaseEmoj(format string, a ...interface{}) {
	format = fmt.Sprintf("%s: %s", symbol.PHASE, format)
	Info(format, a...)
}
