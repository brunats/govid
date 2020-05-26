package processing

import "github.com/brunats/govid/providers"

// Processing generates data
func Processing(data providers.Data) providers.Data {
	data.Processing.MortalityRate = calculateMortalityRate(data.Confirmed, data.Deaths)

	return data
}

func calculateMortalityRate(confirmed int, deaths int) int {
	return ((confirmed * 100) / deaths)
}
