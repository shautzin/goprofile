package goprofile

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

// Flag 名字
var ProfilesFlagName = "profiles"

// 配置文件路径
var ConfigFileFolder = "config/"

// 默认加载的配置文件名前缀
var ConfigurationFilePrefix = "config"

// 默认加载的配置文件名
var DefaultConfigurationFileName = ConfigurationFilePrefix + ".env"

// Inspired from spring framework
// 程序传入参数：-profiles=prod,test
// 自动解析 application.prod.env application.test.env config.env 文件
// 排在最前优先级最高，即 prod 会覆盖 test 和默认文件中的值
func Load() {

	profiles := flag.String(ProfilesFlagName, "", "指定配置 profiles")
	flag.Parse()
	envFiles := getEnvFiles(*profiles)
	log.Println("正在导入配置文件", envFiles)

	// 写在最前面的优先
	for i := range envFiles {
		err := godotenv.Load(envFiles[i])
		if err != nil {
			log.Println("导入配置文件失败", err)
		}
	}
}

// 根据传入的 profile 构建配置文件列表
func getEnvFiles(profile string) []string {
	envFiles := []string{}

	if profile != "" {
		profiles := strings.Split(profile, ",")

		for i := range profiles {
			envFiles = append(envFiles, ConfigFileFolder+ConfigurationFilePrefix+"."+profiles[i]+".env")
		}
	}

	// 最后一个优先级最低
	envFiles = append(envFiles, ConfigFileFolder+DefaultConfigurationFileName)
	return envFiles
}

// 获取配置值
func GetEnv(name string) string {
	return os.Getenv(name)
}
