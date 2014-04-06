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
				book.PlaceBuyOrder("BTCUSD", 1000, 1)
			})

			It("should create a new bid order", func() {
				Expect(book.Bids).To(HaveLen(1))
				Expect(book.Bids[0]).To(Equal(Order{"BTCUSD", 1000, 1}))
			})
		})

		Context("with a fully matching ask order", func() {

			BeforeEach(func() {
				book.PlaceSellOrder("BTCUSD", 1000, 1)
				book.PlaceBuyOrder("BTCUSD", 1000, 1)
			})

			It("doesn't add a buy order", func() {
				Expect(book.Bids).To(HaveLen(0))
			})

			It("removes the matching ask order", func() {
				Expect(book.Asks).To(HaveLen(0))
			})

			It("should create a new executed trade", func() {
				Expect(book.Trades).To(HaveLen(1))
			})
		})
	})

})
