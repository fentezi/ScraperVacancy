package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fentezi/scraper/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Jobs(c *gin.Context) {
	infoDjinni := pkg.ParseDjinni()
	infoDou := pkg.ParserDou()
	jsonDjinni, err := json.MarshalIndent(infoDjinni, "", " ")
	jsonDou, err := json.MarshalIndent(infoDou, "", " ")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	result := append(jsonDjinni, jsonDou...)

	c.Data(http.StatusOK, "application/json", result)
}
