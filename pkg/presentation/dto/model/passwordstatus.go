package presentation

var PasswordStatus_name = map[int32]PasswordStatus{
	0: PasswordStatusNone,
	1: PasswordStatusActive,
	2: PasswordStatusInactive,
	3: PasswordStatusAutoGenerated,
	4: PasswordStatusChangeRequired,
	5: PasswordStatusExpired,
}
var PasswordStatus_value = map[PasswordStatus]int32{
	PasswordStatusNone:           0,
	PasswordStatusActive:         1,
	PasswordStatusInactive:       2,
	PasswordStatusAutoGenerated:  3,
	PasswordStatusChangeRequired: 4,
	PasswordStatusExpired:        5,
}
