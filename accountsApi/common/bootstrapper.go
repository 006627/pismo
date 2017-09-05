package common

// responsible for initializing the entire system
func init() {
	initConfig()
	createDbSession()
}
