package main

import (
	"fmt"

	"io/ioutil"

	"net/http"
)

// Grabweb getting the content from the web page
//skipping errors
func main() {
	// Blank identifier **
	res, _ := http.Get("http://192.168.1.1")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}


package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main()
{
 res,_ := http.Get("url")
 page,_:= ioutil.ReadAll(res.Body)
 res.Body.Close()
 fmt.Printf("%s",page)
 
}
