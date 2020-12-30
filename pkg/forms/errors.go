package forms

type errors map[string][]string

// Add adds a new error to the validation map for a specific field
func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message from the list if it exists
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}