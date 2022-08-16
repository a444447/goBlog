package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode    string
	HttpPort   string
	JwtKey     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassport string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiLiuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("config.ini有误,", err)
		os.Exit(0)
	}
	LoadServer(file)
	LoadDatabase(file)
	LoadQiLiu(file)
}

func LoadServer(file *ini.File) {
	sec := file.Section("server")
	AppMode = sec.Key("AppMode").MustString("debug")
	HttpPort = sec.Key("HttpPort").MustString(":3000")
	JwtKey = sec.Key("JwtKey").MustString("89js82js72")
}

func LoadDatabase(file *ini.File) {
	sec := file.Section("database")
	Db = sec.Key("Db").MustString("mysql")
	DbHost = sec.Key("DbHost").MustString("a44447.com")
	DbPort = sec.Key("DbPort").MustString("3306")
	DbUser = sec.Key("DbUser").MustString("root")
	DbPassport = sec.Key("DbPassport").MustString("root")
	DbName = sec.Key("Dbname").MustString("myGoBlog")
}

func LoadQiLiu(file *ini.File) {
	sec := file.Section("QiLiu")
	AccessKey = sec.Key("AccessKey").String()
	SecretKey = sec.Key("SecretKey").String()
	Bucket = sec.Key("Bucket").String()
	QiLiuServer = sec.Key("QiLiuServer").String()
}
