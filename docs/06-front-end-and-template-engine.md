# Front End and Template Engine
- Including UI(js+html) plus back end code.


## Go template engine

> Template engine is a tool for parsing html fils and replacing the element, to generate the final pages.

There are tow kind of template in golang:
1. `text/template`
2. `html/template`

Go templates using the dynamic-generating mode.

The whole flow:
1. Parsing static html files to generate the templates
2. Render the templates and dynamic element to the final pages

## Writing front end codes

```
.
├── client.go       // proxy, send requests to api module
├── defs.go         // def
├── handlers.go     // handle request
├── main.go         // main
└── template        // folder for static files
```

### 前端两种代理转发方式

跨域和转发请求

Two main modes of proxy:
1. Proxy mode
2. API mode

> **cross-domain** in web 浏览器

跨域是危险的

### 大前端

- web server + templates(static files and js)