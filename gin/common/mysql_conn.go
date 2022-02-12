package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func Mysql() {
	// DSN:Data Source Name
	dsn := "root:@tcp(192.168.56.13:3306)/auto_deploy"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()  // 注意这行代码要写在上面err判断的下面
	fmt.Println("mysql is connected!")

}
