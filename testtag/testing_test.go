package testtag

func init() {
	IsTesting = true
}

func Expose() bool { return IsTesting }
