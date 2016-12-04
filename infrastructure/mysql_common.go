package infrastructure

// Draft funcion that call to mysql and response data
func Get(tbl, id string) (interface{}, error) {
	return map[string]string{
		"id": id,
	}, nil
}
