# CRUD of for tables
- user
- video
- comments
- session

### Session
- What is session?

  http is stateless, need to use session to keep the connect status.

- Use in this project

  Save the session id of a login, if the session id is nil, need to login again.

- The work flow of session

  ![](https://tva1.sinaimg.cn/large/007S8ZIlly1ge6e9x31uvj310a0n0thx.jpg)


## Middleware

Add middleware before handler, the whole work flow is like this:
```
main -> middleware -> def(message, err) -> handlers -> dbops -> response
```
### How to add the middleware ?

The interface is go and Java is different. In go is ducking type, if you have the same method signature, it seems to go that you have implements the interface.

We need to do something in the middle ware,like check session and do other auth things.
```go
type middleWareHandler struct {
	r *httprouter.Router
}

// factory method
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

// Implements of ServeHTTP method, so the middleware handler can response intercept the reqest.
// Each request need to check session and do auth things
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	m.r.ServeHTTP(w, r)
}

// main func file only put some defs, logic code should put in other files.
func main() {
	r := RegisterHandlers()
	// intercept each request
  middleWareHandler := NewMiddleWareHandler(r)

	http.ListenAndServe(":8000", middleWareHandler)
}
```