package gogen

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

const (
	defaultMysqlTableName = "test"
	defaultMysqlUser      = "root"
	defaultMysqlPassword  = "password"
	defaultMysqlAddr      = "127.0.0.1"
	defaultMysqlPort      = 3306

	defaultPort = 8888
	defaultHost = "0.0.0.0"
	etcDir      = "etc"

	// LOG_MODE 日志打印模式，console 控制台 console,file,volume
	LOG_MODE_CONSOLE = "console"
	LOG_MODE_FILE    = "file"
	LOG_MODE_VOLUME  = "volume"

	// LOG_ENCODING 日志格式, json 格式 或者 plain 纯文本
	LOG_ENCODING_PLAIN = "plain"
	LOG_ENCODING_JSON  = "json"

	defaultLogTimeFormat = "2006-01-02T15:04:05.000Z07:00" // 日期格式化
	defaultLogPath       = "logs"                          // 日志在文件输出模式下，日志输出路径

	// LOG_LEVEL 日志级别
	LOG_LEVEL_DEBUG = "debug"

	defaultCompress               = false // 是否压缩日志
	defaultStat                   = true  //是否开启 stat 日志
	defaultLogKeepDays            = 180   // 日志保留天数，只有在文件模式才会生效
	defaultLogStackCoolDownMillis = 100   // 堆栈打印冷却时间
	defaultLogMaxSiz              = 100

	LOG_ROTATION_DAILY = "daily"
)

//go:embed etc.tpl
var etcTemplate string

// 数据库配置
type MysqlCfg struct {
	TableName string `json:"table_name"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Addr      string `json:"addr"`
	Port      int64  `json:"port"`
}

// 日志配置
type LogCfg struct {
	ServiceName         string `json:"service_name"`
	Mode                string `json:"mode"`
	Encoding            string `json:"encoding"`
	TimeFormat          string `json:"time_format"`
	Path                string `json:"path"`
	Level               string `json:"level"`
	Compress            bool   `json:"compress"`
	Stat                bool   `json:"stat"`
	KeepDays            int64  `json:"keep_days"`
	StackCoolDownMillis int64  `json:"stack_cool_down_millis"`
	MaxSize             int64  `json:"max_size"`
	Rotation            string `json:"rotation"`
}

// etc配置文件
type Config struct {
	ServiceName string     `json:"service_name"`
	Host        string     `json:"host"`
	Port        int64      `json:"port"`
	Mysql       []MysqlCfg `json:"mysql"`
	Log         LogCfg     `json:"log"`
}

// MysqlCfg的构造函数
func (m *MysqlCfg) initMysqlCfg(tableName, user, password, addr string, port int64) {
	m.TableName = tableName
	m.User = user
	m.Password = password
	m.Addr = addr
	m.Port = port
}

// 初始化
func (c *Config) init(api *spec.ApiSpec) {
	service := api.Service
	c.ServiceName = service.Name
	c.Host = defaultHost
	c.Port = defaultPort
	c.Mysql = []MysqlCfg{
		{}, // 创建一个空的MysqlCfg实例
	}
	c.Mysql[0].initMysqlCfg(defaultMysqlTableName, defaultMysqlUser, defaultMysqlPassword, defaultMysqlAddr, defaultMysqlPort)

	c.Log = LogCfg{
		ServiceName:         service.Name,
		Mode:                LOG_MODE_CONSOLE,
		Encoding:            LOG_ENCODING_PLAIN,
		TimeFormat:          defaultLogTimeFormat,
		Path:                defaultLogPath,
		Level:               LOG_LEVEL_DEBUG,
		Compress:            defaultCompress,
		Stat:                defaultStat,
		KeepDays:            defaultLogKeepDays,
		StackCoolDownMillis: defaultLogStackCoolDownMillis,
		MaxSize:             defaultLogMaxSiz,
		Rotation:            LOG_ROTATION_DAILY,
	}
}

func genEtc(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	serviceName := api.Service.Name
	filename, err := format.FileNamingFormat(cfg.NamingFormat, serviceName)
	if err != nil {
		return err
	}

	configData := Config{}
	configData.init(api)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          etcDir,
		filename:        fmt.Sprintf("%s.yaml", filename),
		templateName:    "etcTemplate",
		category:        category,
		templateFile:    etcTemplateFile,
		builtinTemplate: etcTemplate,
		data:            configData,
	})
}
