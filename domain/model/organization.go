package model

// Organization ...
type Organization struct {
	BaseWithHashID
	Name              string             `json:"name"`
	Users             []User             `json:"organization_users" gorm:"many2many:organization_users"`
	OrganizationUsers []OrganizationUser `json:"-"`
}

// TableName ...
func (Organization) TableName() string {
	return "organizations"
}
