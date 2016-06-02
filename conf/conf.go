package conf

import (
	"encoding/json"
	"os"

	log "github.com/inconshreveable/log15"
)

//Conf 配置信息
type Conf struct {
	MongoURL    string
	MongoUser   string
	MongoPsw    string
	MongoAuthDB string
	MongoDB     string
}

//Cfg 全局配置文件变量
var Cfg Conf

//loadDataFromFile 读取文件
func loadDataFromFile(confFileURL string, v interface{}) error {
	file, err := os.Open(confFileURL)
	if err != nil {
		log.Error("打开配置文件错误", confFileURL, err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(v)
	if err != nil {
		log.Error("读取配置文件错误", confFileURL, err)
	}
	return err
}

//LoadConf 读取配置文件
func LoadConf() error {
	log.Info("Loading Conf data")
	var confFilePath string

	//配置文件路径
	confFilePath = "/Users/guoqiang/gopath/src/gin_miaomiaola/blog.conf"

	log.Debug("load conf", "filepath", confFilePath)
	err := loadDataFromFile(confFilePath, &Cfg)
	if err != nil {
		log.Error("加载配置失败", "confFilePath", confFilePath, "error", err)
		return err
	}
	log.Info("Conf data loaded")
	return nil
}

//init 初始化
func init() {
	log.Info("----开始初始化配置参数")
	err := LoadConf()
	if err != nil {
		log.Error("初始化参数失败", "error", err)
	}

	log.Info("++++完成参数配置")
}
