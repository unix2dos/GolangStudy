/*

 */

package main

import (
	"fmt"
	"net/http"
	"time"
)

var globalSession *session.Manager

func init() {
	globalSession, _ := NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		fmt.Fprintf(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form("username"))
		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if createtime.(int64)+360 < (time.Now().Unix()) {
		globalSession.SessionDestroy(w, r)
		sess = globalSession.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", ct.(int)+1)
	}
	fmt.Fprintf(w, sess.Get("countnum"))
}

func main() {

}
