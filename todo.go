package todo

import "errors"

type (
	TodoList struct {
		Id          int    `json:"id" db:"id"`
		Title       string `json:"title" db:"title" binding:"required"`
		Description string `json:"description" db:"description" binding:"required" `
	}
	UserList struct {
		Id          int    `json:"id" db:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	ListItem struct {
		Id     int `json:"id"`
		ListId int `json:"list_id"`
		ItemId int `json:"item_id"`
	}
	TodoItem struct {
		Id          int    `json:"id"`
		Title       string `json:"title" db:"title" binding:"required"`
		Description string `json:"description" db:"description" binding:"required"`
		Done        bool   `json:"done" db:"done"`
	}

	UpdateListInput struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	UpdateItemInput struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Done        *bool   `json:"done"`
	}
)

func ValidateUpdate(input UpdateListInput) error {
	if input.Title == nil && input.Description == nil {
		return errors.New("nothing to update")
	}
	return nil
}
