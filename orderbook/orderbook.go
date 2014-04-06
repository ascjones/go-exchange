package orderbook

import (
	"time"
)

type Order struct {
	Pair     string
	Price    float64 // be careful with floating point arithmetic losing precision
	Quantity float64
}

type Trade struct {
	Pair      string
	Price     float64
	Quantity  float64
	Timestamp time.Time
}

type OrderBook struct {
	Bids, Asks []Order
	Trades     []Trade
}

func NewOrderBook() *OrderBook {
	return &OrderBook{make([]Order, 0), make([]Order, 0), make([]Trade, 0)}
}

func (book *OrderBook) PlaceBuyOrder(pair string, price float64, quantity float64) {
	for i := 0; i < len(book.Asks); i++ {
		sell := book.Asks[i]
		if sell.Price == price && sell.Quantity == quantity {
			// a = append(a[:i], a[i+1:]...)
			// remove the ask order
			book.Asks = append(book.Asks[:i], book.Asks[i+1])
			book.Trades = append(book.Trades, Trade{pair, price, quantity, time.Now()})
			return
		}
	}
	order := Order{pair, price, quantity}
	book.Bids = append(book.Bids, order)
}

func (book *OrderBook) PlaceSellOrder(pair string, price float64, quantity float64) {

}
