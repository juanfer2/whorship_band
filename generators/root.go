package generators

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

type Data struct {
	ProductTarget   string
	ConcreteTargets []string
	Properties      []Property
}

type Property struct {
	Name     string
	TypeName string
}

func GenerateInterfaceTemplate(fileName string) {
	data := Data{
		ProductTarget:   "house",
		ConcreteTargets: []string{"normal", "igloo"},
		Properties: []Property{
			{"windowType", "string"},
			{"doorType", "string"},
			{"floor", "int"},
		},
	}

	processTemplate(utilities.GetRootFolder()+"/generators/templates/modules/domain.tmpl", "fileName.go", data)
	fmt.Println("Remember to edit the files that contain the Concrete Targets!")
}

func processTemplate(fileName string, outputFile string, data Data) {
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(fileName))
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}
	outputPath := "./tmp/" + outputFile
	fmt.Println("Writing file: ", outputPath)
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}

func rocessConcreteTargets(fileName string, data Data) {
	for _, value := range data.ConcreteTargets {
		newData := data
		newData.ConcreteTargets = []string{value}
		processTemplate(fileName, value+".go", newData)
	}
}
