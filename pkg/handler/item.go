package handler

import (
	"github.com/Grishun/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) newItem(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := h.services.TodoItem.NewItem(userId, listId, input)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"itemId": itemId,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

	userId, err := GetUserId(c)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	items := make([]todo.TodoItem, 0)

	items, err = h.services.TodoItem.GetAllItems(userId, listId)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)

}
func (h *Handler) getItemById(c *gin.Context) {

	userId, err := GetUserId(c)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	item, err := h.services.TodoItem.GetItemById(userId, itemId)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.UpdateItemInput
	err = c.BindJSON(&input)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.TodoItem.Update(userId, itemId, input)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "successful",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "successful",
	})
}
