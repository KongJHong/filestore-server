/*
 * @Descripttion: 对结构体数组FileMeta的自定义排序
 * @version:  5
 * @Author: KongJHong
 * @Date: 2019-08-02 14:17:01
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-02 15:38:17
 */

 package meta
 
 import "time"

 const baseFormat = "2006-01-02 15:04:05"

type ByUploadTime []FileMeta

func (s ByUploadTime) Len() int{
	return len(s)
}

func (s ByUploadTime) Less(i,j int) bool{

	iTime,_ := time.Parse(baseFormat, s[i].UploadAt)
	jTime,_ := time.Parse(baseFormat, s[j].UploadAt)
	
	return iTime.UnixNano() > jTime.UnixNano() 
}

func (s ByUploadTime) Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}