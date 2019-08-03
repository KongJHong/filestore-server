/*
 * @Descripttion: 维护db链接
 * @version: 6
 * @Author: KongJHong
 * @Date: 2019-08-02 19:40:28
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-03 15:06:56
 */
package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //导入这个库（为了让他自己初始化，并导入到sql中）是为了给database/sql库使用
	"fmt"
	"os"
	"log"
)

//全局唯一指定db句柄
var db *sql.DB

func init(){
	db,_ = sql.Open("mysql","root:123456@tcp(172.17.0.2:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)	//设置最大同时活跃数
	err := db.Ping()			//进行链接测试
	if err != nil{
		fmt.Println("Failed to connect to mysql,err",err.Error())
		os.Exit(1)
	}
}

//DBConn 返回数据库链接对象
func DBConn() *sql.DB{
	return db
}

//ParseRows 解析SQL Query返回，按一级目录划分，一级目录->二级目录
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}


func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}