package orderbook

type Order struct {
	Pair  string
	Price float64 // be careful with floating point arithmetic losing precision
}

type OrderBook struct {
	Bids, Asks []Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{make([]Order, 0), make([]Order, 0)}
}

func (book *OrderBook) PlaceBuyOrder(pair string, price float64) {
	order := Order{pair, price}
	book.Bids = append(book.Bids, order)
}
