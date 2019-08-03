# GO实战仿百度云盘

项目来自慕课网课程：GO实战仿百度云盘

## 基础功能

- 服务架构：一个基本能用的文件上传服务
- 基本接口：基本功能接口（上传/下载/查询/删除）
- 逻辑代码演示：代码实操及功能流程演示

```
————filestore-server
 |
 |---db : 数据库相关函数
 |    |-- mysql
 |    |     |--conn.go : 管理mysql链接对象，返回数据库单例
 |    |-- file.go : 数据库API函数，调用conn.go的单例
 |
 |
 |---doc : 一些重要文档资料
 |    |-- table.sql : 数据库表创建语句，方便查询
 |    |
 |    |
 |
 |
 |---handler : 路由器映射API
 |    |-- handler.go : 路由器映射API，程序主逻辑
 |    |
 |    |
 |
 |---meta : 文件元数据相关
 |    |-- filemeta.go : 文件元数据，辅助handler.go的API进行对文件的增删改查操作
 |    |-- sort.go     : 辅助文件元数据排序
 |    |
 |    
 |---static : 保存静态文件的文件夹
 |    |-- view : *.html文件，用于浏览器返回
 |
 |
 |---util : 工具文件夹
 |    |-- util.go : 加密，路径合法化判断，文件大小的一些可共用的函数
 |
 |
```


```go
main.go
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

[配置docker](https://blog.csdn.net/bingzhongdehuoyan/article/details/79411479)

[docker配置Mysql主从同步](https://www.cnblogs.com/songwenjie/p/9371422.html)

![](https://kongjhong-image.oss-cn-beijing.aliyuncs.com/img/{330FF055-031D-3284-7E2D-A859739E4E17}.jpg)

单点模式发送故障时对整个系统的影响很大，因此主从模式是相当优秀的解决方案

**表字段说明**

![](https://kongjhong-image.oss-cn-beijing.aliyuncs.com/img/{948868D9-1C0C-AC58-0B87-D050CB2B0F5E}.jpg)

**使用MySQL小结**

- 通过sql.DB来管理数据库链接对象
- 通过sql.Open来创建协程安全的sql.DB对象
- 优先使用`Prepared Statement`

**本章小结**

1. MySQL的特点与应用场景
2. 主从架构与文件表设计逻辑
3. Golang与MySQL的亲密接触

