package model

import (
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
	close float64,
	low float64,
	volume int,
	updatedAt time.Time,
) StockQuota {
	return StockQuota{
		open:      open,
		high:      high,
		close:     close,
		low:       low,
		volume:    volume,
		value:     close * float64(volume),
		updatedAt: updatedAt,
	}
}

func (s *StockQuota) GetOpen() float64 {
	return s.open
}

func (s *StockQuota) SetOpen(open float64) {
	s.open = open
}

func (s *StockQuota) GetHigh() float64 {
	return s.high
}

func (s *StockQuota) SetHigh(high float64) {
	s.high = high
}

func (s *StockQuota) GetClose() float64 {
	return s.close
}

func (s *StockQuota) SetClose(close float64) {
	s.close = close
}

func (s *StockQuota) GetLow() float64 {
	return s.low
}

func (s *StockQuota) SetLow(low float64) {
	s.low = low
}

func (s *StockQuota) GetVolume() int {
	return s.volume
}

func (s *StockQuota) SetVolume(volume int) {
	s.volume = volume
}

func (s *StockQuota) GetValue() float64 {
	return s.value
}

func (s *StockQuota) SetValue(value float64) {
	s.value = value
}

func (s *StockQuota) GetUpdatedAt() time.Time {
	return s.updatedAt
}

func (s *StockQuota) SetUpdatedAt(updatedAt time.Time) {
	s.updatedAt = updatedAt
}
