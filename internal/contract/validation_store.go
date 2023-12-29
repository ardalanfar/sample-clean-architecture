package contract

import "context"

type ValidatorStore interface {
	DoesUsernameActiveStore(ctx context.Context, username string) (bool, error)
	DoesIDExistStore(ctx context.Context, id uint) (bool, error)
}
