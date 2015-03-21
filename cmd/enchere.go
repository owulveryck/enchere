package main

import (
	"flag"
	"fmt"
	"github.com/owulveryck/enchere"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url" // Package url parses URLs and implements query escaping. See RFC 3986.
	"strings"
)

func main() {

	username := flag.String("user", "foo", "username to connect with")
	password := flag.String("password", "bar", "password to connect with")
	baseurl := flag.String("baseurl", "www.url.com", "the base url")
	flag.Parse()
	// Cookies
	jar, _ := cookiejar.New(nil)
	// Instanciate a client  http://golang.org/pkg/net/http/
	client := &http.Client{Jar: jar}
	// Get the login page http://golang.org/pkg/net/http/#Get
	// the get function returns:
	// resp *Response, err error
	// Values maps a string key to a list of values. It is typically used for query parameters and form values
	formData := url.Values{"user_name": {*username}, "password": {*password}}

	// DO perform the login !
	// PostForm issues a POST to the specified URL, with data's keys and values urlencoded as the request body.
	fullUrl := []string{"http://", *baseurl, "/login/"}
	response, err := client.PostForm(strings.Join(fullUrl, ""), formData)
	//  The client must close the response body when finished with it:
	// A defer statement pushes a function call onto a list.
	// The list of saved calls is executed after the surrounding function returns.
	// Defer is commonly used to simplify functions that perform various clean-up actions.
	// Therefore the response is closed at the exit of the current function (main)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	log.Println("Login page: ", response.StatusCode)

	// fmt.Println(jsonData)
	// Now Display all Earth.Data.List.Nom et Earth.Data.List.Valeur
	var bids chan *enchere.Bids = make(chan *enchere.Bids)
	var auctions chan *enchere.Auctions = make(chan *enchere.Auctions)
	// Now we launch the update in another thread
	go enchere.FullUpdateAuctionsList(client, auctions, baseurl)
	go enchere.UpdateAuctionsList(client, bids, baseurl)
	go enchere.WriteData(bids)
	var input string
	fmt.Scanln(&input)
}
