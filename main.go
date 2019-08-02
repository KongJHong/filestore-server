package main
import (
	"filestore-server/handler"
	"net/http"
	"fmt"
)
func main(){
	http.HandleFunc("/file/upload",handler.UploadHandler)	//添加路由  按接口书写顺序看代码
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		fmt.Printf("Failed to start server,err:%s\n",err.Error())
	}
}