package handler

import (
	"mezink-assignment/internal/model"
	"mezink-assignment/internal/service"
	"mezink-assignment/shared/base"
	"mezink-assignment/transport/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IRecordHandler interface {
	CreateRecord(c *gin.Context)
	GetRecord(c *gin.Context)
	SearchRecords(c *gin.Context)
	UpsertRecord(c *gin.Context)
	DeleteRecord(c *gin.Context)
}

type RecordHandler struct {
	Service service.IRecordService
}

func (h *Handler) CreateRecord(c *gin.Context) {
	var request model.RecordRequest
	if err := c.BindJSON(&request); err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	if err := request.Validate(); err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.Service.CreateRecord(request.Record()); err != nil {
		response.JSON(c, base.Failure(http.StatusInternalServerError, err.Error()))
		return
	}

	response.JSON(c, base.Success(http.StatusCreated))
}

func (h *Handler) GetRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := h.Service.GetRecordByID(id)
	if err != nil {
		response.JSON(c, base.Failure(http.StatusInternalServerError, err.Error()))
		return
	}

	response.JSON(c, base.Success(http.StatusOK, base.BaseData{
		Field: res.FieldName(),
		Data:  res,
	}))
}

func (h *Handler) SearchRecords(c *gin.Context) {
	var request model.RecordFilterRequest
	if err := c.BindJSON(&request); err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	if err := request.Validate(); err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := h.Service.GetRecordsByFilter(request)
	if err != nil {
		response.JSON(c, base.Failure(http.StatusInternalServerError, err.Error()))
		return
	}

	response.JSON(c, base.Success(http.StatusOK, base.BaseData{
		Field: res.FieldName(),
		Data:  res,
	}))
}

func (h *Handler) UpsertRecord(c *gin.Context) {
	var request model.RecordRequest
	if err := c.BindJSON(&request); err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	if err := request.Validate(); err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	record := request.Record()
	record.ID = int64(id)

	if err := h.Service.UpsertRecord(record); err != nil {
		response.JSON(c, base.Failure(http.StatusInternalServerError, err.Error()))
		return
	}

	response.JSON(c, base.Success(http.StatusOK))
}

func (h *Handler) DeleteRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JSON(c, base.Failure(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.Service.DeleteRecord(id); err != nil {
		response.JSON(c, base.Failure(http.StatusInternalServerError, err.Error()))
		return
	}

	response.JSON(c, base.Success(http.StatusOK))
}
