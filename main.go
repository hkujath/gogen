package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
)

var (
	flagTemplateFile = flag.String("f", "", "template file")
	flagTypeName     = flag.String("t", "", "data type")
)

func init() {
	flag.Parse()
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Arguments missing")
		fmt.Println("Function call is: gogen [template file] [type]")
		os.Exit(1)
	}

	templateFileName := os.Args[1]
	typeName := os.Args[2]

	// Parse template file
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		fmt.Printf("Parsing of file %s failed.\n%s", templateFileName, err)
		os.Exit(1)
	}

	// Generate the new output file
	outName := fmt.Sprintf("gogen_%s_gen.go", typeName)
	fd, err := os.OpenFile(outName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Couldn't open out file %s.\n%s", outName, err)
		os.Exit(1)
	}
	defer fd.Close()

	// Define base template and write template with datatype to file
	data := struct {
		T string
	}{
		typeName,
	}
	t.Execute(fd, data)

	//if *flagTemplateFile == ""{
	//	fmt.Println("Template file missing")
	//	os.Exit(1)
	//}
	//
	//if *flagTypeName == ""{
	//	fmt.Println("Type to generate is missing missing")
	//	os.Exit(1)
	//}
	//

}
