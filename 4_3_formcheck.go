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


*/

package main
import "fmt"

func main (){
	fmt.Println("13333333333")
	fmt.Println("13333333333")
	fmt.Println("13333333333")
	fmt.Println("13333333333")
	fmt.Println("13333333333")
}