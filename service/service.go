package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	ht "github.com/v-grabko1999/go-html2json"
)

var k string
var v string

//convertion of html to json
func HtmlToJson(w http.ResponseWriter, r *http.Request) {
	//result := []map[string]interface{}{}
	resp, err := http.Get("https://www.autoblog.com/2016/03/11/p0110-obd-ii-trouble-code-intake-air-temperature-sensor-circui/") //passing url and returns http response
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body) //reads the http response and returns byte val
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.Println(string(html))
	//Conversion of html to json
	d, err := ht.New(strings.NewReader(string(html))) //pass the html into
	if err != nil {
		fmt.Println(err)
	}
	obj, err := d.ToJSON() //converts dom type to json,returns byte val
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(obj))

	/*json.Unmarshal(obj, &result)
	fmt.Println(len(result))
	for _, i := range result {

	}
	/*json.Unmarshal([]byte(v), &result)

	for _, i := range result {

		for k, v = range i {
			fmt.Println(k, v)

		}

	}*/
}
