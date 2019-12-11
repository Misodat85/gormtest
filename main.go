package main

import(
	//"time"
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/Misodat85/gormtest/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Misodat85/gormtest/model"
	"github.com/Misodat85/gormtest/data"
)
/* TODO:
ルーティングをせずgormのmethodを用いてデータベースとの接続とテーブルの作成を行う。
テーブルを作成後、テーブルをデータベースから取得してに全て表示する。

SignupHandler内の処理に倣い、json形式のstringAutoMigrationになげる？
func (ctrl *IsController)SignupHandler(context *gin.Context) {
	var user, data model.User

	ctrl.DB.Create(&user)
	context.JSON(200, gin.H{"msg": ""})
}
*/

func main(){
	//var data model.User
	//データを作り、DBに接続しテーブルを作成する。
	db, err := initializeDataBase()
	if err != nil {
		fmt.Printf("Faild connecting mysql")
	}
	
	
	if db !=nil{
		fmt.Printf("TODO List Registered")
	}
	db.Close()
	/*TODO: 
	さきほど初期化した同じテーブルに複数回データを登録する。
	*/
	
}

func initializeDataBase()(*gorm.DB, error){
	var db *gorm.DB
	//var data data.Data
	var err error
	// var count time.Duration
	token := config.GetConnectionToken()//GetConnectionTokenが出力した文字列をtokenに入れる
	
	//DBへの接続を複数回試みる処理
	db,err = gorm.Open("mysql",token)//tokenを用いてDBに接続をする
	if err == nil{
		db.AutoMigrate(&model.User{})//modelで構造体として定義したメンバに合わせてDBにテーブルを作成する。
		return db,nil
	}
	log.Fatal(err)
	return nil, err
}
func modeChange(i int){
	switch i {
		case 0:
			//TODO:データベースに日付とやることを登録する
			data := data.DataGenerator()
			db.Create(&data)
		case 1:
			//TODO:データベースからある日付のTODOListを引っ張ってくる
			
		case 2:
			//TODO:内容を更新する
	}
}
/*
//TODO:データ登録をする
func registerDataBase()(*gorm.DB,error){
	var db *gorm.DB
	var err error
	db.Create(&model.User{})
}

*/