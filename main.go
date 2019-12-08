package main

import(
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/Misodat85/gormtest/config"
	"github.com/Misodat85/gormtest/model"
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

	//データを作り、DBに接続しテーブルを作成する。
	db, err := initializeDataBase()
	if err != nil {
		fmt.printf("Faild connecting mysql")
	}
	/*TODO: 
	さきほど初期化した同じテーブルに複数回データを登録する。
	*/
}

func initializeDataBase()(*gorm.DB, error){
	var db *gorm.DB
	//var data data.Data
	var err error
	var count time.Duration
	token := config.GetCo@nnectionToken()//GetConnectionTokenが出力した文字列をtokenに入れる
	user := User{
		ID:0001,
		NAME:"takuya",
		EMAIL:"sample1@sample.com"
		PASSWORD:"passwd"
	}
	//DBへの接続を複数回試みる処理
	count = 1
	for {
		if count > 5 {
			return nil, fmt.Errorf("")
		}
		db,err = gorm.Open("mysql",token)//tokenを用いてDBに接続をする
		if err == nil{
			db.AutoMigrate(&user)//modelで構造体として定義したメンバに合わせてDBにテーブルを作成する。
			return db,nil
		}
		time.Sleep(3 * time.Second)
		count++
	}
	return nil, err
}
