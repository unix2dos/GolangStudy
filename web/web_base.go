/*
scheme://host[:port#]/path/.../[?query-string][#anchor]
scheme         指定低层使用的协议(例如：http, https, ftp)
host           HTTP服务器的IP地址或者域名
port#          HTTP服务器的默认端口是80，这种情况下端口号可以省略。如果使用了别的端口，必须指明，例如 http://www.cnblogs.com:8080/
path           访问资源的路径
query-string   发送给http服务器的数据
anchor         锚
*/

/*
HTTP请求包, 请求头和请求体有个空行
GET /domains/example/ HTTP/1.1		//请求行: 请求方法 请求URI HTTP协议/协议版本
Host：www.iana.org				//服务端的主机名
User-Agent：Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.4 (KHTML, like Gecko) Chrome/22.0.1229.94 Safari/537.4			//浏览器信息
Accept：text/html,application/xhtml+xml,application/xml;q=0.9,*\/*;q=0.8	//客户端能接收的mine
Accept-Encoding：gzip,deflate,sdch		//是否支持流压缩
Accept-Charset：UTF-8,*;q=0.5		//客户端字符编码集
//空行,用于分割请求头和消息体
//消息体,请求资源参数,例如POST传递的参数
*/

/*
HTTP响应包

HTTP/1.1 200 OK						//状态行
Server: nginx/1.0.8					//服务器使用的WEB软件名及版本
Date:Date: Tue, 30 Oct 2012 04:14:25 GMT		//发送时间
Content-Type: text/html				//服务器发送信息的类型
Transfer-Encoding: chunked			//表示发送HTTP包是分段发的
Connection: keep-alive				//保持连接状态
Content-Length: 90					//主体内容长度
//空行 用来分割消息头和主体
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"... //消息体


1XX 提示信息 - 表示请求已被成功接收，继续处理
2XX 成功 - 表示请求已被成功接收，理解，接受
3XX 重定向 - 要完成请求必须进行更进一步的处理
4XX 客户端错误 - 请求有语法错误或请求无法实现
5XX 服务器端错误 - 服务器未能实现合法的请求
*/

/*
http 无状态协议, 所以需要通过cookie标识
Connection: keep-alive
从HTTP/1.1起，默认都开启了Keep-Alive保持连接特性，简单地说，当一个网页打开完成后，客户端和服务器之间用于传输HTTP数据的TCP连接不会关闭，如果客户端再次访问这个服务器上的网页，会继续使用这一条已经建立的TCP连接。
Keep-Alive不会永久保持连接，它有一个保持时间，可以在不同服务器软件（如Apache）中设置这个时间。



HTTP中的GET, POST, PUT, DELETE就对应着对这个资源的查，改，增，删4个操作。
POST是不具有幂等性
*/
