package enum

type ActionType string

const (
	AddItem     ActionType = "add"
	UpdateItem  ActionType = "update"
	DeleteItems ActionType = "delete"
	GetSnapshot ActionType = "snapshot"
)
