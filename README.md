# GO实战仿百度晕盘

项目来自慕课网课程：GO实战仿百度云盘

## 基础功能

- 服务架构：一个基本能用的文件上传服务
- 基本接口：基本功能接口（上传/下载/查询/删除）
- 逻辑代码演示：代码实操及功能流程演示

```go
//上传接口
http.HandleFunc("/file/upload",handler.UploadHandler)  
//上传成功返回页面         
http.HandleFunc("/file/upload/suc", handler.UploadSucHandler) 
//查询文件元信息  
http.HandleFunc("/file/meta", handler.GetFileMetaHandler)   
//查询多个最新的元信息 limit=?
http.HandleFunc("/file/query", handler.FileQueryHandler)
//下载接口
http.HandleFunc("/file/download", handler.DownloadHandler)
//文件元信息修改 op: 0表示修改 1表示其他 filehash: 文件hash值 filename:新名称
http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
//删除文件 filehash：文件hash
http.HandleFunc("/file/delete", handler.FileDeleteHandler)

```

