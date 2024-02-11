package model

import (
	"encoding/json"
	"time"
)

// StockQuota is a value object that represents the stock quota.
type StockQuota struct {
	open      float64
	high      float64
	close     float64
	low       float64
	volume    int
	value     float64
	updatedAt time.Time
}

// NewStockQuota creates a new StockQuota.
func NewStockQuota(
	open float64,
	high float64,
	closePrice float64,
	low float64,
	volume int,
	updatedAt time.Time,
) StockQuota {
	return StockQuota{
		open:      open,
		high:      high,
		close:     closePrice,
		low:       low,
		volume:    volume,
		value:     closePrice * float64(volume),
		updatedAt: updatedAt,
	}
}

func (x *StockQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Open   float64 `json:"open,omitempty"`
		High   float64 `json:"high,omitempty"`
		Close  float64 `json:"close,omitempty"`
		Low    float64 `json:"low,omitempty"`
		Volume int     `json:"volume,omitempty"`
		Value  float64 `json:"value,omitempty"`
	}{
		Open:   x.open,
		High:   x.high,
		Close:  x.close,
		Low:    x.low,
		Volume: x.volume,
		Value:  x.value,
	})
}

func (x *StockQuota) GetOpen() float64 {
	return x.open
}

func (x *StockQuota) SetOpen(open float64) {
	x.open = open
}

func (x *StockQuota) GetHigh() float64 {
	return x.high
}

func (x *StockQuota) SetHigh(high float64) {
	x.high = high
}

func (x *StockQuota) GetClose() float64 {
	return x.close
}

func (x *StockQuota) SetClose(closePrice float64) {
	x.close = closePrice
}

func (x *StockQuota) GetLow() float64 {
	return x.low
}

func (x *StockQuota) SetLow(low float64) {
	x.low = low
}

func (x *StockQuota) GetVolume() int {
	return x.volume
}

func (x *StockQuota) SetVolume(volume int) {
	x.volume = volume
}

func (x *StockQuota) GetValue() float64 {
	return x.value
}

func (x *StockQuota) SetValue(value float64) {
	x.value = value
}

func (x *StockQuota) GetUpdatedAt() time.Time {
	return x.updatedAt
}

func (x *StockQuota) SetUpdatedAt(updatedAt time.Time) {
	x.updatedAt = updatedAt
}
