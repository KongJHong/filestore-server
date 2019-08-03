/*
 * @Descripttion: 该类用于response返回，因为返回的数据很大，因此把返回的内容封装成JSON格式
 * @version: 10
 * @Author: KongJHong
 * @Date: 2019-08-03 16:10:41
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-03 17:01:47
 */
package util

import (
	"encoding/json"
	"fmt"
	"log"
)

//RespMsg http响应数据的通用结构
type RespMsg struct {
	Code int 			`json:"code"`			//错误码
	Msg  string 		`json:"msg"`		//提示信息
	Data interface{} 	`json:"data"`	//数据返回
}


//NewRespMsg 生成response对象
func NewRespMsg(code int,msg string,data interface{}) (*RespMsg){
	return &RespMsg{
		Code:code,
		Msg:msg,
		Data:data,
	}
}

//JSONBytes 对象转json格式的二进制数组
func (resp *RespMsg)JSONBytes() []byte{
	r,err := json.Marshal(resp)
	if err != nil{
		log.Println(err)
	}
	return r
}

//JSONString 对象转json格式的string
func (resp *RespMsg)JSONString() string{
	r,err := json.Marshal(resp)
	if err != nil{
		log.Println(err)
	}
	return string(r)
}

//GenSimpleRespString 只包含code和message的响应体(string)
func GenSimpleRespString(code int,msg string) string {
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code,msg)
}

//GenSimpleRespStream 只包含code和message的响应体([]byte])
func GenSimpleRespStream(code int,msg string) []byte {
	return []byte(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code,msg))
}