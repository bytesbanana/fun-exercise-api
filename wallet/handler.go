package wallet

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var WalletType = map[string]string{
	"Savings":      "Savings",
	"CreditCard":   "Credit Card",
	"CryptoWallet": "Crypto Wallet",
}

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets(walletType string) ([]Wallet, error)
	CreateWallet(wallet Wallet) (*Wallet, error)
	UpdateWallet(wallet Wallet) (*Wallet, error)
	DeleteWallet(id int) error
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// GetWallet
//
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		500	{object}	Err
//	@Failure		400	{object}	Err
//	@Param   wallet_type  query	string	false	"Wallet type"	Enums(Savings, CreditCard, CryptoWallet)
func (h *Handler) GetWallet(c echo.Context) error {
	walletType := c.QueryParam("wallet_type")

	if walletType != "" {
		if _, ok := WalletType[walletType]; !ok {
			return c.JSON(http.StatusBadRequest, Err{Message: "Invalid wallet type"})
		}
	}

	wallets, err := h.store.Wallets(WalletType[walletType])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

// Create new wallet
//
// @Summary		Create new wallet
// @Description	Create new wallet
// @Tags			wallet
// @Accept			json
// @Produce		json
// @Router			/api/v1/wallets [post]
// @Success		200	{object}	Wallet
// @Failure		400	{object}	Err
// @Failure		500	{object}	Err
// @Param   wallet  body		Wallet	true	"Wallet"
func (h *Handler) CreateWallet(c echo.Context) error {

	var w Wallet
	if err := c.Bind(&w); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	wallet, err := h.store.CreateWallet(w)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, wallet)
}

// UpdateWallet
//
// @Summary		Update wallet
// @Description	Update wallet
// @Tags			wallet
// @Accept			json
// @Produce		json
// @Router			/api/v1/wallets/{id} [put]
// @Success		200	{object}	Wallet
// @Failure		400	{object}	Err
// @Failure		500	{object}	Err
// @Param   id  path		int	true	"Wallet id"
// @Param   wallet  body		Wallet	true	"Wallet"
func (h *Handler) UpdateWallet(c echo.Context) error {

	var w Wallet
	if err := c.Bind(&w); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	pWalletId := c.Param("id")
	walletId, err := strconv.Atoi(pWalletId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid wallet id"})
	}
	w.ID = walletId

	wallet, err := h.store.UpdateWallet(w)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, wallet)
}

// DeleteWallet
//
// @Summary		Delete wallet
// @Description	Delete wallet
// @Tags			wallet
// @Accept			json
// @Produce		json
// @Router			/api/v1/wallets/{id} [delete]
// @Success		204
// @Failure		400	{object}	Err
// @Failure		500	{object}	Err
// @Param   id  path		int	true	"Wallet id"
func (h *Handler) DeleteWallet(c echo.Context) error {
	pWalletId := c.Param("id")
	walletId, err := strconv.Atoi(pWalletId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "Invalid wallet id"})
	}

	if err := h.store.DeleteWallet(walletId); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
