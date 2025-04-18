package repository

import (
	"database/sql"
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

type AcoountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AcoountRepository {
	return &AcoountRepository{db: db}
}

func (r *AcoountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare(`
	   INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at)
	   VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreateAt,
		account.UpdateAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *AcoountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var createdAT, UpdateAT time.Time

	err := r.db.QueryRow(`
	    SELECT id, name, email, api_key, balance, created_at, updated_at
	    FROM accounts
		WHERE api_key = $1
	`, apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAT,
		&UpdateAT,
	)

	if err == sql.ErrConnDone {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	account.CreateAt = createdAT
	account.UpdateAt = UpdateAT
	return &account, nil

}

func (r *AcoountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAT, UpdateAT time.Time

	err := r.db.QueryRow(`
	    SELECT id, name, email, api_key, balance, created_at, updated_at
	    FROM accounts
		WHERE id = $1
	`, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAT,
		&UpdateAT,
	)

	if err == sql.ErrConnDone {
		return nil, domain.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}

	account.CreateAt = createdAT
	account.UpdateAt = UpdateAT
	return &account, nil
}

func (r *AcoountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow(`SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`,
		account.ID).Scan(&currentBalance)

	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
	   UPDATE accounts
	   SET balance = $1, update_at = $2
	   WHERE id = $3
	`, account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}
	return tx.Commit()
}
