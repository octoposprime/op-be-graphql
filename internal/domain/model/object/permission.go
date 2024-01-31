package domain

var PERMISSIONS []Permission = []Permission{
	PermissionNone,

	PermissionUserRead,
	PermissionUserList,
	PermissionUserCreate,
	PermissionUserUpdateRole,
	PermissionUserUpdateBase,
	PermissionUserUpdateStatus,
	PermissionUserChangePassword,
	PermissionUserDelete,
}

type Permission struct {
	Policy string
}

var (
	PermissionNone Permission = Permission{Policy: "none"}

	PermissionUserRead           Permission = Permission{Policy: "user.read"}
	PermissionUserList           Permission = Permission{Policy: "user.list"}
	PermissionUserCreate         Permission = Permission{Policy: "user.create"}
	PermissionUserUpdateRole     Permission = Permission{Policy: "user.update.role"}
	PermissionUserUpdateBase     Permission = Permission{Policy: "user.update.base"}
	PermissionUserUpdateStatus   Permission = Permission{Policy: "user.update.status"}
	PermissionUserChangePassword Permission = Permission{Policy: "user.change.password"}
	PermissionUserDelete         Permission = Permission{Policy: "user.delete"}
)

func GetPermissions() []Permission {
	return PERMISSIONS
}
