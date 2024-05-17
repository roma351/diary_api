package objects

import "strings"

func GetName(firstName *string, lastName *string) string {
	var name []string
	if lastName != nil {
		if *lastName != "" {
			name = append(name, *lastName)
		}
	}
	if firstName != nil {
		if *firstName != "" {
			name = append(name, *firstName)
		}
	}
	return strings.Join(name, " ")
}
