package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Config 配置结构体，用于存储应用相关的配置.
// 不用 viper.Get，是因为这种方式能更加清晰的知道应用提供了哪些配置项.
type Config struct {
	ServerMode string
	JWTKey     string
	Expiration time.Duration
}

// UnionServer 定义一个联合服务器. 根据 ServerMode 决定要启动的服务器类型.
type UnionServer struct {
	cfg *Config
}

// NewUnionServer 根据配置创建联合服务器.
func (cfg *Config) NewUnionServer() (*UnionServer, error) {
	return &UnionServer{cfg: cfg}, nil
}

// Run 运行应用
func (u *UnionServer) Run() error {
	fmt.Printf("ServerMode from ServerOptions: %s\n", u.cfg.JWTKey)
	fmt.Printf("ServerMode from Viper: %s\n\n", viper.GetString("jwt-key"))

	jsonData, _ := json.MarshalIndent(u.cfg, "", " ")
	fmt.Println(string(jsonData))

	select {}
}
