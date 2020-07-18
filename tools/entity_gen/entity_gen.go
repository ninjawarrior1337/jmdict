package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"text/template"
)

var entityTemplate = template.Must(template.New("entity_map").Parse(`// Code generaed by entity map generator; DO NOT EDIT.
package jmdict

var entity = map[string]string{
	{{- range $i, $e := .}}
	"{{$i}}": "{{$e}}",
	{{- end}}
}
`))

var entityRegex = regexp.MustCompile(`<!ENTITY (.+) "(.+)">`)

//var endingMap = map[string]string{
//	"u&#39": "う",
//	"ku&#39": "く",
//	"su&#39": "す",
//	"tsu&#39": "つ",
//	"nu&#39": "ぬ",
//	"hu/fu&#39": "ふ",
//	"mu&#39": "む",
//	"ru&#39": "る",
//	"dzu&#39": "づ",
//}

func main() {
	in := flag.String("i", "JMdict_e", "input path for the JMdict file")
	out := flag.String("o", "entity.go", "output path for the go file containing entity map")
	flag.Parse()

	var entityMap = map[string]string{}

	inF, err := ioutil.ReadFile(*in)
	if err != nil {
		log.Fatalln(err)
	}
	matches := entityRegex.FindAllStringSubmatch(string(inF), -1)

	for _, match := range matches {
		entityMap[match[1]] = match[2]
	}

	outF, err := os.OpenFile(*out, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	defer outF.Close()
	outF.Truncate(0)
	entityTemplate.Execute(outF, entityMap)
}
