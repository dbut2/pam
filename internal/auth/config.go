package auth

type Config struct {
	creds []byte
}

func ConfigWithCreds(creds []byte) Config {
	return Config{
		creds: creds,
	}
}
