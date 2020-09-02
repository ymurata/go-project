package model

// OrganizationUser ...
type OrganizationUser struct {
	BaseWithHashID
	OrganizationID BaseWithHashID `json:"organization_id"`
	UserID         BaseWithHashID `json:"user_id"`
}

// TableName ...
func (OrganizationUser) TableName() string {
	return "organization_users"
}
