package musik

import "os"

func Dangdut() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = ":80"
	}
	return
}

func Koplo() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	return
}

func Tarling() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	return
}
