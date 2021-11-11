package util

// WrapError returns map[string]interface{} which represents error
func WrapError(e error) map[string]interface{} {
	return map[string]interface{}{"error": e.Error()}
}
