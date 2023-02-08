package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main(){

	name, numOfRepo, err := githubInfo("jonddon")
	if err != nil {
		log.Fatalf("Error: %s", err)
		
	}
	fmt.Printf("%s has %d repositories\n", name, numOfRepo);
	// resp, err := http.Get("http://api.github.com/users/jonddon")
	// if err != nil {
	// 	log.Fatalf("Error: %s", err)
	// }
	// if resp.StatusCode != http.StatusOK {
	// 	log.Fatalf("Error: %s", resp.Status)
	// }
	// fmt.Printf("Content-Type %s\n", resp.Header.Get("Content-Type"))
	// //_ is an unused variable
	// //io.Copy is to copy resp.Body into the console output
	// // if _,err := io.Copy(os.Stdout, resp.Body); err != nil{
	// // 	log.Fatalf("Error: can't copy %s", err)
	// // } 
	// var r Reply
	// dec:= json.NewDecoder(resp.Body)
	// if err :=dec.Decode(&r); err != nil{
	// 	log.Fatalf("error: can't copy %s", err)
	// }
	// fmt.Println(r)
	// fmt.Printf("%#v", r)
}

 func githubInfo(login string) (string, int, error){
	 //url.PathEscape to remove space from loginbif there is any
	 url := "http://api.github.com/users/" + url.PathEscape(login)
	 resp, err := http.Get(url)
	 if err != nil{
		 log.Fatalf("error: can't get github info: %s", err)
		 return "", 0, err
	 }

	 if resp.StatusCode != http.StatusOK{
		 log.Fatalf("error: %s", resp.Status)
		 return "", 0, fmt.Errorf(resp.Status)
	 }

	 fmt.Printf("Content-Type %s\n", resp.Header.Get("Content-Type"))

	//  var r Reply
	// anonymous struct is a struct without s type 
	// and used mainly for one of task like  api call
	 var r struct { // anonymous struct
		Name string
		Public_Repos int
		Public_Gists int
		User_type string `json:"type"`
	}
	 dec:=json.NewDecoder(resp.Body)
	 if err:=dec.Decode(&r); err != nil{
		 return "", 0, err
	 }

	//  fmt.Printf("%#v", r)
	 return r.Name, r.Public_Repos, nil

 }

// in returning vakues you can use a different name 
// but you need to tag the field with the json name
// type Reply struct{
// 	Name string
// 	Public_Repos int
// 	Public_Gists int
// 	User_type string `json:"type"`
// }
/* JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int64, int, uint8, ...
array <-> []any ([]interface{})
object <-> map[string]any, struct
encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/