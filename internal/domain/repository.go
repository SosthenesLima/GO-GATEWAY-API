package domain

type AcoountRepository interface {
	Save(account *Account) error
	FindByAPIKey(APIKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	Update(Account *Account) error
}
