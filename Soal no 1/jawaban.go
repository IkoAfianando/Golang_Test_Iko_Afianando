package main

import "fmt"

func fungsiUnik(a, b []string) []string {

	uniqueElements := make(map[string]bool)

	for _, elem := range a {
		uniqueElements[elem] = true
	}

	for _, elem := range b {
		uniqueElements[elem] = true
	}

	var result []string
	for elem := range uniqueElements {
		result = append(result, elem)
	}

	return result

}

func main() {
	// should print Adit, Elga, Okta, Sari
	fmt.Println(fungsiUnik([]string{"Adit", "Elga", "Okta"}, []string{"Okta", "Sari", "Elga"}))
}
