package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := 0; j < len(podium); j++ {
			if i < j {
				top := podium[i]
				podium[i] = podium[j]
				podium[j] = top
			}
		}
	}
	return podium
}
