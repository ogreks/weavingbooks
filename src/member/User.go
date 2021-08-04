package member

type UserApi interface {
    CreateUser(user User) User
    GetUser(account interface{}) User
    Login(account, password string) User
    Logout(id int64)
}

type User struct {
    Id int64 `json:"id"`
    Account string `json:"account"`
    Name string `json:"name"`
    Password string `json:"password"`
    Locker bool `json:"locker"`
    Locknums int64 `json:"locknums"`
    Createtime int64 `json:"createtime"`
}