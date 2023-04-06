func convertTemperature(celsius float64) []float64 {
	kelv := celsius + 273.15
	fnhren := celsius*1.80 + 32.00

	var arr = [2]float64{kelv, fnhren}
	return arr
}