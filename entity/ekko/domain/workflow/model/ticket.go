package model

import (
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

// Ticket is a value object that represents a ticket.
type Ticket struct {
	ID        string `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

func (t *Ticket) BubbleContainer() *linebot.BubbleContainer {
	return &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   t.Title,
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeXxl,
				},
				&linebot.SeparatorComponent{
					Type: linebot.FlexComponentTypeSeparator,
				},
			},
		},
	}
}

// FlexMessage returns a flex message of the ticket.
func (t *Ticket) FlexMessage() *linebot.FlexMessage {
	return linebot.NewFlexMessage("Ticket", t.BubbleContainer())
}

// NewTicket creates a new ticket.
func NewTicket(title string) (*Ticket, error) {
	return &Ticket{
		Title: title,
	}, nil
}

// Tickets is a collection of tickets.
type Tickets []*Ticket

func (x Tickets) FlexMessage() *linebot.FlexMessage {
	var contents []*linebot.BubbleContainer
	for _, ticket := range x {
		contents = append(contents, ticket.BubbleContainer())
	}

	return linebot.NewFlexMessage("Ticket List", &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: contents,
	})
}
