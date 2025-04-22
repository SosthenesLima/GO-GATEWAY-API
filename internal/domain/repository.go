package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(APIKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	FindAll() ([]*Account, error)
	UpdateBalance(Account *Account) error
}
