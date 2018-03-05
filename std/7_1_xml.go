/*
	解析:  func Unmarshal(data []byte, v interface{}) error


<?xml version="1.0" encoding="utf-8"?>
<servers version="2">
	<server>
		<serverName>Shanghai_VPN</serverName>
		<serverIP>127.0.0.1</serverIP>
	</server>
	<server>
		<serverName>Beijing_VPN</serverName>
		<serverIP>127.0.0.2</serverIP>
	</server>
</servers>

1. 如果struct的一个字段是string或者[]byte类型且它的tag含有",innerxml"，Unmarshal将会将此字段所对应的元素内所有内嵌的原始xml累加到此字段上
2. 如果struct中有一个叫做XMLName，且类型为xml.Name字段，那么在解析的时候就会保存这个element的名字到该字段,如上面例子中的servers。
3. 如果某个struct字段的tag定义中含有XML结构中element的名称，那么解析的时候就会把相应的element值赋值给该字段，如上servername和serverip定义。
4. 如果某个struct字段的tag定义了中含有",attr"，那么解析的时候就会将该结构所对应的element的与字段同名的属性的值赋值给该字段，如上version定义。
5. 如果某个struct字段的tag定义 型如"a>b>c",则解析的时候，会将xml结构a下面的b下面的c元素的值赋值给该字段。
6. 如果某个struct字段的tag定义了"-",那么不会为该字段解析匹配任何xml数据。
7. 如果struct字段后面的tag定义了",any"，如果他的子元素在不满足其他的规则的时候就会匹配到这个字段。
8. 如果某个XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上，这个字段的类型可能是[]byte或string,如果没有这样的字段存在，那么注释将会被抛弃。
9. 为了正确解析，go语言的xml包要求struct定义中的所有字段必须是可导出的（即首字母大写）
*/

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
	// Description string   `xml:",innerxml"`
}
type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("/Users/liuwei/golang/src/daydayup/7_1.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	v := Servers{}
	err = xml.Unmarshal(bytes, &v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
