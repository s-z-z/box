package cprt

import "github.com/fatih/color"

func Green(format string, a ...interface{}) string {
	return color.HiGreenString(format, a...)
}
func Blue(format string, a ...interface{}) string {
	return color.HiBlueString(format, a...)
}
func Yellow(format string, a ...interface{}) string {
	return color.HiYellowString(format, a...)
}
func Red(format string, a ...interface{}) string {
	return color.HiRedString(format, a...)
}
func Magenta(format string, a ...interface{}) string {
	return color.HiMagentaString(format, a...)
}
