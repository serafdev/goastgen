package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)
func Gen(){
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".go"){
				dump, err := Parse(path)
				if err != nil {
					log.Println(err)
				}
				WriteString(path + "ast", dump)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
