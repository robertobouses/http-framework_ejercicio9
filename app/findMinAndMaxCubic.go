package app

func (s *Service) FindMinAndMaxCubic(cubics []int) (int, int) {
	if len(cubics) == 0 {
		return 0, 0
	}

	minCubic := cubics[0]
	maxCubic := cubics[0]

	for _, cubic := range cubics {
		if cubic < minCubic {
			minCubic = cubic
		}
		if cubic > maxCubic {
			maxCubic = cubic
		}
	}

	return minCubic, maxCubic
}
