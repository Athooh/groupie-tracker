package api 

func GetLocations() ([]Location, error) {
    var locations []Location
    err := fetchData("/locations", &locations)
    return locations, err
}



