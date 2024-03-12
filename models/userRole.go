package models

type UserRole struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
}

type UserRoleListItem struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
}

type UserRoleCreateModel struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
}

type UserRoleUpdateModel struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
}
