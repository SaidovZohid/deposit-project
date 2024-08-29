package v1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/SaidovZohid/deposit-project/api/models"
	"github.com/SaidovZohid/deposit-project/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (h *handlerV1) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "strconv atoi not worked",
		})
		return
	}
	data, err := h.strg.User().GetById(ctx, int64(idInt))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "user not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "user get by id not worked",
		})
		return
	}
	ctx.JSON(http.StatusOK, parseUserRepoToApi(data))
}

func parseUserRepoToApi(user *repo.UserModelResp) models.UserModelResp {
	resp := models.UserModelResp{
		Id:        user.Id,
		Email:     user.Email,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt.Format(time.RFC1123Z),
	}
	if user.FullName.Valid {
		resp.FullName = &user.FullName.String
	}
	if user.PhoneNumber.Valid {
		resp.PhoneNumber = &user.PhoneNumber.String
	}
	if user.UpdatedAt.Valid {
		tm := user.UpdatedAt.Time.Format(time.Layout)
		resp.UpdatedAt = &tm
	}
	return resp
}
