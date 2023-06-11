package format_plugin

// recover错误，转string
func ErrorToString(r interface{}) string {
	switch t := r.(type) {
	case error:
		return t.Error()
	default:
		return r.(string)
	}
}
