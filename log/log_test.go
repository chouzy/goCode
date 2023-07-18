package log

import "testing"

func TestLog(t *testing.T) {
	t.Log("start...")
	z := Zap{
		Level:         "debug",
		Prefix:        "[goProject] ",
		Format:        "json",
		Director:      "log",
		EncodeLevel:   "LowercaseLevelEncoder",
		StacktraceKey: "stacktrace",
		MaxAge:        30,
		ShowLine:      true,
		LogInConsole:  true,
	}
	log := InitZap(&z)

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")

	t.Log("end..")
}
