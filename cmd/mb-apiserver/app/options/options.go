package options

import (
	"errors"
	"fmt"
	"github.com/MmringZee/practise-miniblog/internal/apiserver"
	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"time"
)

var availableServerModes = sets.New(
	"grpc",
	"grpc-gateway",
	"gin",
)

// ServerOptions 包含服务器配置选项
type ServerOptions struct {
	// ServerMode 定义服务器模式: gRPC、Gin HTTP、HTTP Reverse Proxy
	ServerMode string `json:"server-mode" mapstructure:"server-mode"`
	// JWTKey 定义 JWT 密钥
	JWTKey string `json:"jwt-key" mapstructure:"jwt-key"`
	// Expiration 定义 JWT Token 的过期时间
	Expiration time.Duration `json:"expiration" mapstructure:"expiration"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		ServerMode: "grpc-gateway",
		JWTKey:     "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		Expiration: 2 * time.Hour,
	}
}

// AddFlags 将 ServerOptions 的选项绑定到命令行标志.
// 通过使用 pflag 包，可以实现从命令行中解析这些选项的功能.
// 这里设置的信息都可以通过 -h 访问.
func (receiver *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&receiver.ServerMode, "server-mode", receiver.ServerMode, fmt.Sprintf("Server mode, available options: %v", availableServerModes.UnsortedList()))
	fs.StringVar(&receiver.JWTKey, "jwt-key", receiver.JWTKey, "JWT signing key. Must be at least 6 characters long.")
	fs.DurationVar(&receiver.Expiration, "expiration", receiver.Expiration, "The expiration duration of JWT tokens.")
}

// Validate 校验 ServerOptions 中的选项是否合法.
func (receiver *ServerOptions) Validate() error {
	errs := []error{}

	// 校验 ServerMode 是否有效s
	// fmt.Errorf 一般用于需要插入动态内容时
	if !availableServerModes.Has(receiver.ServerMode) {
		errs = append(errs, fmt.Errorf("invalid server mode: must be one of %v", availableServerModes.UnsortedList()))
	}

	// 校验 JWTKey 长度
	// errors.New 一般用于不需要嵌入动态内容时, 性能更高
	if len(receiver.JWTKey) < 6 {
		errs = append(errs, errors.New("JWTKey must be at least 6 characters long"))
	}

	// 合并所有错误并返回
	return utilerrors.NewAggregate(errs)
}

// Config 基于 ServerOptions 创建新的 apiserver.Config.
func (s *ServerOptions) Config() (*apiserver.Config, error) {
	// cmd/mb-apiserver/app/ 目录下的文件导入了 internal/apiserver 目录下的包，也就是"控制面"依赖"数据面". 为了避免循环依赖，要避免反向导入.
	return &apiserver.Config{
		ServerMode: s.ServerMode,
		JWTKey:     s.JWTKey,
		Expiration: s.Expiration,
	}, nil
}
