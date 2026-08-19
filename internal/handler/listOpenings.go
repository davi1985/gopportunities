package handler

import (
	"net/http"

	"github.com/davi1985/gopportunities/internal/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	rawOpenings := []schemas.Opening{}

	if err := db.Find(&rawOpenings).Error; err != nil {
		sendError(
			ctx,
			http.StatusInternalServerError,
			"error listing openings",
		)
		return
	}

	openings := make(
		[]schemas.OpeningResponse,
		len(rawOpenings),
	)

	for i, raw := range rawOpenings {
		openings[i] = schemas.NewOpeningResponse(raw)
	}

	sendSuccess(
		ctx,
		http.StatusOK,
		"list-openings",
		&openings,
	)
}
