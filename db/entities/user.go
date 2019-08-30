package entities

type UserEntity struct {
	Email string `gorm:"unique_index"`
	Name  string
}

func (UserEntity) TableName() string {
	return "users"
}
