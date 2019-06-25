package config

import (
	"flag"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type logConf struct {
	Level string
	Path  string
}

// Base 基本配置
type Base struct {
	Port  int
	Debug bool
	Log   logConf
}

// Mysql 配置
type MysqlConf struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string `yaml:"dbName"`
}

// 如有新的配置按照mysql的配置增加即可

var C = struct {
	Base      `yaml:",inline"`
	MysqlConf MysqlConf `yaml:"mysql"`
}{}

func init() {
	flag.String("c", "", "config file path")

	file := ""
	for i := range os.Args {
		if os.Args[i] == "-c" && len(os.Args) > i+1 {
			file = os.Args[i+1]
			break
		}
	}
	if file == "" {
		logrus.Fatalln("not specify config file")
		return
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Fatalf("read config file %s failed, error: %v", file, err)
	} else {
		logrus.Infof("read config file %s success", file)
	}

	if err := yaml.Unmarshal(content, &C); err != nil {
		logrus.Fatalf("parse config file %s to base failed, error: %v", file, err)
	}

	if C.Log.Level == "debug" {
		logrus.Infoln("set log level to debug")
		logrus.SetLevel(logrus.DebugLevel)
	}
}
