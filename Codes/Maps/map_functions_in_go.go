/*
  Here we have checked whether the value is present in map
  if present then how to delete from map
*/
package main

import (
	"errors"
	"fmt"
	"sort"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {

	_, ok := users[name]
	if ok == false {
		return false, errors.New("not found")
	} else if users[name].scheduledForDeletion == false {
		return false, nil
	} else {
		delete(users, name)
		return true, nil
	}

}

// don't touch below this line

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted:", name)
		return
	}
	fmt.Println("Did not delete:", name)
}

func main() {
	users := map[string]user{
		"john": {
			name:                 "john",
			number:               18965554631,
			scheduledForDeletion: true,
		},
		"elon": {
			name:                 "elon",
			number:               19875556452,
			scheduledForDeletion: true,
		},
		"breanna": {
			name:                 "breanna",
			number:               98575554231,
			scheduledForDeletion: false,
		},
		"kade": {
			name:                 "kade",
			number:               10765557221,
			scheduledForDeletion: false,
		},
	}
	test(users, "john")
	test(users, "musk")
	test(users, "santa")
	test(users, "kade")

	keys := []string{}
	for name := range users {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	fmt.Println("Final map keys:")
	for _, name := range keys {
		fmt.Println(" - ", name)
	}
}