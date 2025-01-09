package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	mergers := make([][]int, 0, len(points))
	for _, point := range points {
		if len(mergers) == 0 || mergers[len(mergers)-1][1] < point[1] {
			mergers = append(mergers, point)
		} else {
			mergers[len(mergers)-1][1] = max(mergers[len(mergers)-1][1], point[1])
		}
	}
	return len(mergers)
}
