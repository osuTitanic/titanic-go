package passwords

var passwordCache = map[string]bool{}

func GetPasswordCache() map[string]bool {
	return passwordCache
}

func ClearPasswordCache() {
	passwordCache = map[string]bool{}
}
