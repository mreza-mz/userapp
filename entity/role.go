package entity

type Role uint8

const (
	TenantRole Role = iota + 1
	ManagerRole
)

const (
	TenantRoleStr  = "tenant"
	ManagerRoleStr = "manager"
)

func (r Role) String() string {
	switch r {
	case TenantRole:
		return TenantRoleStr
	case ManagerRole:
		return ManagerRoleStr
	}

	return ""
}

func MapToRoleEntity(roleStr string) Role {
	switch roleStr {
	case ManagerRoleStr:
		return ManagerRole
	case TenantRoleStr:
		return TenantRole
	}

	return Role(0)
}
