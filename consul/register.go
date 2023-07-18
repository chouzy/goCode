package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// RegisterService 服务注册
func RegisterService(c *Server) error {
	conf := api.DefaultConfig()
	conf.Address = c.ConsulAddr
	conf.Token = c.ConsulToken
	client, err := api.NewClient(conf)
	if err != nil {
		fmt.Printf("new client error: %v\n", err)
		return err
	}

	asr := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", c.Name, c.IP, c.Port), // 节点名
		Name:    c.Name,                                        // 服务名
		Tags:    c.Tag,                                         // 标签
		Address: c.IP,                                          // 服务ip
		Port:    c.Port,                                        // 服务端口
		Check: &api.AgentServiceCheck{
			Interval:                       c.Interval.String(),                           // 健康检查时间间隔
			GRPC:                           fmt.Sprintf("%v:%v/%v", c.IP, c.Port, c.Name), // 使用grpc执行健康检查, health.Check() 方法进行响应
			DeregisterCriticalServiceAfter: c.Deregister.String(),                         // 注销时间
		},
	}
	if err := client.Agent().ServiceRegister(asr); err != nil {
		fmt.Printf("service register error: %v\n", err)
		return err
	}
	return nil
}
