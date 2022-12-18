package slices

//func Deduplicate(slice []string) []string {
//	j := 0
//	for i := 1; i < len(slice); i++ {
//		if slice[j] == slice[i] {
//			continue
//		}
//		j++
//		// preserve the original data
//		// slice[i], slice[j] = slice[j], slice[i]
//		// only set what is required
//		slice[j] = slice[i]
//	}
//	return slice[:j+1]
//}

func Deduplicate(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveIndex(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
