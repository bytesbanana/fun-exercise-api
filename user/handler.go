package user

import (
	"net/http"
	"strconv"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	WalletsByUserID(userId int) ([]wallet.Wallet, error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// WalletByUserIdHandler
//
//	@Summary		Get all wallets by user id
//	@Description	Get all wallets by user id
//	@Router			/api/v1/users/{user_id}/wallets [get]
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	wallet.Wallet
//	@Failure		400	{object}	Err
//	@Failure		500	{object}	Err
//	@Param   user_id  path	string	true "User id"
func (h *Handler) WalletByUserIdHandler(c echo.Context) error {
	pUserId := c.Param("id")
	userId, err := strconv.Atoi(pUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid user id"})
	}

	wallets, err := h.store.WalletsByUserID(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}
