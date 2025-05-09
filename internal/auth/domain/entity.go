package domain

import "gorm.io/gorm"

type UserRole string

const (
	RoleAdmin          UserRole = "Admin"
	RoleWarehouseStaff UserRole = "WarehouseStaff"
	RoleCustomer       UserRole = "Customer"
	RoleMember         UserRole = "Member"
)

// Database Model User
type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Role == "" {
		u.Role = string(RoleMember)
	}
	return nil
}
