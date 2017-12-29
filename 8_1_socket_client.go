/*
### 通信函数
func (c *TCPConn) Write(b []byte) (int, error)
func (c *TCPConn) Read(b []byte) (int, error)

### 地址信息
type TCPAddr struct {
	IP IP
	Port int
	Zone string // IPv6 scoped addressing zone
}

### 获取addr
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only), TCP(IPv6-only)或者TCP(IPv4, IPv6的任意一个)。
addr表示域名或者IP地址，例如"www.google.com:80" 或者"127.0.0.1:22"。


### 客户端连接函数
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
net参数是"tcp4"、"tcp6"、"tcp"中的任意一个，分别表示TCP(IPv4-only)、TCP(IPv6-only)或者TCP(IPv4,IPv6的任意一个)
laddr表示本机地址，一般设置为nil
raddr表示远程的服务地址

### 服务器监听函数
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
func (l *TCPListener) Accept() (Conn, error)


### 控制连接函数

设置建立连接的超时时间，客户端和服务器端都适用，当超过设置时间时，连接自动关闭。
func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)



用来设置写入/读取一个连接的超时时间。当超过设置时间时，连接自动关闭。
func (c *TCPConn) SetReadDeadline(t time.Time) error
func (c *TCPConn) SetWriteDeadline(t time.Time) error



设置keepAlive属性，是操作系统层在tcp上没有数据和ACK的时候，会间隔性的发送keepalive包，操作系统可以通过该包来判断一个tcp连接是否已经断开，在windows上默认2个小时没有收到数据和keepalive包的时候人为tcp连接已经断开，这个功能和我们通常在应用层加的心跳包的功能类似。
func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
