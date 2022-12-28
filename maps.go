package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["Spain"] = "Spanish" 
	languages["France"] = "French"
	fmt.Println(languages["Spain"])
	fmt.Println(languages)

	scores := map[string]int {
		"Pedro" : 5,
		"Alberto" : 8,
		"Fernando" : 4}
	scores["Juan"] = 3
	scores["Fernando"] = 6
	fmt.Println(scores)

	delete(scores, "Pedro")
	fmt.Println(scores)

	for key, value := range scores {
		fmt.Printf("Person %s with score %d \n", key, value)
	}

	score, exists := scores["Juan"]
	fmt.Printf("%d %t \n", score, exists)
}