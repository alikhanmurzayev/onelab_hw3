package importsort

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func MyImportSort(path string) {
	//Reading file
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not find file: %s", err)
	}

	content := string(f)

	startImports := strings.Index(content,"(")
	endImports := strings.Index(content, ")")
	myImports := content[startImports+1 : endImports] //taking only imports
	myImports = strings.Trim(myImports, "\n")
	fields := strings.Fields(myImports) //putting imports in slice

	sort.Strings(fields)

	//Preparing to write sorted imports
	sortedImports := "(\n"
	for _, val := range fields {
		sortedImports += "\t" + val + "\n"
	}
	sortedImports += ")\n"

	content = content[ : startImports] + sortedImports + content[endImports+1: ]
	ioutil.WriteFile("testfile.go", []byte(content), 0644)
}