A simple web framework that allow you to create web pages and embed codes in <?go ?> tags.

### Install ###
**COMMAND**
```
$ go get code.google.com/p/gopages
```
**PACKAGE**
```
$ go get code.google.com/p/gopages/pkg
```
### Examples ###
```
$ go get code.google.com/p/gopages/examples
$ go get github.com/abiosoft/gopages-sample
```
### Documentation ###
http://go.pkgdoc.org/code.google.com/p/gopages/pkg
### Hello world example ###
  * create project folder 'hello'
  * create index.ghtml
```
 <html> 
 	<body> 
 	<?go print("<h1>Hello World</h1>") ?> 
 	</body> 
 </html> 
```
  * create hello.go
```
 package main 

 import (
    _ "hello/pages" //generated package	
    "code.google.com/p/gopages/pkg"
    "net/http"
 )
 
 func main(){
    http.HandleFunc("/", gopages.Handler("index.ghtml"))
    http.ListenAndServe(":9999", nil)
 }
```
  * build
```
$ gopages get
```
  * run
```
$ $GOPATH/bin/hello
```
  * finally, point your browser to localhost:9999