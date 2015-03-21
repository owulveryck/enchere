package enchere

// ************************ JSON <=> Structure asignation ***********************************

// Every 20s
type Auctions struct {
	Cmd      string `json:"cmd"`
	Version  string `json:"version"`
	Status   int    `json:"status"`
	Response struct {
		Reference struct {
			ImageBase string  `json:"image_base"`
			SmsCost   float32 `json:"sms_cost"`
			SmsMsisdn string  `json:"sms_msisdn"`
			Timestamp string  `json:"timestamp"`
		} `json:"reference"`
		Items []struct {
			Title          string  `json:"title"`            //"Le nouveau lot d'Apple"
			Type           int     `json:"type"`             // 1
			AuctionID      int     `json:"auction_id"`       // 3261773
			PrototypeID    int     `json:"prototype_id"`     // 20351
			ProductID      int     `json:"product_id"`       // 5630
			ProductGroupID int     `json:"product_group_id"` // 0
			IntlAuctionID  int     `json:"intl_auction_id"`  // 4304690
			CategoryID     string  `json:"category_id"`      // "898"
			Rrp            string  `json:"rrp"`              // "1815.99"
			PoolID         string  `json:"pool_id"`          // "3"
			ShippingCosts  float32 `json:"shipping_costs"`   // 15
			SmsKeyword     string  `json:"sms_keyword"`      // "FGTT"
			AuctionData    struct {
				State       int      `json:"state"`        // 3
				Timeout     int      `json:"timeout"`      // 5
				CreditCost  int      `json:"credit_cost"`  // 8
				DateOpens   string   `json:"date_opens"`   // 2015-03-14 17:00:00
				Flags       []string `json:"flags"`        // "intl"
				DateTimeout string   `json:"date_timeout"` // 2015-03-16 22:10:52
				LastBid     struct {
					DateBid       string  `json:"date_bid"`       // "2015-03-16 22:10:47"
					HighestBidder string  `json:"highest_bidder"` // "Steffan86"
					HighestBid    float32 `json:"highest_bid"`    // 289.36
				} `json:"last_bid"`
				Availability struct {
					Type      int    `json:"type"`       // 1
					TimeStart string `json:"time_start"` // "09:00"
					TimeEnd   string `json:"time_end"`   // "03:00"
				} `json:"availability"`
				BiddingHistory []struct {
					UserName string  `json:"user_name"`
					BidValue float32 `json:"bid_value"`
				} `json:"bidding_history"`
			} `json:"auction_data"`
			Images     []string `json:"images"`
			BuynowData struct {
				ItemsLeft     int     `json:"items_left"`     // 10000001
				BuynowMessage string  `json:"buynow_message"` // "Les prix chutent quand vous enchérissez !"
				BasePrice     float32 `json:"base_price"`     // 1815.99
			} `json:"buynow_data"`
			Segments []int `json:"segments"`
		} `json:"items"`
	} `json:"response"`
}

// Every 2s
type Bids struct {
	Cmd      string `json:"cmd"`
	Version  string `json:"version"`
	Status   int    `json:"status"`
	Response struct {
		Items []struct {
			AuctionID     int     `json:"auction_id"`
			State         int     `json:"state"`
			Timeout       int     `json:"timeout"`
			DateBid       string  `json:"date_bid"`
			DateTimeout   string  `json:"date_timeout"`
			HighestBidder string  `json:"highest_bidder"`
			HighestBid    float32 `json:"highest_bid"`
		} `json:"items"`
		Reference struct {
			Timestamp string `json:"timestamp"`
		} `json:"reference"`
	} `json:"response"`
}
