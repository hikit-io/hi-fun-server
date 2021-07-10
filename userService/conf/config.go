package conf

import "os"

type TypeModel string

const (
	ModelDev  TypeModel = "dev"
	ModelTest TypeModel = "test"
	ModelPro  TypeModel = "pro"
)

func (t TypeModel) LoadConf() {
	switch t {
	case ModelDev:
		basicConf = BasicConf{
			Port:  "",
			Model: "",
		}
	case ModelTest:
		basicConf = BasicConf{
			Port:  "",
			Model: "",
		}
	case ModelPro:
		basicConf = BasicConf{
			Port:  "",
			Model: "",
		}
	}
}

type BasicConf struct {
	Port  string
	Model TypeModel
}

type MysqlConf struct {
	Addr string
}

type RedisConf struct {
	Addr string
}

var (
	basicConf = BasicConf{Port: "8020"}
	mysqlConf = MysqlConf{Addr: ""}
	redisConf = RedisConf{Addr: ""}
)

func init() {
	mode, _ := os.LookupEnv("MODE")
	basicConf.Model = TypeModel(mode)

}
