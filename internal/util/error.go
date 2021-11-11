package util

func WrapError(e error) map[string]interface{} {
	return map[string]interface{}{"error": e.Error()}
}
