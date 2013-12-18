package orderbook_test

import (
	. "github.com/ascjones/go-exchange/orderbook"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Orderbook", func() {
	var book *OrderBook

	BeforeEach(func() {
		book = NewOrderBook()
	})

	It("should be empty of orders", func() {
		Expect(book.Bids).To(HaveLen(0))
		Expect(book.Asks).To(HaveLen(0))
	})
})
