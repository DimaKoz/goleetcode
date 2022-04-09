package easy

func areAlmostEqual(s1 string, s2 string) bool {

	diff := 0
	const maxDif = 2
	indexes := []int{0, 0}
	for i := 0; i < len(s1) && diff <= maxDif; i++ {
		if s1[i] != s2[i] {
			if diff < maxDif {
				indexes[diff] = i
			}
			diff++
		}
	}
	if diff != maxDif && diff != 0 {
		return false
	}
	return s1[indexes[0]] == s2[indexes[1]] && s2[indexes[0]] == s1[indexes[1]]
}
