package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type item struct {
	name        string
	description string
	minimumBid  float64
	bestBid     float64
	bidder      string
}

// Keeps track of all the items.
var items = make([]item, 0)

// Handles adding an item to the service
func add(w http.ResponseWriter, req *http.Request) {
	var name = req.URL.Query().Get("name")
	var description = req.URL.Query().Get("description")
	var bidder = req.URL.Query().Get("bidder")
	var minimumBid, err1 = strconv.ParseFloat(req.URL.Query().Get("min"), 64)
	var bestBid, err2 = strconv.ParseFloat(req.URL.Query().Get("best"), 64)
	if err1 != nil || err2 != nil {
		fmt.Fprintf(w, "invaild number\n")
		return

	}
	var i = item{
		name:        name,
		description: description,
		minimumBid:  minimumBid,
		bestBid:     bestBid,
		bidder:      bidder,
	}
	items = append(items, i)
	fmt.Fprintf(w, "added %s to the auction\n", name)
}

// Handles placing a bid, either accepting or rejecting the bid
func bid(w http.ResponseWriter, req *http.Request) {

	// Gets the values from the url
	var name = req.URL.Query().Get("name")
	var bidder = req.URL.Query().Get("bidder")
	var amount, err1 = strconv.ParseFloat(req.URL.Query().Get("amt"), 64)

	if err1 != nil {
		fmt.Fprintf(w, "invaild number\n")
		return

	}

	// Looks for the item by going through each
	for i := 0; i < len(items); i++ {
		if name == items[i].name {
			//If the item's bid is not greater than the best bid, reject
			if items[i].bestBid >= amount {
				fmt.Fprintf(w, "could not place bid, your bid of %f is not more than the current bid of %f\n", amount, items[i].bestBid)
				return

			}
			// If the item's bid is not at least the minimum bid, reject
			if items[i].minimumBid >= amount {
				fmt.Fprintf(w, "could not place bid, your bid of %f is not more than the minimum bid of %f\n", amount, items[i].minimumBid)
				return

			}

			// If we get passed the previous if statements, then set this as the best bid
			items[i].bestBid = amount
			items[i].bidder = bidder
			fmt.Fprintf(w, "ok")
			return
		}

	}

	fmt.Fprintf(w, "could not find item named %s\n", name)
}

// Handles looking up an item. Responds with the item if found, otherwise responds by saying not found.
func lookup(w http.ResponseWriter, req *http.Request) {
	//Get the name from the url
	var name = req.URL.Query().Get("name")

	// Search each item for an item with the same name
	for i := 0; i < len(items); i++ {
		if name == items[i].name {
			// If we find the item, display it.
			fmt.Fprintf(w, "name: %s\ndescription: %s\nminimum bid: %f\nbest bid: %f\nbidder: %s\n", items[i].name, items[i].description, items[i].minimumBid, items[i].bestBid, items[i].bidder)
			return

		}

	}

	//If we didn't find the item, display that we didn't find it
	fmt.Fprintf(w, "could not place bid: could not find item named %s\n", name)
}

func main() {
	// Sets up each handler
	http.HandleFunc("/add", add)
	http.HandleFunc("/bid", bid)
	http.HandleFunc("/lookup", lookup)
	fmt.Println("Listening at http://localhost:3000")

	// Starts the service on port 3000
	http.ListenAndServe(":3000", nil)

}
