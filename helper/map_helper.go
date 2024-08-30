package helper

func GetFromMapWithDefaultValue(mmap map[string]string, key string, defaultValue string) string {
	val, ok := mmap[key]
	if !ok {
		return defaultValue
	}
	return val
}
