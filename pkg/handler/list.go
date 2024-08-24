package handler

import (
	"github.com/Grishun/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) newList(c *gin.Context) {

	userId, err := GetUserId(c)
	if err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.Todolist.CreateList(userId, input)
	if err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"listId": listId,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {

	userId, err := GetUserId(c)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.Todolist.GetAllLists(userId)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{Data: lists})

}
func (h *Handler) getListById(c *gin.Context) {
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

	list, err := h.services.Todolist.GetListById(userId, listId)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := GetUserId(c)

	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	} else if _, err := h.services.Todolist.GetListById(userId, listId); err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Todolist.Update(userId, listId, input); err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"status of update": "successful",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	} else if _, err := h.services.Todolist.GetListById(userId, listId); err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.Todolist.DeleteList(userId, listId); err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "successful deleted",
	})
}
