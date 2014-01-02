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

	Describe("Placing a buy order", func() {

		Context("with no matching sell orders", func() {

			BeforeEach(func() {
				book.PlaceBuyOrder("BTCUSD", 1000)
			})

			It("should create a new bid order", func() {
				Expect(book.Bids).To(HaveLen(1))
				Expect(book.Bids[0]).To(Equal(Order{"BTCUSD", 1000}))
			})
		})
	})

})
