package dates

const DaysInWeek = 7

func WeekToDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysToWeek(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
