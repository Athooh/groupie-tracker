package api

func GetDates() ([]Date, error) {
	var dates []Date
	err := fetchData("/dates", &dates)
	return dates, err
}
