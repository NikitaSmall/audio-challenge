/*Package config holds ways to get and setup environment and configs.
 * In current implementation this is only one function to setup
 * environment from `.env` file.
 */
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
