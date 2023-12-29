package validator

import (
	"Farashop/internal/contract"
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
)

func DoesUsernameActive(ctx context.Context, contract contract.ValidatorStore) validation.RuleFunc {
	return func(value interface{}) error {
		username := value.(string)
		ok, err := contract.DoesUsernameActiveStore(ctx, username)
		if err != nil || !ok {
			return err
		}
		return nil
	}
}

func DoesIDExist(ctx context.Context, contract contract.ValidatorStore) validation.RuleFunc {
	return func(value interface{}) error {
		id := value.(uint)
		ok, err := contract.DoesIDExistStore(ctx, id)
		if err != nil || !ok {
			return err
		}
		return nil
	}
}
