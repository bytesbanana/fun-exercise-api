package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"
)

type StubUserHandler struct {
	wallets []wallet.Wallet
	err     error
}

func (w *StubUserHandler) WalletsByUserID(userId int) ([]wallet.Wallet, error) {
	filteredWallets := []wallet.Wallet{}
	for _, w := range w.wallets {
		if w.UserID == userId {
			filteredWallets = append(filteredWallets, w)
		}
	}

	return filteredWallets, w.err
}

func TestUser(t *testing.T) {

	t.Run("given user id should return list of wallets", func(t *testing.T) {
		t.Parallel()
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id/wallets")
		c.SetParamNames("id")
		c.SetParamValues("2")

		handlers := New(&StubUserHandler{
			wallets: []wallet.Wallet{
				{
					ID:         1,
					UserID:     1,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Credit Card",
					Balance:    100,
					CreatedAt:  time.Now(),
				},
				{
					ID:         2,
					UserID:     2,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Credit Card",
					Balance:    100,
					CreatedAt:  time.Now(),
				}, {
					ID:         3,
					UserID:     2,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Credit Card",
					Balance:    100,
					CreatedAt:  time.Now(),
				},
			},
		})
		handlers.WalletByUserId(c)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		resp := []wallet.Wallet{}
		json.Unmarshal(rec.Body.Bytes(), &resp)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		if len(resp) != 2 {
			t.Errorf("expected status code %d but got %d", 1, len(resp))
		}

	})

}
