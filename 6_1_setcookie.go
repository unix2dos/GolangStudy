/*
	type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

// MaxAge=0 means no 'Max-Age' attribute specified.
// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int  //过期时间
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
	//设置Cookie
	http.SetCookie(w ResponseWriter, cookie *Cookie)

	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)


	//读取Cookie
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, cookie)

	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie.Name)
	}


}

*/