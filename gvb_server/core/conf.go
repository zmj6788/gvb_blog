package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const configFile = "settings.yaml"
// 读取配置操作
// 读取yaml文件的配置
func InitConf() {
	c := &config.Config{}
	yamlConfig, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("read config file error, %v", err))
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatal("unmarshal config file error, %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}

func SetYaml() error{
	data := global.Config
	//将结构体信息转换为yaml格式
	byteData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	//写入文件
	err = ioutil.WriteFile(configFile, byteData, 0644)
	if err != nil {
		return  err
	}
	return nil
}
