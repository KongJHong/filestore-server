package main
import (
	"filestore-server/handler"
	"net/http"
	"fmt"
)
func main(){
	http.HandleFunc("/file/upload",handler.UploadHandler)	//添加路由
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("file/meta", handler.GetFileMetaHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Printf("Failed to start server,err:%s\n",err.Error())
	}
}