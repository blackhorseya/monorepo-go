package model

import (
	"time"
)

// OrderType is an enum for the type of order.
type OrderType int

const (
	OrderTypeLimit OrderType = iota
	OrderTypeMarket
)

var orderTypeMap = map[OrderType]string{
	OrderTypeLimit:  "limit",
	OrderTypeMarket: "market",
}

// String returns the string representation of the order type.
func (x *OrderType) String() string {
	return orderTypeMap[*x]
}

// SideType is an enum for the side of the order.
type SideType int

const (
	SideTypeBuy SideType = iota
	SideTypeSell
)

var sideTypeMap = map[SideType]string{
	SideTypeBuy:  "buy",
	SideTypeSell: "sell",
}

// String returns the string representation of the side type.
func (x *SideType) String() string {
	return sideTypeMap[*x]
}

// SubmitOrder is an entity for submitting an order.
type SubmitOrder struct {
	ID string `json:"id"`

	Symbol string    `json:"symbol"`
	Side   SideType  `json:"side"`
	Type   OrderType `json:"type"`

	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type OrderStatus int

const (
	// OrderStatusNew is a status for a new order.
	OrderStatusNew OrderStatus = iota

	// OrderStatusFilled is a status for a filled order.
	OrderStatusFilled

	// OrderStatusPartiallyFilled is a status for a partially filled order.
	OrderStatusPartiallyFilled

	// OrderStatusRejected is a status for a rejected order.
	OrderStatusRejected

	// OrderStatusCanceled is a status for a canceled order.
	OrderStatusCanceled
)

var orderStatusMap = map[OrderStatus]string{
	OrderStatusNew:             "new",
	OrderStatusFilled:          "filled",
	OrderStatusPartiallyFilled: "partially_filled",
	OrderStatusRejected:        "rejected",
	OrderStatusCanceled:        "canceled",
}

func (x *OrderStatus) String() string {
	return orderStatusMap[*x]
}

func (x *OrderStatus) Closed() bool {
	return *x == OrderStatusFilled || *x == OrderStatusRejected || *x == OrderStatusCanceled
}

// Order is an entity for an order.
type Order struct {
	SubmitOrder

	Status    OrderStatus `json:"status"`
	IsWorking bool        `json:"is_working"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
