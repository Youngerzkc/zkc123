package model
import(
	"time"
)
//User 用户
type User struct{
	ID   int  `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"column:name"`
	Pass string `json:"-" gorm:"column:password"`
	Create time.Time `json:"create_at" gorm:"column:create_at"`
}