package mongoModels

import "learning-platform/board/internal/models/enum"

type BoardItem struct {
	ObjectID    string           `json:"object_id" bson:"object_id"`
	AuthorID    int64            `json:"author_id" bson:"author_id"`
	Type        enum.ItemType    `json:"type" bson:"type"`
	Left        int64            `json:"left" bson:"left"`
	Top         int64            `json:"top" bson:"top"`
	Width       int64            `json:"width" bson:"width"`
	Height      int64            `json:"height" bson:"height"`
	ZIndex      int64            `json:"z_index" bson:"z_index"`
	Angle       int64            `json:"angle" bson:"angle"`
	FigureType  *enum.FigureType `json:"figure_type,omitempty" bson:"figure_type,omitempty"`
	Points      []Point          `json:"points,omitempty" bson:"points,omitempty"`
	Color       *string          `json:"color,omitempty" bson:"color,omitempty"`
	Fill        *string          `json:"fill,omitempty" bson:"fill,omitempty"`
	StrokeWidth *int64           `json:"stroke_width,omitempty" bson:"stroke_width,omitempty"`
	Text        *string          `json:"text,omitempty" bson:"text,omitempty"`
	FontSize    *int64           `json:"font_size,omitempty" bson:"font_size,omitempty"`
	ImageURL    *string          `json:"image_url,omitempty" bson:"image_url,omitempty"`
}

type Point struct {
	Left int64 `json:"left" bson:"left"`
	Top  int64 `json:"top" bson:"top"`
}
