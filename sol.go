package sol

type ReachDataFormat struct {
	Diff int
	Min  int
	Max  int
}

func FindBestInOut(history []int) ReachDataFormat {
	result := ReachDataFormat{0, 0, 0}
	for idx := range history {
		tempResult := FindCurrentBest(history, idx)
		if result.Diff <= tempResult.Diff {
			result = tempResult
		}
	}
	return result
}

func FindCurrentBest(history []int, targetIdx int) ReachDataFormat {
	result := ReachDataFormat{0, 0, 0}
	for idx, val := range history {
		Diff := diff(val, history[targetIdx])
		Max := max(idx, targetIdx)
		Min := min(idx, targetIdx)
		if result.Diff <= Diff && history[Max] > history[Min] {
			result.Diff = Diff
			if Max == idx {
				result.Max = idx
				result.Min = targetIdx
			} else if idx < targetIdx {
				result.Max = targetIdx
				result.Min = idx
			}
		}
	}

	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if b > a {
		return a
	}
	return b
}
func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
