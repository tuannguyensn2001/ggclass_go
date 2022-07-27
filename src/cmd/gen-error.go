package cmd

import (
	"bufio"
	"fmt"
	"ggclass_go/src/app"
	"github.com/spf13/cobra"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type tmp struct {
	Fields []Parent
}

type Child struct {
	Name       string
	Message    string
	StatusCode int
	Code       string
}

type Parent struct {
	Name     string
	Children []Child
}

func genError() *cobra.Command {
	return &cobra.Command{
		Use: "gen-error",
		Run: func(cmd *cobra.Command, args []string) {

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Nhap category: ")
			parent, err := reader.ReadString('\n')
			parent = strings.TrimSuffix(parent, "\n")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Print("Nhap title: ")
			child, err := reader.ReadString('\n')
			child = strings.TrimSuffix(child, "\n")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Print("Nhap message: ")
			message, err := reader.ReadString('\n')
			message = strings.TrimSuffix(message, "\n")
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Print("Nhap status: ")
			statusCodeStr, err := reader.ReadString('\n')
			statusCodeStr = strings.TrimSuffix(statusCodeStr, "\n")
			if err != nil {
				log.Fatalln(err)
			}
			statusCode, err := strconv.Atoi(statusCodeStr)
			if err != nil {
				fmt.Println("status code is not valid", err)
			}

			m, err := app.LoadErr("./error.yml")
			if err != nil {
				fmt.Print("fail load err", err)
			}

			nextCode := 0
			for _, v := range m {
				nextCode += len(v)
			}
			nextCode++

			httpError := app.HttpError{
				StatusCode: statusCode,
				Message:    message,
				Code:       strconv.Itoa(nextCode),
			}

			_, ok := m[parent]
			if !ok {
				m[parent] = make(map[string]app.HttpError)
			}

			m[parent][child] = httpError

			file := "error.yml"

			f, err := os.Create(file)
			if err != nil {
				fmt.Println("err create file", err)
			}

			parentFields := make([]Parent, 0)

			for k, v := range m {
				item := Parent{
					Name: k,
				}
				children := make([]Child, 0)
				for k1, v1 := range v {
					item2 := Child{
						Name:       k1,
						Message:    v1.Message,
						StatusCode: v1.StatusCode,
						Code:       v1.Code,
					}
					children = append(children, item2)
				}
				item.Children = children
				parentFields = append(parentFields, item)
			}

			templ := template.Must(template.ParseFiles("src/templates/error.tmpl"))
			templ.Execute(f, tmp{
				Fields: parentFields,
			})
			f.Close()

			//deleteLine(file)
		},
	}
}

func deleteLine(path string) {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.TrimSuffix(string(read), "\n"))
	ioutil.WriteFile(path, []byte(read), fs.ModePerm)
}
