/*
 * @Descripttion: 用于用户存储进入mysql的类
 * @version: 9
 * @Author: KongJHong
 * @Date: 2019-08-03 14:18:13
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-03 21:11:02
 */
package db

import(
	mydb "filestore-server/db/mysql"
	"fmt"
)

//UserSignup 通过用户名及密码完成user表的注册操作
func UserSignup(username ,passwd string) bool{
	stmt,err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user(`user_name`,`user_pwd`) values(?,?)")

	if err != nil{
		fmt.Println("Failed to insert user,err:",err.Error())
		return false;
	}

	defer stmt.Close()

	ret,err := stmt.Exec(username,passwd)
	if err != nil{
		fmt.Println("Failed to insert user,err",err.Error())
		return false
	}

	if rowsAffected,err := ret.RowsAffected();nil == err && rowsAffected > 0 {
		return true
	}

	fmt.Println("重复注册")
	return false
}

//UserSignin 判断密码是否一致
func UserSignin(username,encpwd string) bool{
	stmt,err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")

	if err != nil{
		fmt.Println(err.Error())
		return false
	}

	rows,err := stmt.Query(username)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}else if rows == nil{
		fmt.Println("username not found:" + username)
		return false
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) >0 && string(pRows[0]["user_pwd"].([]byte))==encpwd{
		return true
	}
	
	return false
}

//UpdateToken 刷新用户登录的token
func UpdateToken(username,token string)bool{
	stmt,err := mydb.DBConn().Prepare(
		"replace into tbl_user_token(`user_name`,`user_token`) values(?,?)")

	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_,err = stmt.Exec(username,token)

	if err != nil{
		fmt.Println(err.Error())
		return false
	}

	return true

}

//GetUserToken: 查询制定用户的token值
func GetUserToken(username string) (string){
	stmt,err := mydb.DBConn().Prepare(
		"select `user_token` from table tbl_user_token where user_name=? limit 1")
	if err != nil{
		fmt.Println(err.Error())
		return ""
	}
	
	defer stmt.Close()

	rows,err := stmt.Query(username)
	if err != nil{
		fmt.Println(err.Error())
		return ""
	}else if rows == nil{
		fmt.Println("username not found:" + username)
		return ""
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) >0 {
		return string(pRows[0]["user_token"].([]byte))
	}
	
	return ""
}

//User 用于用户信息查询返回的结构体
type User struct{
	Username 	string
	Email 		string
	Phone 		string
	SignupAt 	string
	LastActiveAt string
	Status 		int
}

//GetUserInfo 用户信息查询
func GetUserInfo(username string) (User,error){
	user := User{}

	stmt,err := mydb.DBConn().Prepare(
		"select user_name,signup_at from tbl_user where user_name=? limit 1")

	if err != nil{
		fmt.Println(err.Error())
		return user,err
	}

	defer stmt.Close()

	//执行查询的操作
	err = stmt.QueryRow(username).Scan(&user.Username,&user.SignupAt)
	if err != nil{
		return user,err
	}

	return user,nil
}

