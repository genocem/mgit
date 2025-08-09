package logic

func DeleteElemsFromSlice(slice []string, elems []string) []string {
	if len(elems) == 0 {
		return slice
	}
	hashmap := make(map[string]int)
	for _, e := range slice {
		hashmap[e] = 1
	}
	for _, elem := range elems {
		hashmap[elem] = 0
	}
	var filtered []string
	for key, value := range hashmap {
		if value == 1 {
			filtered = append(filtered, key)
		}
	}
	return filtered
}
