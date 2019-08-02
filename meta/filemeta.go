/*
 * @Descripttion: 保存文件的元信息
 * @version: 序号-4
 * @Author: KongJHong
 * @Date: 2019-08-02 10:27:28
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-02 10:40:21
 */


package meta

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

//UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta){
	fileMetas[fmeta.FileSha1] = fmeta
}

//GetFileMeta: 通过sha1获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta{
	return fileMetas[fileSha1]
}