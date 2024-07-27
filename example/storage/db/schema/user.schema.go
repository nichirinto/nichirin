package schema

type User struct {
	Id       string `gorm:"column:id;primaryKey:true"`
	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}
