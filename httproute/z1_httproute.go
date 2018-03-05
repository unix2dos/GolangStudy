package main

/*

httprouter的使用

### Named parameters
Pattern: /user/:user
 /user/gordon              match
 /user/you                 match
 /user/gordon/profile      no match
 /user/                    no match

### Catch-All parameters
Pattern: /src/*filepath
 /src/                     match
 /src/somefile.go          match
 /src/subdir/somefile.go   match

*/
import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

type HostSwitch map[string]http.Handler

func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403) // Or Redirect?
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	// log.Fatal(http.ListenAndServe(":8080", router))

	hs := make(HostSwitch)
	hs["localhost:12345"] = router

	// Use the HostSwitch to listen and serve on port 12345
	log.Fatal(http.ListenAndServe(":12345", hs))
}
