package agg

import (
	"fmt"
	"log"

	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/model"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

// Asset is an aggregate root.
type Asset struct {
	model.Car
	Distance float64 `json:"distance"`
}

// BubbleContainer is to create a bubble container for car.
func (x *Asset) BubbleContainer() *linebot.BubbleContainer {
	return &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Size: linebot.FlexBubbleSizeTypeKilo,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   x.Id,
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeXxl,
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: fmt.Sprintf("↔️ %.2f km", x.Distance),
				},
				&linebot.SeparatorComponent{
					Type: linebot.FlexComponentTypeSeparator,
				},
				&linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: linebot.NewURIAction("Google Map", x.Location.GoogleMap()),
				},
			},
		},
	}
}

type Assets []*Asset

// FlexMessage is to create a flex message for car list.
func (x Assets) FlexMessage() *linebot.FlexMessage {
	var contents []*linebot.BubbleContainer
	for _, asset := range x {
		contents = append(contents, asset.BubbleContainer())
	}

	log.Println(contents)

	return linebot.NewFlexMessage("Car List", &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: contents,
	})
}
