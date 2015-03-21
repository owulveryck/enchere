package enchere

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func FullUpdateAuctionsList(client *http.Client, auctions chan *Auctions, baseurl *string) {
	log.Println("UpdateAution: Entering the FullUpdateAuctionList function")
	var localAuctions *Auctions
	// Entering an infinite loop
	for {
		fullUrl := []string{"http://", *baseurl, "/json/site/load/current/refresh/"}
		response, err := client.Get(strings.Join(fullUrl, ""))
		if err != nil {
			panic(err)
		}
		log.Println("FullUpdateAuction: update page: ", response.StatusCode)

		body, err := ioutil.ReadAll(response.Body)
		// And display it!
		//fmt.Print(string(body))

		// Now let's play with Json

		err = json.Unmarshal([]byte(body), &localAuctions)
		if err != nil {
			log.Println(err)
		} else {
			//log.Println("UpdateAuction: feeding the channel")
			auctions <- localAuctions
		}
		// Now sleep for 20 seconds
		time.Sleep(20 * time.Second)
	}
}
func UpdateAuctionsList(client *http.Client, bids chan *Bids, baseurl *string) {
	var localBids *Bids
	log.Println("UpdateAution: Entering the updateAuctionList function")
	// Entering an infinite loop
	for {
		fullUrl := []string{"http://", *baseurl, "/json/auction/update/"}
		response, err := client.Get(strings.Join(fullUrl, ""))
		if err != nil {
			panic(err)
		}
		log.Println("UpdateAuction: update page: ", response.StatusCode)

		body, err := ioutil.ReadAll(response.Body)
		// And display it!
		//fmt.Print(string(body))

		err = json.Unmarshal([]byte(body), &localBids)
		if err != nil {
			log.Println(err)
		} else {
			// log.Println("UpdateAuction: feeding the channel")
			bids <- localBids
		}
		// Now sleep for 2 seconds
		//log.Println("UpdateAuction: sleeping")
		time.Sleep(2 * time.Second)
	}
}
