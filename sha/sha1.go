package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	
}

// cat http.log.gz| gunzip |sha1sum
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
	r, err := gzip.NewReader(file)
	if err != nil{
		return "", err
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