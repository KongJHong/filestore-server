# GO实战仿百度云盘

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

## 服务架构变迁
![](https://kongjhong-image.oss-cn-beijing.aliyuncs.com/img/{507E0415-6D1E-C59C-EA71-57DB581B41DD}.jpg)


### 部署MySQL主从模式

![](https://kongjhong-image.oss-cn-beijing.aliyuncs.com/img/{330FF055-031D-3284-7E2D-A859739E4E17}.jpg)

单点模式发送故障时对整个系统的影响很大，因此主从模式是相当优秀的解决方案

**表字段说明**

![](https://kongjhong-image.oss-cn-beijing.aliyuncs.com/img/{948868D9-1C0C-AC58-0B87-D050CB2B0F5E}.jpg)
