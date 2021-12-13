package authentication

type UseCase interface {
	GenerateJWT(login string, password string) (string, error)
}
