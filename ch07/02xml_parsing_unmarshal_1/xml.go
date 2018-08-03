// リスト7.2
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type A struct {
	XMLName xml.Name `xml:"B"`
	Posts   []Post   `xml:"posts>post"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Post struct {
	Id      string `xml:"id, attr"`
	Content string `xml:"content"`
	Author1 Author `xml:"author1"`
	Author2 Author `xml:"author2"`
}

func main() {
	xmlFile, err := os.Open(os.Getenv("GOPATH") + "/src/github.com/Saya-K/golang-sample/ch07/02xml_parsing_unmarshal_1/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post A
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
