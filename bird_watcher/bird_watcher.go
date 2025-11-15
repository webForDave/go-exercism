package birdwatcher


func TotalBirdCount(birdsPerDay []int) int{
	if len(birdsPerDay) == 0 {
		return 0
	}

	sum := 0
	for _, val := range birdsPerDay{
		sum += val
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int{

	start, end, sum := 0, week * 7, 0

	if week * 7 > 7 {
		start = (week * 7) - 7
	}
	
	for _, v := range birdsPerDay[start:end]{
		sum += v
	}
	
	return sum
}

func FixBirdCountLog(birdsPerDay []int) []int{
	for i, v := range birdsPerDay{
		if i % 2 == 0{
			birdsPerDay[i] = v + 1
		}
	}

	return birdsPerDay
}