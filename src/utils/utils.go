package utils

import (
	"bufio"
	"go/build"
	"io/ioutil"
	"log"
	"os"
)

func ReadTextByline(filePath string) ([]string){
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var keys []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keys = append(keys,scanner.Text() )
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return keys
}

func ReadTextAll(filePath string)(string){
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return string(b)
}

func GoPath()(string){
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

func WriteToFileOverride(filePath string ,context string) bool{
	d1 := []byte(context)
	if err := ioutil.WriteFile(filePath, d1, 0644);err!=nil{
		log.Fatal(err)
		return false
	}
	return true
}

