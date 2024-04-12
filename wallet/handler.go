package wallet

import (
	"net/http"

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
