//Generated by gopages from src/hello.ghtml, do not edit
//This file will be overwritten during build

package pages

import (
	"code.google.com/p/gopages/pkg"
	"fmt"
	"net/http"
	"time"
)

func Rendersrchelloghtml(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	writer.Header().Set("Content-Type", "text/html")
	print := func(toPrint ...interface{}) {
		fmt.Fprint(writer, toPrint...)
	}
	formValue := func(keyToGet string) string {
		return request.FormValue(keyToGet)
	}
	print("")
	formValue("") // prevent initialization runtime error

	fmt.Fprint(writer, `
<html>
	<head><title>`)
	print("Hello with gopages")
	fmt.Fprint(writer, `</title>
	<body>
		<a href="echo">Echo example</a><br>
		<img src="gopages.png" />
		`)

	for i := 1; i < 5; i++ {

		fmt.Fprint(writer, `
			<h`)
		print(i)
		fmt.Fprint(writer, `>Hello gopages</h`)
		print(i)
		fmt.Fprint(writer, `>
		`)

	}

	fmt.Fprint(writer, `
		<hr>
		`)

	print("page generated on " + time.Now().String())

	fmt.Fprint(writer, `
	</body>
</html>
`)

}
func init() {
	gopages.ParsedPaths["pages/srchelloghtml.go"] = "src/hello.ghtml"
}
