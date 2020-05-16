package userconf

import (
	"github.com/athomecomar/envconf"
)

func GetDATABASE_NAME() (db string) {
	switch envconf.GetENV() {
	case envconf.Development:
		db = "db_dev"
	case envconf.Staging:
		db = "db_stg"
	case envconf.Production:
		db = "db_prd"
	}

	return
}
