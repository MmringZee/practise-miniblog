package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	// defaultHomeDir 定义放置 miniblog 服务配置的默认目录.
	defaultHomeDir = ".miniblog"

	// defaultConfigName 指定 miniblog 服务的默认配置文件名.
	defaultConfigName = "mb-apiserver.yaml"
)

// onInitialize 设置需要读取的配置文件名、环境变量，并将其内容读取到 viper 中.
func onInitialize() {
	if configFile != "" {
		// 当配置文件路径不为空时, 从命令行选项指定的配置文件中读取配置参数
		viper.SetConfigFile(configFile)
	} else {
		// 用户启动程序没有传入`--config`参数, 使用默认配置文件路径和名称
		for _, dir := range searchDirs() {
			// 将 dir 目录加入到配置文件搜索路径
			viper.AddConfigPath(dir)
		}

		// 设置配置文件格式为 YAML
		viper.SetConfigType("yaml")

		// 配置文件名称
		viper.SetConfigName(defaultConfigName)
	}

	// 读取环境变量并设置前缀
	setupEnvironmentVariables()

	// 读取配置文件.如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read viper configuration file, err: %v", err)
	}

	// 打印当前使用的配置文件，方便调试
	log.Printf("Using config file: %s", viper.ConfigFileUsed())
}

// searchDirs 返回默认的配置文件搜索目录
func searchDirs() []string {
	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	// 若获取用户主目录失败, 打印错误信息并退出
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}

// setupEnvironmentVariables 配置环境变量规则
func setupEnvironmentVariables() {
	// 允许 viper 自动匹配环境变量
	viper.AutomaticEnv()
	// 设置环境变量前缀
	// 防止同服务器上其他服务环境变量的影响
	viper.SetEnvPrefix("MINIBLOG")
	// 替换环境变量 key 中的分隔符"."和"-"为"_"
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

// filePath 获取默认配置文件的完整路径.
func filePath() string {
	home, err := os.UserHomeDir()
	// 如果不能获取用户主目录，则记录错误并返回空路径
	cobra.CheckErr(err)
	return filepath.Join(home, defaultHomeDir, defaultConfigName)
}
