package slices

func Deduplicate(slice []string) []string {
	j := 0
	for i := 1; i < len(slice); i++ {
		if slice[j] == slice[i] {
			continue
		}
		j++
		// preserve the original data
		// slice[i], slice[j] = slice[j], slice[i]
		// only set what is required
		slice[j] = slice[i]
	}
	return slice[:j+1]
}
