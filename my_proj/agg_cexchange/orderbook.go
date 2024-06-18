package main

type Order struct {
	Size      float64
	Limit     *Limit
	Timestamp int64
}

type Limit struct {
	Price       float64
	Orders      []*Orderbook
	TotalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

type Orderbook struct {
	Asks []*Limit
	Bids []*Limit
}
