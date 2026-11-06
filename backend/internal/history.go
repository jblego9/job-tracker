package internal

type MonthCount struct {
	Month string
	Count int
}

type History [12]MonthCount

func CreateEmptyHistory() History {
	var history History

	history[0].Month = "Jan"
	history[1].Month = "Feb"
	history[2].Month = "Mar"
	history[3].Month = "Apr"
	history[4].Month = "May"
	history[5].Month = "Jun"
	history[6].Month = "Jul"
	history[7].Month = "Aug"
	history[8].Month = "Sep"
	history[9].Month = "Oct"
	history[10].Month = "Nov"
	history[11].Month = "Dec"

	for i := range len(history) {
		history[i].Count = 0
	}

	return history
}
