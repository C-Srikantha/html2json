package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	ht "github.com/v-grabko1999/go-html2json"
	"golang.org/x/net/html/atom"
)

var result []byte
var v string

type Htmlread struct {
	Name    string
	Text    string
	Element []map[string]interface{}
}
type Detail struct {
	General_meaning    string
	Causes             string
	Symptoms           string
	Mechanic_diagnosis string
	Severity_level     string
	Suggested_repairs  string
}

var det Detail
var id = "p0110"

//convertion of html to json
func HtmlToJson(w http.ResponseWriter, r *http.Request) {
	var hold []Htmlread

	//var count int = 0
	var str []string

	resp, err := http.Get("https://www.autoblog.com/2016/03/11/" + id + "-obd-ii-trouble-code-intake-air-temperature-sensor-circui/") //passing url and returns http response
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
	d, err := ht.New(strings.NewReader(string(html))) //pass the html var
	if err != nil {
		fmt.Println(err)
	}
	obj, err := d.ToJSON() //converts dom type to json,returns byte val
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, string(obj))

	obj1, _ := d.ByClass("post-body")
	for _, i := range obj1 {

		obj2, _ := i.ByTag(atom.P)
		for index, j := range obj2 {
			if index != 0 {
				result, _ = j.ToJSON()
				json.Unmarshal(result, &hold)
				//fmt.Println(hold[0].Element)
				str = append(str, hold[0].Text)

			}
		}
	}

	det.General_meaning = str[0]
	det.Causes = str[1] + "," + str[2] + "," + str[3] + "," + str[4] + "," + str[6] + "," + str[7]
	det.Symptoms = str[7]
	det.Mechanic_diagnosis = str[8]
	det.Severity_level = str[11] + "," + str[12]
	det.Suggested_repairs = str[13] + "," + str[14] + "," + str[15]
	fmt.Println(det.General_meaning)
	/*obj1, _ := d.ByTag(atom.H3)

	for _, J := range obj1 {

		obj2, _ := d.ByTag(atom.P)
		for index, j := range obj2 {
			if index != 0 {
				result, _ = j.ToJSON()
				json.Unmarshal(result, &hold)
				fmt.Println(string(result))

			}
		}
	}*/
	writefile()
}

func writefile() {
	file, _ := os.Create("dtc1.csv")
	defer file.Close()
	writefile := csv.NewWriter(file)
	defer writefile.Flush()
	row := []string{id, det.General_meaning, det.Causes, det.Symptoms, det.Mechanic_diagnosis, det.Severity_level, det.Suggested_repairs}
	_ = writefile.Write(row)

}
