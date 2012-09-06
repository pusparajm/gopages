// Copyright 2010 Abiola Ibrahim <abiola89@gmail.com>. All rights reserved.
// Use of this source code is governed by New BSD License
// http://www.opensource.org/licenses/bsd-license.php
// The content and logo is governed by Creative Commons Attribution 3.0
// The mascott is a property of Go governed by Creative Commons Attribution 3.0
// http://creativecommons.org/licenses/by/3.0/

package gopages

import "net/http"

var ParsedPages = make(map[string]func(http.ResponseWriter, *http.Request))

// Handler to use with http.HandleFunc. page is the path to gopages file
func Handler(page string) func(http.ResponseWriter, *http.Request) {
	return ParsedPages[page]
}
