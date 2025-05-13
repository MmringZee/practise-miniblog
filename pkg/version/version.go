package version

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"runtime"
)

var (
	// gitVersion 是语义化的版本号。
	gitVersion = "v0.0.0-master+$Format:%h$"

	// buildDate 是 ISO8601 格式的构建时间， $（date -u +'%Y-%m-%dT%H:%M:%SZ'） 命令的输出。
	buildDate = "1970-01-01T00:00:00Z"

	// gitCommit 是 Git 的 SHA1 值，$（git rev-parse HEAD） 命令的输出。
	gitCommit = "$Format:%H$"

	// gitTreeState 代表构建时 Git 仓库的状态，可能的值有：clean， dirty。
	gitTreeState = ""
)

// Info 包含了版本信息.
type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String 返回人性化的版本信息字符串.
func (i Info) String() string {
	return i.GitVersion
}

// ToJSON 以 JSON 格式返回版本信息.
func (i Info) ToJSON() string {
	s, _ := json.Marshal(i)
	return string(s)
}

// Text 将版本信息编码为 UTF-8 格式的文本, 并返回.
// 引入 github.com/gosuri/uitable 包生成表格化输出.
// 可以方便地控制命令行输出的格式, 提高可读性.
func (i Info) Text() string {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", i.GitVersion)
	table.AddRow("gitCommit:", i.GitCommit)
	table.AddRow("gitTreeState:", i.GitTreeState)
	table.AddRow("buildDate:", i.BuildDate)
	table.AddRow("goVersion:", i.GoVersion)
	table.AddRow("compiler:", i.Compiler)
	table.AddRow("platform:", i.Platform)

	return table.String()
}

// Get 返回详尽的代码库版本信息，用来标明二进制文件由哪个版本的代码构建。
func Get() Info {
	// 以下变量通常由 -ldflags 进行设置
	return Info{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
