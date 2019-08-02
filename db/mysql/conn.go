/*
 * @Descripttion: 维护db链接
 * @version: 6
 * @Author: KongJHong
 * @Date: 2019-08-02 19:40:28
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-02 20:04:16
 */
package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //导入这个库（为了让他自己初始化，并导入到sql中）是为了给database/sql库使用
	"fmt"
	"os"
)

//全局唯一指定db句柄
var db *sql.DB

func init(){
	db,_ = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)	//设置最大同时活跃数
	err := db.Ping()			//进行链接测试
	if err != nil{
		fmt.Println("Failed to connect to mysql,err",err.Error())
		os.Exit(1)
	}
}



//DBConn: 返回数据库链接对象
func DBConn() *sql.DB{
	return db
}


