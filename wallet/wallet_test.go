package wallet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

type StubWalletHandler struct {
	wallets []Wallet
	err     error
}

func (s *StubWalletHandler) Wallets(walletType string) ([]Wallet, error) {

	if _, ok := WalletType[walletType]; ok {
		filteredWallets := []Wallet{}
		for _, w := range s.wallets {
			if w.WalletType == walletType {
				filteredWallets = append(filteredWallets, w)
			}
		}
		return filteredWallets, nil
	}

	return s.wallets, s.err
}
func (w *StubWalletHandler) CreateWallet(wallet Wallet) (*Wallet, error) {
	lastWalletId := 0
	if len(w.wallets) > 0 {
		lastWalletId = w.wallets[len(w.wallets)-1].ID
	}

	wallet.ID = lastWalletId + 1
	wallet.CreatedAt = time.Now()
	w.wallets = append(w.wallets, wallet)
	return &w.wallets[len(w.wallets)-1], nil
}

func setup(t *testing.T, buildRequestFunc func() *http.Request) (echo.Context, *httptest.ResponseRecorder) {
	t.Parallel()
	e := echo.New()
	req := buildRequestFunc()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/wallets")

	return c, rec
}

func TestWallet(t *testing.T) {
	// t.Parallel()

	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		c, rec := setup(t, func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/", nil)
		})

		handlers := New(&StubWalletHandler{
			err: echo.ErrInternalServerError,
		})
		handlers.GetWallet(c)
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

	t.Run("given user able to getting wallet should return all list of wallets", func(t *testing.T) {
		c, rec := setup(t, func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/", nil)
		})

		handlers := New(&StubWalletHandler{
			wallets: []Wallet{
				{
					ID:         1,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Credit Card",
					Balance:    100,
					CreatedAt:  time.Now(),
				},
				{
					ID:         2,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Credit Card",
					Balance:    100,
					CreatedAt:  time.Now(),
				},
			},
		})

		handlers.GetWallet(c)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		resp := []Wallet{}
		json.Unmarshal(rec.Body.Bytes(), &resp)
		if len(resp) != 2 {
			t.Errorf("expected status code %d but got %d", 1, len(resp))
		}

	})

	t.Run("given wallet type should return list of wallets type", func(t *testing.T) {
		q := make(url.Values)
		q.Set("wallet_type", "Savings")
		c, rec := setup(t, func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		})

		handlers := New(&StubWalletHandler{
			wallets: []Wallet{
				{
					ID:         1,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Savings",
					Balance:    100,
					CreatedAt:  time.Now(),
				},
				{
					ID:         2,
					UserName:   "John Doe",
					WalletName: "John's Wallet",
					WalletType: "Credit Card",
					Balance:    100,
					CreatedAt:  time.Now(),
				},
			},
		})

		handlers.GetWallet(c)
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		resp := []Wallet{}
		json.Unmarshal(rec.Body.Bytes(), &resp)
		if len(resp) != 1 {
			t.Errorf("expected wallet length %d but got %d", 1, len(resp))
		}
	})

	t.Run("given invalid wallet type should return 400 and error message", func(t *testing.T) {
		q := make(url.Values)
		q.Set("wallet_type", "InvalidWalletType")
		c, rec := setup(t, func() *http.Request {
			return httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		})

		handlers := New(&StubWalletHandler{})
		handlers.GetWallet(c)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rec.Code)
		}
		resp := &Err{}
		json.Unmarshal(rec.Body.Bytes(), resp)
		want := &Err{Message: "Invalid wallet type"}
		if resp.Message != want.Message {
			t.Errorf("expected status code %s but got %s", want.Message, resp.Message)
		}

	})

	t.Run("given wallet info should create new wallet", func(t *testing.T) {

		c, rec := setup(t, func() *http.Request {
			walletJSON := `{
				"user_id": 2,
				"user_name": "Chivas",
				"wallet_name": "My Saving",
				"wallet_type": "Savings",
				"balance": 100
			}`
			return httptest.NewRequest(http.MethodPost, "/", strings.NewReader(walletJSON))

		})

		handlers := New(&StubWalletHandler{
			wallets: []Wallet{},
		})
		handlers.CreateWallet(c)
		if rec.Code != http.StatusCreated {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}
		resp := &Wallet{}
		json.Unmarshal(rec.Body.Bytes(), resp)
		if resp.ID == 0 {
			t.Errorf("expected status code %d but got %d", 1, resp.ID)
		}
	})
}
