# Go Homework 6

This project hosts a server allowing people to place bids on items via http query parameters.

The server runs with go. 

## Usage

To add an item, use the following url:

http://localhost:3000?name=FruitSalad&description=Fruit&bidder=Dawn

To place a bid, type the following url:

http://localhost:3000?name=FruitSalad&bidder=Max&amt=600

To lookup an item, type the following:

http://localhost:3000?name=FruitSalad