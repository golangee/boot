// build provides information about the build time.
package build

import "encoding/json"

// These variables will be set at compile time, as linker flags
var (
	Time   string
	Commit string
)

// Environment describe the Time and Commit of a build
type Environment struct {
	Time   string
	Commit string
}

// String returns a json marshalled text value
func (e Environment) String() string {
	b, _ := json.Marshal(Env())
	return string(b)
}

// Env returns the Time and Commit of the current build.
func Env() Environment {
	return Environment{
		Time:   Time,
		Commit: Commit,
	}
}
