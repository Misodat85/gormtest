package config

import (
	"fmt"
)

//DBに接続するする際のTokenはどんなメンバか
type dataBaseConfig struct {
	User string
	Pass string
	IP   string
	Port string
	Name string
}

var c dataBaseConfig

const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

func init() {
	c = dataBaseConfig{
		User: "root",//mysqlのユーザー名
		Pass: "p2hack",//password
		IP:   "127.0.0.1",//IPアドレス
		Port: "3306",//ポート番号
		Name: "test2",//テーブル名
	}

	err := checkElements(c)
	if err != nil {
		// TODO: Faild to get env Value
	}
}

func checkElements(c dataBaseConfig) error {
	if c.User == "" {
		return fmt.Errorf("User error")
	}
	if c.Pass == "" {
		return fmt.Errorf("invalid password")
	}
	if c.IP == "" {
		return fmt.Errorf("")
	}
	if c.Port == "" {
		return fmt.Errorf("")
	}
	if c.Name == "" {
		return fmt.Errorf("")
	}
	return nil
}

func GetConnectionToken() string {
	return fmt.Sprintf(accessTokenTemplate, c.User, c.Pass, c.IP, c.Port, c.Name)
}
