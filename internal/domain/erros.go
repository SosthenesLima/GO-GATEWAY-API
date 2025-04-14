package domain

import "errors"

var (
	// ErrAccountNotFound é retornado quando uma conta não é encontrada.
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplitedAPIKey é retornado quando há rentativa de criar conta com API key duplicada.
	ErrDuplitedAPIKey = errors.New("api key alreadry exists")
	// ErrInvoiceNotFoun é um retornado quando uma fatura não é encontrada.
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAcces é retornado quando há tentativa de acesso não autorizado a um recurso.
	ErrUnauthorizedAccess = errors.New("unuathorized access")
)
