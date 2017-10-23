package util

import (
	"io/ioutil"
	"path/filepath"
	"os/exec"
	"fmt"
	"log"
	"os"
)


func Mkdir(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	return path
}

func Save_bytes_as(bytes []byte, filename string) {
	err := ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	//log.Println("saved as: "+ filename)
}

func Read_bytes_from(filename string) []byte {
	log.Println("reading from: "+ filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes
}

func List_dir(wildcard string) []string{
	files, _ := filepath.Glob(wildcard)
	return files
}

func File_exists(filename string) bool{
	_, err := os.Stat(filename);
	if os.IsNotExist(err) {
		return false
	}else{
		return true
	}
}

func Exec_command(cmd string, args []string) []byte{
	log.Println(cmd)
	log.Println(args)
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))
	return out
}
