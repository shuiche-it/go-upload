# golang写的一个接收上传文件的服务器

### 前言
因为要从一个采集器里往外传一个文件,采集器只支持curl,所用就得又一个文件接收服务器了.

### 使用
将main.go文件直接复制到 gopath目录下,直接执行
~~~
go run main.go
~~~

curl的语句为:

~~~
 curl -F "action=upload" -F 'id=@/opt/dvr_rdk/ti816x/upgarade' http://192.168.2.181:9090/upload
~~~

其中 
- **192.168.2.181** 是服务器的地址.就是curl要发送的地址.
- **id** 为接收字段名
- **/opt/dvr_rdk/ti816x/upgarade** 为要发送的文件
