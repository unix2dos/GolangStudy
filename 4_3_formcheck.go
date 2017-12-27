/*
必填字段
if len(r.Form["username"][0])==0  获取map
r.Form.Get() 获取单个

数字
getint,err:=strconv.Atoi(r.Form.Get("age"))
if err!=nil{}

if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
	return false
}

中文
unicode 包提供的 func Is(rangeTab *RangeTable, r rune) bool

if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
	return false
}

英文
if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
	return false
}

电子邮件地址
if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
	fmt.Println("no")
}else{
	fmt.Println("yes")
}

手机号码
if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
	return false
}


身份证号码
//验证15位身份证，15位的是全部数字
if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
	return false
}
//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
	return false
}


日期和时间
t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
fmt.Printf("Go launched at %s\n", t.Local())




####### 预防跨站脚本JavaScript、VBScript、 ActiveX或Flash以欺骗用户. 对XSS最佳的防护应该结合以下两种方法：
一是验证所有输入数据，有效检测攻击;
另一个是对所有输出数据进行适当的处理，以防止任何已成功注入的脚本在浏览器端运行

func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
func HTMLEscapeString(s string) string //转义s之后返回结果字符串
func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串


但是如果想正常输出未转义
import "text/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")

输出
Hello, <script>alert('you have been pwned')</script>!


import "html/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", template.HTML("<script>alert('you have been pwned')</script>")) //没有template.HTML会被转义掉
输出
Hello, <script>alert('you have been pwned')</script>!



####### 防止多次递交表单

<input type="hidden" name="token" value="{{.}}">
设置一个hidden get的时候设置时间戳设置到页面并且存到session
post的进行验证


 ####### 上传文件

 要使表单能够上传文件，首先第一步就是要添加form的enctype属性，enctype属性有如下三种情况:
1. application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
2. multipart/form-data	  不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
3. text/plain	  空格转换为 "+" 加号，但不对特殊字符编码。



<form enctype="multipart/form-data" action="/upload" method="post">
  <input type="file" name="uploadfile" />
  <input type="hidden" name="token" value="{{.}}"/>
  <input type="submit" value="upload" />
</form>



通过下面的实例我们可以看到我们上传文件主要三步处理：

1. 表单中增加enctype="multipart/form-data"
2. 服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
3. 使用r.FormFile获取文件句柄，然后对文件进行存储等处理。



*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)

		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return
		}
		defer f.Close()
		io.Copy(f, file)

	}
}
func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}
