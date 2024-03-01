package utils

func RoleMap(role uint) string {
	if role == 0 {
		return "User"
	} else if role == 1 {
		return "Student"
	} else if role == 2 {
		return "Admin"
	}
	return "Invalid Role"
}
