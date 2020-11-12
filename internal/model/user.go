package model

type User struct {
	*Model
	Password string `json:"password" form:"password"`
	NickName string `json:"nickName" form:"nickName"`
	Avatar   string `json:"avatar" form:"avatar"`
}

func (u User) TabName() string {
	return `user`
}

func (u User) Create() error {
	return DB.Create(&u).Error
}

func (u *User) FindOne() error {
	return DB.First(u,u.Id).Error
}

func (u User) Update() error {
	return DB.Model(&u).Updates(u).Where("id=?", u.Id).Error
}
