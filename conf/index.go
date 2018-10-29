package conf

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

type MySQL struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type Config struct {
	MySQL       MySQL
	HttpAddr    string
	HttpPort    string
	JwtSecret   string
	PageLimit   uint
	SessionName string
	Home        string
	PassLen     uint
	TitleLen    uint
	BasePublic string
	PublicHost string
}

var GlobalConfig *Config

// 这里会导致不好测试
//func init() {
//	GlobalConfig = getConfig()
//	// TODO 自动解析json文件
//}

func New() *Config {
	if GlobalConfig == nil {
		GlobalConfig = getConfig()
	}
	return GlobalConfig

}

func getPath() (filePath string) {

	const dir = "conf/json"
	m := gin.Mode()

	if m == gin.DebugMode {
		filePath = fmt.Sprintf("%s/%s.config.json", dir, "dev")
	} else if m == gin.ReleaseMode {
		filePath = fmt.Sprintf("%s/%s.config.json", dir, "pro")
	} else if m == gin.TestMode {
		filePath = fmt.Sprintf("%s/%s.config.json", dir, "test")
	}
	return
}

// 初始化时获取Config对象
func getConfig() *Config {
	var c Config
	configFilePath := getPath()
	if len(configFilePath) == 0 {
		panic("can't get config path")
	}
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil && err != io.EOF {
		log.Fatal("read config file error ", err.Error())
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Fatal("transform config filr error ", err.Error())
	}
	return &c
}

// 获取http监听地址和端口URL
func (config *Config) GetHttpAddrPort() string {
	if len(config.HttpAddr) == 0 {
		config.HttpAddr = "127.0.0.1"
		log.Warnf("set default server host/ip is %s", config.HttpAddr)
	}

	if len(config.HttpPort) == 0 {
		config.HttpPort = "3000"
		log.Warnf("set default server port is %s", config.HttpPort)
	}
	ap := config.HttpAddr + ":" + config.HttpPort
	return ap
}

// 获取数据库的链接URL 主要用于常规增删改查
func (config *Config) GetMySQLUrl() (string, error) {

	var mysql = &config.MySQL
	if len(mysql.Username) == 0 {
		return "", errors.New("database username must config")
	}
	if len(mysql.Password) == 0 {
		return "", errors.New("database password must config")
	}
	if len(mysql.Database) == 0 {
		return "", errors.New("database name must config")
	}
	if len(mysql.Host) == 0 {
		mysql.Host = "127.0.0.1"
		log.Warnf("set default mysql host/ip is %s", mysql.Host)
	}
	if len(mysql.Port) == 0 {
		mysql.Port = "3306"
		log.Warnf("set default mysql port is %s", mysql.Port)
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port, config.MySQL.Database)
	return url, nil
}

//
//// 获取MySQL的URL，主要用于创建数据库
//func (config Config) GetMySqlURL() string {
//	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port)
//}
//
//// 获取config对象
//func GetConfig()(  *Config){
//	if config == nil {
//		return NewConfig(cmds.GetEnv())
//	}else {
//		return config
//	}
//}
