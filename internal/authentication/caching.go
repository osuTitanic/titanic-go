package passwords

// TODO: Make sure this doesn't cause memory issues in the future
//       There should be some sort of max size, but this will do (for now *cough*)

var passwordCache = map[string]bool{}

func GetPasswordCache() map[string]bool {
	return passwordCache
}

func ClearPasswordCache() {
	passwordCache = map[string]bool{}
}
