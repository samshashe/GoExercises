package main

import (
	//"bytes"
	///"fmt"
	"html/template"
	"os"
)

type Article struct {
	Name       string
	AuthorName string
	Draft      bool
}

func main5() {
	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
		Draft:      false,
	}
	tmpl, err := template.New("Foo").Parse("{{.Name}}{{if .Draft}} (Draft){{else}} (Published){{end}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goArticle)
	if err != nil {
		panic(err)
	}

	tmpl2, err := template.New("Foo2").Parse(`
    {{range .}}
        <p>{{.Name}} by {{.AuthorName}}</p>
    {{else}}
        <p>No published articles yet</p>
    {{end}}
    `)
	if err != nil {
		panic(err)
	}
	err = tmpl2.Execute(os.Stdout, []Article{goArticle, {Name: "Truth", AuthorName: "Sam"}})
	if err != nil {
		panic(err)
	}
}

/*func main() {
	b := []byte{}
	var i byte
	for i = 33; i < 128; i++ {
		b = append(b, i)
	}

	//fmt.Println(string(b))

	//str := "abcdefg"
	//fmt.Println([]byte(str))

	var buf bytes.Buffer
	buf.Write([]byte("Hello "))
	count, err := fmt.Fprintf(&buf, "World!")
	fmt.Println(count, err)
	buf.WriteTo(os.Stdout)

}
*/
