package log

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                         // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                      // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                      // 输出
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`               // 日志文件夹
	EncodeLevel   string `mapstructure:"encodeLevel" json:"encodeLevel" yaml:"encodeLevel"`       // 编码级别
	StacktraceKey string `mapstructure:"stacktraceKey" json:"stacktraceKey" yaml:"stacktraceKey"` // 栈名

	MaxAge       int  `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`                   // 日志留存时间
	ShowLine     bool `mapstructure:"showLine" json:"showLine" yaml:"showLine"`             // 显示行
	LogInConsole bool `mapstructure:"logInConsole" json:"logInConsole" yaml:"logInConsole"` // 输出控制台
}

// ZapEncodeLevel 获取编码器级别
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	// 以json格式输出时如果使用带颜色的编码器会出现乱码
	switch z.EncodeLevel {
	case "LowercaseLevelEncoder": // 小写编码器
		return zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 将日志级别转换为 zapcore.Level 类型
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
