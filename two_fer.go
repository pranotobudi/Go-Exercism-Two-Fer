// Package twofer implements string processing functions.
package twofer

// ShareWith function receive 1 string parameter 'name', and return a string.
// it will return 'One for X, one for me.' Where X is the given name.
//it will return 'One for you, one for me.' if the name is missing
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
