package data

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/Misodat85/gormtest/model"
)

/*
TODO:
データベースのテーブルに登録するキーとバリューを
上に定義した構造体の通りに
複数個のデータ生成する

*/
type Data struct {
	DB	*gorm.DB
}
/*

func dataGenerator(data *Data){
	var user, data model.User
	data.DB.Create(&user)

}

{
	"id": "000001",
	"name": "takuya",
	"email": "sample@example.com",
	"password": "hogehoge"
}

*/