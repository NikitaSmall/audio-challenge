package config

import (
	"github.com/subosito/gotenv"
)

// InitEnv can load env variables from different sources.
// Count them in params when call this function.
func InitEnv(filenames ...string) {
	for _, file := range filenames {
		gotenv.Load(file)
	}
}
