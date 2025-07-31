package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	// ?
	if day == 0 {
		return nil
	}

	var getCost []float64
	for i := 1; i < len(costs); i++ {
		if costs[i].day == day {
			getCost = append(getCost, costs[i].value)
		}
	}
	return getCost
}
