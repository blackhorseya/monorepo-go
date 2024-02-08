package model

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

// FlexMessage is used to create a flex message.
func (s *Stock) FlexMessage() *linebot.FlexMessage {
	return linebot.NewFlexMessage("StockInfo", &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   s.Symbol + " " + s.Name,
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
									Text:  "$" + strconv.FormatFloat(s.Price, 'f', 2, 64),
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

func (x *StockInfo) MarshalJSON() ([]byte, error) {
	type Alias StockInfo

	return json.Marshal(&struct {
		*Alias
		Date      string `json:"date,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	}{
		Alias:     (*Alias)(x),
		Date:      x.Date.AsTime().UTC().Format(time.RFC3339),
		UpdatedAt: x.UpdatedAt.AsTime().UTC().Format(time.RFC3339),
	})
}

func (x *DailyQuote) MarshalJSON() ([]byte, error) {
	type Alias DailyQuote

	return json.Marshal(&struct {
		*Alias
		Date string `json:"date,omitempty"`
	}{
		Alias: (*Alias)(x),
		Date:  x.Date.AsTime().UTC().Format(time.RFC3339),
	})
}
