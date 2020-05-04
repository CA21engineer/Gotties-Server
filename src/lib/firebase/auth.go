package firebase

//https://firebase.google.com/docs/auth/admin/
type Auth struct {

}

type AuthImpl interface {
	IsUserLoggedIn(string)error
}

func NewAuth() AuthImpl  {
	return &Auth{}
}

//受け取ったuuidからそのユーザーがすでにログイン済みかを確認する
func (a *Auth) IsUserLogedIn(uuid string) error{

}