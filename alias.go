package box

import (
	"github.com/s-z-z/box/log"
	"github.com/s-z-z/box/signals"
)

var (
	SetupSignalContext = signals.SetupSignalContext

	Log        = log.Log
	LoggerFrom = log.FromContext
	LoggerInto = log.IntoContext
	SetLogger  = log.SetLogger
)
