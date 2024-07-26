package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	/*CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`*/
}
