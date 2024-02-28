package agg

import (
	"encoding/json"
	"strconv"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

// Stock is an aggregate root that represents the stock.
type Stock struct {
	stock         *model.Stock
	recentQuota   model.StockQuota
	earningsCalls []model.EarningsCall
}

// NewStock is the constructor of Stock.
func NewStock(stock *model.Stock) Stock {
	return Stock{
		stock: stock,
	}
}

// NewStockWithQuota is the constructor of Stock with recent quota.
func NewStockWithQuota(stock *model.Stock, recentQuota model.StockQuota) Stock {
	return Stock{
		stock:       stock,
		recentQuota: recentQuota,
	}
}

func (x *Stock) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*model.Stock
		RecentQuota *model.StockQuota `json:"recent_quota,omitempty"`
	}{
		Stock:       x.stock,
		RecentQuota: &x.recentQuota,
	})
}

func (x *Stock) UnmarshalJSON(bytes []byte) error {
	aux := &struct {
		*model.Stock `json:",inline"`
		RecentQuota  *model.StockQuota `json:"recent_quota,omitempty"`
	}{}

	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	x.stock = aux.Stock
	x.recentQuota = *aux.RecentQuota

	return nil
}

func (x *Stock) GetSymbol() string {
	return x.stock.Symbol
}

func (x *Stock) GetName() string {
	return x.stock.Name
}

func (x *Stock) GetIndustryCategory() string {
	return x.stock.IndustryCategory
}

func (x *Stock) GetExchangeName() string {
	return x.stock.ExchangeName
}

func (x *Stock) GetRecentQuota() *model.StockQuota {
	return &x.recentQuota
}

// FlexMessage is used to create a flex message.
func (x *Stock) FlexMessage() *linebot.FlexMessage {
	return linebot.NewFlexMessage("StockInfo", &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   x.stock.Symbol + " " + x.stock.Name,
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeXxl,
					Margin: linebot.FlexComponentMarginTypeMd,
				},
				&linebot.SeparatorComponent{
					Type:   linebot.FlexComponentTypeSeparator,
					Margin: linebot.FlexComponentMarginTypeXxl,
				},
				&linebot.BoxComponent{
					Type:    linebot.FlexComponentTypeBox,
					Layout:  linebot.FlexBoxLayoutTypeVertical,
					Margin:  linebot.FlexComponentMarginTypeXxl,
					Spacing: linebot.FlexComponentSpacingTypeSm,
					Contents: []linebot.FlexComponent{
						&linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent{
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  "Price",
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#555555",
									Flex:  linebot.IntPtr(0),
								},
								&linebot.TextComponent{
									Type:  linebot.FlexComponentTypeText,
									Text:  "$" + strconv.FormatFloat(x.recentQuota.GetClose(), 'f', 2, 64),
									Size:  linebot.FlexTextSizeTypeSm,
									Color: "#111111",
									Align: linebot.FlexComponentAlignTypeEnd,
								},
							},
						},
					},
				},
			},
		},
	})
}
