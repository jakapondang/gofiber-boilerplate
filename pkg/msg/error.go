package msg

func PanicLogging(err interface{}) {
	if err != nil {
		panic(err)
	}
}
