package db

func Persist(db *Datasource) bool {
	return false
}

func Restore(version string) *Datasource {
	return New()
}
