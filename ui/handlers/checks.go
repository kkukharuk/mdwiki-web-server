package handlers

func CheckUser(username, password string) bool {
	return map[string]string{"admin": "Qwer1234"}[username] == password
}
