package processing

import (
	"github.com/brunats/govid/providers"
)

// Processing generates data
func Processing(data *providers.Data) {
	data.Processing.MortalityRate = calculateMortalityRate(data.Confirmed, data.Deaths)
}

func calculateMortalityRate(confirmed int, deaths int) int {
	if deaths == 0 {
		return 0
	}

	return ((deaths * 100) / confirmed)
}
