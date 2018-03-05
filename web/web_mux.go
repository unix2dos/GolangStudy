/*
ServeMux

type ServeMux struct {
	mu sync.RWMutex   //锁，由于请求涉及到并发处理，因此这里需要一个锁机制
	m  map[string]muxEntry  // 路由规则，一个string对应一个mux实体，这里的string就是注册的路由表达式
	hosts bool // 是否在任意的规则中带有host信息
}


type Handler interface {
	ServeHTTP(ResponseWriter, *Request)  // 路由实现器
}
*/
package main
import "fmt"
import "net/http"

type myMux struct{

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func (p *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
	if (r.URL.Path == "/"){
		sayhelloName(w,r)
		return
	}
	http.NotFound(w,r)
}


func main(){
	mux := &myMux{}
	http.ListenAndServe(":8080", mux)
}