package cprt

import "k8s.io/klog/v2"

func Ok(format string, a ...interface{}) {
	klog.Infoln(Green(format, a...))
}

func Debug(format string, a ...interface{}) {
	klog.Infoln(Magenta(format, a...))
}

func Flag(a interface{}) {
	klog.Infoln(Magenta("##### %s #####", a))
}

func Info(format string, a ...interface{}) {
	klog.Infoln(Blue(format, a...))
}

func Warning(format string, a ...interface{}) {
	klog.Infoln(Yellow(format, a...))
}

func Error(format string, a ...interface{}) {
	klog.Infoln(Red(format, a...))
}
