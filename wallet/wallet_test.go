package wallet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

type StubHandler struct {
	wallets []Wallet
	err     error
}

func (s *StubHandler) Wallets() ([]Wallet, error) {
	return s.wallets, s.err
}

func TestWallet(t *testing.T) {

	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/wallets")

		handlers := New(&StubHandler{
			err: echo.ErrInternalServerError,
		})
		handlers.WalletHandler(c)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
		resp := &Err{}
		json.Unmarshal(rec.Body.Bytes(), resp)
		want := &Err{Message: "code=500, message=Internal Server Error"}
		if resp.Message != want.Message {
			t.Errorf("expected status code %s but got %s", want.Message, resp.Message)
		}
	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/wallets")

		handlers := New(&StubHandler{
			wallets: []Wallet{
				{
					ID:         1,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Create Card",
					Balance:    100,
					CreatedAt:  time.Now(),
				},
				{
					ID:         2,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Create Card",
					Balance:    100,
					CreatedAt:  time.Now(),
				},
			},
		})

		handlers.WalletHandler(c)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		resp := &[]Wallet{}
		json.Unmarshal(rec.Body.Bytes(), resp)

		if len(*resp) != 2 {
			t.Errorf("expected status code %d but got %d", 1, len(*resp))
		}

	})
}
