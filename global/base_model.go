/**
 * @Description ้ๅ
 **/
package global

import (
	"gorm.io/gorm"
	"time"
)

// ้ๅgorm.Model
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
