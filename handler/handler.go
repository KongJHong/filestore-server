/*
 * @Descripttion: 路由器主业务 GET-返回文件 POST-保存文件到FileMeta,以及hash保存
 * @version: 序号-2
 * @Author: KongJHong
 * @Date: 2019-08-02 09:39:03
 * @LastEditors: KongJHong
 * @LastEditTime: 2019-08-02 11:33:51
 */

package handler

import (
	"io/ioutil"
	"net/http"
	"io"
	"os"
	"fmt"
	"time"
	"filestore-server/meta"	//文件元信息
	"filestore-server/util"
	"encoding/json"
)

/*
 * UploadHandler： 处理文件上传
 * 1.获取上传页面
 * 2.选取本地文件,form形式上传文件
 * 3.云端接受文件流，写入本地存储
 * 4.云端更新文件元信息集合
 */
func UploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		//返回上传html页面
		data,err := ioutil.ReadFile("./static/view/index.html")
		if err!=nil{
			io.WriteString(w,"internel server error")
			return
		}
		io.WriteString(w,string(data))
	}else if r.Method == "POST"{
		//接收文件流及存储到本地目录

		//1.获取上传文件，file:文件句柄 head:文件头部信息
		file,head,err := r.FormFile("file")
		if err != nil{
			fmt.Printf("Failed to get data,err:%s\n",err.Error())
			return 
		}
		defer file.Close()


		//2. 创建文件元信息fileMeta用于保存文件，方便hash
		fileMeta := meta.FileMeta{
			FileName : head.Filename,
			Location : "/tmp/"+head.Filename,
			UploadAt : time.Now().Format("2006-01-02 15:04:05"),
		}


		//3. 本地创建文件句柄——未做任何操作
		newFile,err := os.Create(fileMeta.Location)
		if err != nil{
			fmt.Printf("Failed to create file,err :%s\n",err.Error())
			return 
		}
		defer newFile.Close()

		//4. 把文件复制到本地——io.Copy改变了读取下标，后面要用Seek改回来
		fileMeta.FileSize,err = io.Copy(newFile, file)
		if err != nil{
			fmt.Printf("Failed to save data into file,err:%s\n", err.Error())
			return 
		}

		//5. 修改文件读取下标，返回文件的sha1 hash值，添加到文件元信息对应管理结构
		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)


		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}


//uploadSucHandler:上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Upload finished")
}

//GetFileMetaHandler:查询文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data,err := json.Marshal(fMeta)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}