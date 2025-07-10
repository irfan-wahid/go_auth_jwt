package models

const UsersTableName = "users"

type Users struct {
	Id        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	Firstname string `gorm:"column:firstname" json:"firstname"`
	Lastname  string `gorm:"column:lastname" json:"lastname"`
}

func (*Users) TableName() string {
	return UsersTableName
}
