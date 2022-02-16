package authentication

import "context"

type UseCase interface {
	GenerateJWT(ctx context.Context, login string, password string) (string, error)
}
