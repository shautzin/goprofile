package goprofile

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"strings"
)

// 配置文件路径
const ApplicationFolder = "config/"

// 默认加载的配置文件名
const DefaultApplicationFileName = "application.env"

// Inspired from spring framework
// 程序传入参数：-profile=prod,test
// 自动解析 application.prod.env application.test.env application.env 文件
// 排在最前优先级最高，即 prod 会覆盖 test 和默认文件中的值
func Load() {

	profile := flag.String("profile", "", "specify application file")
	flag.Parse()
	envFiles := getEnvFiles(*profile)
	log.Println("正在导入配置文件：", envFiles)

	// 写在最前面的优先
	err := godotenv.Load(envFiles...)
	if err != nil {
		log.Println("导入配置文件失败", err)
	}
}

// 根据传入的 profile 构建配置文件列表
func getEnvFiles(profile string) []string {
	envFiles := []string{}

	if profile != "" {
		profiles := strings.Split(profile, ",")

		for i := range profiles {
			envFiles = append(envFiles, ApplicationFolder+"application."+profiles[i]+".env")
		}
	}

	// 最后一个优先级最低
	envFiles = append(envFiles, ApplicationFolder+DefaultApplicationFileName)
	return envFiles
}
