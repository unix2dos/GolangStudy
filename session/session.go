package session

import (
	"fmt"
	"sync"
)

type Session struct {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime)
}

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

var providers = make(map[string]Provider)
func Register(name string, provider Provider){
	if provider == nil {
		panic("is null")
	}
	if _, have := providers[name]; have{
		panic("have")
	}
	providers[name] = provider
}

func NewManager(providerName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("error %q", providerName)
	}
	return &Manager{cookiename: cookieName, provider: provider, maxLifeTime: maxLifeTime}, nil
}


func (manager* Manager) sessionId() string{
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil{
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil  || cookie.Value == ""{
		sid := sessionId()
		session, _ := manager.provider.SessionInit(sid)
		cookie := http.Cookie{name: manager.cookieName, Value:url.QueryEscape(sid), Path:"/",HttpOnly:true,MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.value)
		session, _ := manager.provider.SessionRead(sid)
	}
	return
}

func (manger *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
 cookie, err := r.Cookie(manger.cookieName)
 if err != nil || cookie.Value == ""{
	 return
 }else{
	 manger.lock.Lock()
	 defer manger.lock.Unlock()
	 manger.provider.SessionDestroy(cookie.Value)
	 expriation := time.Now()
	 cookie := http.Cookie{Name:manger.cookieName,Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
	http.SetCookie(w, &cookie)
	}
}

func (manager *Manager) GC(){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func(){manager.GC()})
}