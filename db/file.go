/*
 * @Descripttion: 文件写入mysql
 * @version: 7
 * @Author: KongJHong
 * @Date: 2019-08-02 19:58:33
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-03 11:19:37
 */
package db

import (
	mydb "filestore-server/db/mysql"
	"fmt"
	"database/sql"
)

//OnFileUploadFinish:文件上传完成之后，把文件插入到数据库表中
func OnFileUploadFinish(filehash string,filename string,
	filesize int64,fileaddr string) bool{
	//用预编译的语句Prepare()来写sql语句有很多好处
	//比如可以防止sql的注入攻击
	//同时对于批量插入有很好的效果
	//INSERT IGNORE 与INSERT INTO的区别就是INSERT IGNORE会忽略数据库中已经存在 的数据，
	//如果数据库没有数据，就插入新的数据，如果有数据的话就跳过这条数据。
	//这样就可以保留数据库中已经存在数据，达到在间隙中插入数据的目的。
	stmt,err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file(`file_sha1`,`file_name`,`file_size`,"+
		"`file_addr`,`status`) values(?,?,?,?,1)")

	if err != nil{
		fmt.Println("Failed to prepare statement,err:",err.Error())
		return false
	}

	defer stmt.Close()
	
	//sql语句执行
	ret,err := stmt.Exec(filehash,filename,filesize,fileaddr)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}


	if rf,err := ret.RowsAffected();nil == err{
		if rf <= 0{
			fmt.Printf("File with hash :%s has been uploaded before",filehash)
		}
		return true
	}

	return false
}


type TableFile struct{
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//GetFileMeta:从mysql获取文件元信息
func GetFileMeta(filehash string)(*TableFile,error){
	stmt,err := mydb.DBConn().Prepare(
		"select file_sha1,file_addr,file_name,file_size from tbl_file " +
		"where file_sha1=? and status=1")
	
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	defer stmt.Close()

	tfile := TableFile{}

	err = stmt.QueryRow(filehash).Scan(
		&tfile.FileHash,&tfile.FileAddr,&tfile.FileName,&tfile.FileSize)

	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}

	return &tfile,nil
}
