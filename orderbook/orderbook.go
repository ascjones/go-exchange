package orderbook

type Order struct{}

type OrderBook struct {
	Bids, Asks []Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{make([]Order, 0), make([]Order, 0)}
}
