package main
import (
	"filestore-server/handler"
	"net/http"
	"fmt"
)
func main(){

	// 静态资源处理
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/file/upload",handler.UploadHandler)	//添加路由  按接口书写顺序看代码
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)

	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler) //   /user/signin由signin.html中的上传url返回
	http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))

	//http.HandleFunc("/user/info", handler.UserInfoHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		fmt.Printf("Failed to start server,err:%s\n",err.Error())
	}
}