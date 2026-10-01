package enum

type ItemType string

const (
	TypeLine   ItemType = "LINE"
	TypeFigure ItemType = "FIGURE"
	TypeText   ItemType = "TEXT"
	TypeImage  ItemType = "IMAGE"
)
