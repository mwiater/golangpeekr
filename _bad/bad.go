package helpers

// Bad struct...
type Bad struct {
	Blah                     int
	Etc                      int
	YouShouldntSeeThisStruct string
}

// YouShouldntSeeThisFunction()...
func YouShouldntSeeThisFunction() string {
	return "!!!"
}

// YouShouldntSeeThisFunctionEither()...
func YouShouldntSeeThisFunctionEither() string {
	return "???"
}
