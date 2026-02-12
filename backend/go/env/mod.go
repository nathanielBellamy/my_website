package env

type Env int

func IsProd(mode string) bool {
	return mode == "prod"
}

func IsLocalhost(mode string) bool {
	return mode == "localhost" || mode == ""
}

func IsRemotedev(mode string) bool {
	return mode == "remotedev"
}
