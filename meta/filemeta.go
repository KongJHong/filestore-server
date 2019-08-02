/*
 * @Descripttion: 保存文件的元信息
 * @version: 序号-4
 * @Author: KongJHong
 * @Date: 2019-08-02 10:27:28
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-02 20:22:49
 */


package meta

import (
	"sort"
	mydb "filestore-server/db"
)


//FileMeta: 文件元信息结构
type FileMeta struct{
	FileSha1 string		//sha1 hash值
	FileName string		//文件名
	FileSize int64		//文件大小
	Location string		//文件位置
	UploadAt string		//上传的时间戳
}


var fileMetas map[string]FileMeta //通过FileSha1唯一对应FileMeta信息

func init(){
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMeta:新增/更新文件元信息 （弃用）
func UpdateFileMeta(fmeta FileMeta){
	fileMetas[fmeta.FileSha1] = fmeta
}

// 8-2_ 20:18
//UpdateFileMetaDB:新增/更新文件信息到mysql中
//替换原来的UpdateFileMeta函数
func UpdateFileMetaDB(fmeta FileMeta) bool{
	return mydb.OnFileUploadFinish(
		fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

//GetFileMeta: 通过sha1获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta{
	return fileMetas[fileSha1]
}


//GetLastFileMetas:获取批量的文件元信息列表
func GetLastFileMetas(count int) []FileMeta{
	fMetaArray := make([]FileMeta,len(fileMetas))
	for _,v := range fileMetas{
		fMetaArray = append(fMetaArray,v)
	}
	sort.Sort(ByUploadTime(fMetaArray)) //自定义fMetaArray排序
	return fMetaArray[:count]
}

//RemoveFileMeta:删除元信息
func RemoveFileMeta(fileSha1 string){
	delete(fileMetas, fileSha1)
}