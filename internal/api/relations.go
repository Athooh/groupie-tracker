package api

func GetRelations() ([]Relation, error) {
	var relations []Relation
	err := fetchData("/relations", &relations)
	return relations, err
}
