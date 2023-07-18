package consul

import "time"

type Server struct {
	RunMode     string        `mapstructure:"runMode" yaml:"runMode" json:"runMode"`             // release/debug/test
	Name        string        `mapstructure:"name" yaml:"name" json:"name"`                      // 注册到consul的服务名
	Tag         []string      `mapstructure:"tag" yaml:"tag" json:"tag"`                         // 注意是字符数组形式, 不允许空数据
	IP          string        `mapstructure:"ip" yaml:"ip" json:"ip"`                            // 服务的IP
	Port        int           `mapstructure:"port" yaml:"port" json:"port"`                      // 服务的端口
	ConsulAddr  string        `mapstructure:"consulAddr" yaml:"consulAddr" json:"consulAddr"`    // consul的地址
	ConsulToken string        `mapstructure:"consulToken" yaml:"consulToken" json:"consulToken"` // consul的token
	Interval    time.Duration `mapstructure:"interval" yaml:"interval" json:"interval"`          // 健康检查间隔, 单位：秒
	Deregister  time.Duration `mapstructure:"deregister" yaml:"deregister" json:"deregister"`    // 注销时间, 相当于过期时间, 单位：分钟
}
