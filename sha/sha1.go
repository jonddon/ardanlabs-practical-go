package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1sum("http.log.gz")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("%s\n", sig)

	sig, err = sha1sum("sha1.go")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("%s\n", sig)
}

/*
if file names ends with .gz
	$ cat http.log.gz| gunzip | sha1sum
else
	$ cat http.log.gz| sha1sum
*/
func sha1sum(fileName string ) (string, error){
	//os.Open opens the file in the filesystem
	file, err := os.Open(fileName)
	if err != nil{
		return "", err
	}
	//defer is used to manage resources e.g Database connections, VM connections
	// defer only happens when a function exit
	// defers are called in LIFO order
	/*	defer file.Close() is used to close file because 
	we don't want too many open files that we exhaust 
	the number of files our server can handle
	*/
	defer file.Close()
	//create a file that can be read(unzips the file)
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}
	w := sha1.New()

	//copy the uncompressed file r to w 
	if _, err := io.Copy(w, r); err != nil{
		return "", err
	}
	
	//Get the signature of the uncompressed file.
	//it returns a Hash
	//returns the checksum of the data
	sig:=w.Sum(nil)

	//fmt.Sprintf("%x", sig) converts hash into string
	return fmt.Sprintf("%x", sig), nil
}