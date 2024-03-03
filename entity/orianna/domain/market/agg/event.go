package agg

import (
	"fmt"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/timex"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

// Event is an aggregate root that represents the event.
type Event struct {
	*model.Event
}

// BubbleContainer is to create a bubble container for event.
func (x *Event) BubbleContainer() *linebot.BubbleContainer {
	return &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Size: linebot.FlexBubbleSizeTypeKilo,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   x.Symbol,
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeXxl,
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: fmt.Sprintf("📅 %s", x.OccurredAt.In(timex.LocTaipei).Format("2006-01-02")),
				},
			},
		},
	}
}

// Events is a collection of Event.
type Events []*Event

// FlexMessage is to create a flex message for event list.
func (x Events) FlexMessage() *linebot.FlexMessage {
	var contents []*linebot.BubbleContainer
	for _, asset := range x {
		contents = append(contents, asset.BubbleContainer())
	}

	return linebot.NewFlexMessage("Event List", &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: contents,
	})
}
