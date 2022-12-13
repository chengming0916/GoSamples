package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	Port       int    `yaml:"Port"`
	DbHost     string `yaml:"DbHost"`
	DbPort     int    `yaml:"DbPort"`
	DbName     string `yaml:"DbName"`
	DbUser     string `yaml:"DbUser"`
	DbPassword string `yaml:"DbPassword"`
	Test       []int  `yaml:"Test"`

	Test1 struct {
		Test2 string `yaml:"Test2"`
		Test3 int    `yaml:"Test3"`
	} `yaml:"Test1"`
}

func main() {

	yamlFile, err := ioutil.ReadFile("application.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var cfg *Conf
	err = yaml.Unmarshal([]byte(yamlFile), &cfg)
	if err != nil {
		log.Fatalf("unmarshal: %v", err)
	}

	log.Printf("cfg:%#v \n", cfg)
	// log.Printf("test2:%s \n", cfg.test2)

	// for index, item := range cfg.test3 {
	// 	log.Printf("test3.item%d:%d \n", index, item)
	// }

	// log.Printf("test5:%d \n", cfg.test4.test5)
	// log.Printf("test6:%s \n", cfg.test4.test6)

	configMap := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, configMap)

	if err != nil {
		log.Fatalf("Unmarshal: %v when to map", err)
	}

	for k, v := range configMap {
		_, isStr := v.(string)
		_, isInt32 := v.(int32)
		_, isBool := v.(bool)
		if isStr {
			log.Printf("key:%s, value(string):%s \n", k, v)
		} else if isInt32 {
			log.Printf("key:%s, value(int32):%d \n", k, v)
		} else if isBool {
			log.Printf("key:%s, value(bool):%t \n", k, v)
		} else {
			log.Printf("key:%s, value:%s \n", k, v)
		}
	}

	// 使用map表示已经提前知道结构是什么样的， 可以将map提取为自己想要的格式
	mysqlValueInterface := configMap["mysql"]
	mysqlValue, ok := mysqlValueInterface.(interface{})
	if ok {
		mysqlMap, ok := mysqlValue.(map[interface{}]interface{})
		if ok {
			log.Printf(" map解析结果。 mysqlHost:%s, port:%d", mysqlMap["host"], mysqlMap["port"])
		} else {
			log.Printf("no mysql info")
		}
	} else {
		log.Printf("no mysql value is not interface")
	}

}
