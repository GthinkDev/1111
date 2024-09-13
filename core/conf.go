package core

import (
	"fmt"
	"gin-blog/config"
	"gin-blog/global"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "settings.yaml"

// InitConfig 读取 yaml 文件的配置
func InitConfig() {
	c := &config.Config{}
	yamlConfig, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("读取配置文件出错: %v", err))
	}

	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatalf("解析配置文件出错: %v", err)
	}

	log.Println("yaml 配置文件读取成功~~")
	//fmt.Println(" 配置文件读取成功~~")
	global.Config = c
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile("settings.yaml", byteData, 0644)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功～")
	return err
}
