package handler

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	v1 "github.com/truxcoder/trux-layout-advanced/api/v1"
	"github.com/truxcoder/trux-layout-advanced/internal/model"
	"github.com/truxcoder/trux-layout-advanced/internal/service"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type PersonHandler struct {
	*Handler
	personService service.PersonService
}

func NewPersonHandler(
	handler *Handler,
	personService service.PersonService,
) *PersonHandler {
	return &PersonHandler{
		Handler:       handler,
		personService: personService,
	}
}

func (h *PersonHandler) GetPerson(ctx *gin.Context) {
	var (
		err    error
		id     int64
		person *model.Person
	)
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if person, err = h.personService.GetPerson(ctx, id); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleSuccess(ctx, person)
}

func (h *PersonHandler) GetPeople(ctx *gin.Context) {
	var err error
	var req v1.PersonRequest
	var people []model.Person
	if err = jsoniter.Unmarshal([]byte(ctx.DefaultQuery("json", "{}")), &req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if people, err = h.personService.GetPeople(ctx, &req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleSuccess(ctx, people)
}

func (h *PersonHandler) CreatePerson(ctx *gin.Context) {
	var err error
	var req v1.PersonRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if err = h.personService.CreatePerson(ctx, &req); err != nil {
		h.logger.WithContext(ctx).Error("personService.CreatePerson error", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleCustomSuccess(ctx, v1.MsgCreateSuccess, nil)
}

func (h *PersonHandler) UpdatePerson(ctx *gin.Context) {
	var err error
	var req v1.PersonRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if err = h.personService.UpdatePerson(ctx, &req); err != nil {
		h.logger.WithContext(ctx).Error("personService.UpdatePerson error", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleCustomSuccess(ctx, v1.MsgUpdateSuccess, nil)
}

func (h *PersonHandler) DeletePerson(ctx *gin.Context) {
	var (
		err error
		id  int64
		ids []int64
	)
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	ids = append(ids, id)

	if err = h.personService.DeletePerson(ctx, ids); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleCustomSuccess(ctx, v1.MsgDeleteSuccess, nil)
}
