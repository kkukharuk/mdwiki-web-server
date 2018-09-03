package handlers

func CheckUser(username, password string) bool {
	return map[string]string{"admin": "admin"}[username] == password
}
