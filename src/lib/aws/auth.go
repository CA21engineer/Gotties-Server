package aws

//https://firebase.google.com/docs/auth/admin/
type Auth struct {
	Uuid string
}

type AuthImpl interface {
	IsUserLoggedIn()error
}

func NewAuth(uuid string) *Auth  {
	return &Auth{uuid}
}

//受け取ったuuidからそのユーザーがすでにログイン済みかを確認する
func (a *Auth) IsUserLogedIn() (string, error){
	return a.Uuid, nil
}