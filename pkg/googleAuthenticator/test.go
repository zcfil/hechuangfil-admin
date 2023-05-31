package googleAuthenticator

func createSecret(ga *GAuth) string {
	secret, err := ga.CreateSecret(16)
	if err != nil {
		return ""
	}
	return secret
}

func main() {

}
