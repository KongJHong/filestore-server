/*
 * @Descripttion:  根据文件流和文件句柄计算hash值
 * @version: 序号-3
 * @Author: KongJHong
 * @Date: 2019-08-02 09:45:01
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-02 10:26:54
 */


package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"path/filepath"
	"fmt"
)

type Sha1Stream struct{
	_sha1 hash.Hash
}

func (obj *Sha1Stream) Updata(data []byte){
	if obj._sha1 == nil{
		obj._sha1 = sha1.New()
	}
	obj._sha1.Write(data)
}

func (obj *Sha1Stream) Sum() string{
	return hex.EncodeToString(obj._sha1.Sum([]byte("")))
}

func Sha1(data []byte) string{
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))//加密两次
}

func FileSha1(file *os.File)string{
	_sha1 := sha1.New()
	io.Copy(_sha1,file)
	return hex.EncodeToString(_sha1.Sum(nil)) //加密两次
}

func MD5(data []byte)string{
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

//返回文件的md5
func FileMD5(file *os.File)string{
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

/*
 * 路径文件是否存在
 */
func PathExists(path string) (bool,error){
	_,err := os.Stat(path)
	if err == nil{
		return true,nil
	}

	if os.IsNotExist(err){
		return false,nil
	}

	return false,err
}

//获取文件大小
func GetFileSize(filename string) int64{
	var result int64
	filepath.Walk(filename,func(path string,f os.FileInfo,err error)error{
		result = f.Size()
		return nil
	})

	stat,_ := os.Stat(filename)
	tmp := stat.Size()
	if(result == tmp){
		fmt.Println("GetFileSize方法可改")
	}

	return result
}

