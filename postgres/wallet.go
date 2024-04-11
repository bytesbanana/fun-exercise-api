package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Wallets(walletType string) ([]wallet.Wallet, error) {
	var rows *sql.Rows
	var err error
	if walletType == "" {
		rows, err = p.Db.Query("SELECT * FROM user_wallet")
	} else {
		rows, err = p.Db.Query("SELECT * FROM user_wallet WHERE wallet_type = $1", walletType)
	}

	if err != nil {
		return nil, errors.New("failed to get wallets")
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) WalletsByUserID(userID int) ([]wallet.Wallet, error) {
	var rows *sql.Rows
	var err error
	rows, err = p.Db.Query("SELECT * FROM user_wallet WHERE user_id = $1", userID)
	if err != nil {
		return nil, errors.New("failed to get wallets")
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) CreateWallet(w wallet.Wallet) (*wallet.Wallet, error) {

	stmt := "INSERT INTO user_wallet (user_id, user_name, wallet_name, wallet_type, balance, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"

	row := p.Db.QueryRow(stmt,
		w.UserID, w.UserName, w.WalletName, w.WalletType, w.Balance, time.Now())

	var newWallet wallet.Wallet
	row.Scan(&newWallet.ID, &newWallet.UserID, &newWallet.UserName, &newWallet.WalletName, &newWallet.WalletType, &newWallet.Balance, &newWallet.CreatedAt)

	return &newWallet, nil
}
