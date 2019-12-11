package data

import(
	"strings"
	"bufio"
    "fmt"
    "os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Misodat85/gormtest/model"
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


//TODO: リマインドする日時と、内容を入力して、model.go定義した形で返す
func DataGenerator() (date model.User){
	var data model.User

	fmt.Print("日付を入力してください：")
	data.DATE = StrStdin()
	fmt.Print("内容を入力してください:")
	data.TODO = StrStdin()
	return data 
}

func StrStdin() (stringInput string) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()//
    stringInput = scanner.Text()//バッファ内の文字列を返す
    stringInput = strings.TrimSpace(stringInput)
    return
}

func IntStdin() (int, error) {
    stringInput := StrStdin()
    return strconv.Atoi(strings.TrimSpace(stringInput))
}