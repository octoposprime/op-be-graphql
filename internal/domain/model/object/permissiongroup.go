package domain

var PERMISSIONGROUPS []PermissionGroup = []PermissionGroup{
	PermissionGroupNone,

	PermissionGroupUserManager,
	PermissionGroupUserViewer,
}

type PermissionGroup struct {
	GroupName   string
	Permissions []Permission
}

var (
	PermissionGroupNone PermissionGroup = PermissionGroup{
		GroupName: "none",
		Permissions: []Permission{
			PermissionNone,
		},
	}

	PermissionGroupUserManager PermissionGroup = PermissionGroup{
		GroupName: "user.manager",
		Permissions: []Permission{
			PermissionUserRead,
			PermissionUserList,
			PermissionUserCreate,
			PermissionUserUpdateRole,
			PermissionUserChangePassword,
			PermissionUserUpdateBase,
			PermissionUserUpdateStatus,
			PermissionUserDelete,
		},
	}
	PermissionGroupUserViewer PermissionGroup = PermissionGroup{
		GroupName: "user.viewer",
		Permissions: []Permission{
			PermissionUserRead,
			PermissionUserList,
		},
	}
)

func GetPermissionGroups() []PermissionGroup {
	return PERMISSIONGROUPS
}
