package domain

var ROLES []Role = []Role{
	RoleAdmin,
	RoleManager,
	RoleUser,
}

type Role struct {
	RoleName         string
	PermissionGroups []PermissionGroup
}

var (
	RoleAdmin Role = Role{
		RoleName: "ADMIN",
		PermissionGroups: []PermissionGroup{
			PermissionGroupNone,
			PermissionGroupUserManager,
		},
	}
	RoleManager Role = Role{
		RoleName: "MANAGER",
		PermissionGroups: []PermissionGroup{
			PermissionGroupNone,
			PermissionGroupUserViewer,
		},
	}
	RoleUser Role = Role{
		RoleName: "USER",
		PermissionGroups: []PermissionGroup{
			PermissionGroupNone,
			PermissionGroupUserViewer,
		},
	}
)

func GetRoles() []Role {
	return ROLES
}
