package enchere

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//
//func OpenDatabase() (*sql.driver.Conn) {
//	log.Println("Entering OpenDatabase")
//
//}
//
//func WriteDataAuction(auction chan *Auctions) {
//	log.Println("Write auctions")
//}
//
// We create a channel
func WriteData(bid chan *Bids) {
	log.Println("Entering WriteDate")
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `
	create table bids (id INTEGER PRIMARY KEY AUTOINCREMENT, Timestamp  text, AuctionID integer ,State integer, Timeout integer, DateBid text, DateTimeout text, HighestBidder text, HighestBid text) ;
	create table auctions (id INTEGER PRIMARY KEY AUTOINCREMENT, Title text, AuctionID integer); 
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	for {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare("insert into bids(AuctionID, Timestamp, State, Timeout, DateBid, DateTimeout, HighestBidder, HighestBid) values(?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		log.Println("WriteData: Reading Channel")
		mybid := <-bid
		timestamp := mybid.Response.Reference.Timestamp
		log.Println("WriteData: ", timestamp)
		for _, mybid := range mybid.Response.Items {
			//log.Println("WriteData: Inserting auction:", mybid.AuctionID)

			_, err = stmt.Exec(mybid.AuctionID, timestamp, mybid.State, mybid.Timeout, mybid.DateBid, mybid.DateTimeout, mybid.HighestBidder, mybid.HighestBid)
			if err != nil {
				log.Fatal(err)
			}
		}
		tx.Commit()
	}
}
