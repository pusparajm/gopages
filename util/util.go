// Copyright 2010 Abiola Ibrahim <abiola89@gmail.com>. All rights reserved.
// Use of this source code is governed by New BSD License
// http://www.opensource.org/licenses/bsd-license.php
// The content and logo is governed by Creative Commons Attribution 3.0
// The mascott is a property of Go governed by Creative Commons Attribution 3.0
// http://creativecommons.org/licenses/by/3.0/

package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//StringBuilder type like StringBuilder in java
type StringBuilder struct {
	String string
}

//Creates an instance of StringBuilder
func NewStringBuilder(s string) *StringBuilder {
	sBuilder := new(StringBuilder)
	sBuilder.String = s
	return sBuilder
}

//Retrieves the string Content
func (this *StringBuilder) Content() string {
	return this.String
}

//Appends string to the end of the string content
func (this *StringBuilder) Append(s string) {
	this.String = this.String + s
}

//Deletes string from start index to end
func (this *StringBuilder) Delete(start, end int) {
	part1 := this.String[0:start]
	part2 := this.String[end:]
	this.String = part1 + part2
}

//Deletes the remaining string from index start
func (this *StringBuilder) DeleteTillEnd(start int) {
	this.String = this.String[0:start]
}

//Empties the string content
func (this *StringBuilder) Reset() {
	this.String = ""
}

//Returns the index of the first occurence of a particular string
func (this *StringBuilder) Index(s string) int {
	return strings.Index(this.Content(), s)
}

//Returns the length of the string content
func (this *StringBuilder) Len() int {
	return len(this.Content())
}

//Returns the substring from start to end index
func (this *StringBuilder) Sub(start, end int) string {
	return this.Content()[start:end]
}

//Returns the remaining string from the start index
func (this *StringBuilder) SubEnd(start int) string {
	return this.Content()[start:]
}

//QuoteParser to parse quotes e.g. { } or <?go ?>
type QuoteParser struct {
	buffer, static   *StringBuilder
	outer, inner     []string
	opening, closing string
}

//Creates a new QuoteParser with string s, opening and closing string
func NewQuoteParser(s, opening, closing string) *QuoteParser {
	parser := new(QuoteParser)
	parser.buffer, parser.static = NewStringBuilder(s), NewStringBuilder(s)
	parser.opening, parser.closing = opening, closing
	//parser.inner, parser.outer = new(vector.StringVector), new(vector.StringVector)
	return parser
}

//Parses the string content in it
func (this *QuoteParser) Parse() (err error) {
	for this.HasNext() {
		_, _, err = this.Next()
		if err != nil {
			return
		}
	}
	_, _, err = this.Next()
	return
}

//Returns the array of contents embedded in the quotes
func (this *QuoteParser) Parsed() []string {
	return this.inner
}

//Returns the array of contents outside the quotes
func (this *QuoteParser) Outer() []string {
	return this.outer
}

//Parses the next set and returns the embedded and outer strings with an error if any
//This method deletes the parsed string from the content
//If there is still need to parse whole content, use Reset()
func (this *QuoteParser) Next() (inner, outer string, err error) {
	start := this.buffer.Index(this.opening)
	if start >= 0 {
		start += len(this.opening)
	}
	end := -1
	if start >= 0 {
		end = strings.Index(this.buffer.SubEnd(start), this.closing)
		if end >= 0 {
			end += start
		}
	}
	//	end := this.buffer.Index(this.closing)
	if end < 0 && start >= 0 {
		err = errors.New("no closing string '" + this.closing + "' found near " + this.buffer.SubEnd(start-len(this.opening)))
		return
	}
	if start >= end && end > -1 {
		err = errors.New("no matching closing string '" + this.closing + "' found near " + this.buffer.SubEnd(start-len(this.opening)))
		return
	}
	if this.HasNext() {
		inner = this.buffer.Sub(start, end)
	}
	//if len(this.opening) > 4 {
	//	println(start, end, this.closing, inner)
	//}
	//println(this.opening)
	if this.buffer.Len() > 0 {
		l := this.buffer.Index(this.opening)
		if l < 0 {
			l = this.buffer.Len()
		}
		outer = this.buffer.Sub(0, l)
		if start >= 0 {
			l = end + len(this.closing)
		}
		this.buffer.Delete(0, l)
		this.inner = append(this.inner, inner)
		this.outer = append(this.outer, outer)
	}
	return
}

//Resets the content to its state before parsing
func (this *QuoteParser) Reset() {
	this.buffer = this.static
	this.outer = []string{}
	this.inner = []string{}
}

//Checks whether there is next set of data to parse
func (this *QuoteParser) HasNext() (res bool) {
	start := this.buffer.Index(this.opening)
	end := -1
	if start >= 0 {
		end = strings.Index(this.buffer.SubEnd(start), this.closing)
		if end >= 0 {
			end += start
		}
	}
	//end := this.buffer.Index(this.closing)
	return (start >= 0 && end >= 0 && start < end)
}

//Returns the remaining content in the buffer being used
func (this *QuoteParser) String() string {
	return this.buffer.Content()
}

//public variable to store settings
var Config map[string][]string

const (
	SETTINGS   = "pages.json"
	PATHS_FILE = "pages/.pages"
)

//Settings file type
type Settings struct {
	Data map[string][]string
}

//loads settings from pages.settings
func LoadSettings() (s *Settings, err error) {
	s = new(Settings)
	err = s.parse()
	return
}

//parse the informations in the settings file
func (this *Settings) parse() (err error) {
	settings, err := ioutil.ReadFile(SETTINGS)
	config := make(map[string]string)
	if err == nil {
		err = json.Unmarshal(settings, &config)
		if err != nil {
			return
		}
	}
	this.Data = make(map[string][]string)
	if _, ok := config["extensions"]; !ok {
		this.Data["extensions"] = []string{"ghtml"}
	} else {
		this.Data["extensions"] = strings.Split(config["extensions"], " ")
	}
	if _, ok := config["folders"]; !ok {
		this.Data["folders"] = []string{"."}
	} else {
		this.Data["folders"] = strings.Split(config["folders"], " ")
	}
	err = this.GeneratePages()
	return
}

//generates all .go source files
func (this *Settings) GeneratePages() (err error) {
	if len(this.Data["extensions"]) == 0 {
		return
	}
	var pages []string
	for i := 0; i < len(this.Data["folders"]); i++ {
		pages, err = this.iterFiles(this.Data["folders"][i], pages)
		if err != nil {
			break
			return
		}
	}
	this.Data["pages"] = pages
	//if len(pages) > 0 {
	//	AddHandlers(pages)
	//}
	return
}

//loops through root and subfolders to locate files with
//extensions specified in settings file
func (this *Settings) iterFiles(f string, pages []string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	stat, er := file.Stat()
	if er != nil {
		err = er
		return nil, err
	}
	if stat.IsDir() {
		dirs, err := file.Readdir(-1)
		if err != nil {
			return nil, err
		}
		for _, d := range dirs {
			pages, err = this.iterFiles(path.Join(file.Name(), d.Name()), pages)
		}
	} else {
		if hasExt(file.Name(), this.Data["extensions"]) {
			err = generate(file.Name())
			if err != nil {
				println(err.Error())
				return nil, err
			}
			pages = append(pages, file.Name())
		}
		file.Close()
	}
	return pages, err
}

//check if the file has the extension
func hasExt(filename string, ext []string) bool {
	extn := path.Ext(filename)
	for _, e := range ext {
		if "."+e == extn {
			return true
		}
	}
	return false
}

//direct function to generate .go source file
func generate(page string) (err error) {
	p, err := NewPage(page)
	if err != nil {
		return
	}
	err = p.ParseToFile()
	return
}
