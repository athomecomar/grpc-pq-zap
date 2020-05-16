package userconf

import "github.com/athomecomar/envconf"

func GetPORT() (port string) {
	switch envconf.GetENV() {
	case envconf.Development:
		port = ":50051"
	case envconf.Staging, envconf.Production:
		port = ":30051"
	}
	return
}
