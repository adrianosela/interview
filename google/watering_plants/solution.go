package main

/*
--------------------- PROBLEM: --------------------------------
- You and your friends are both watering plants.

- Plants are ordered in an array of integers indicating
  how much water a given plant needs.

- You start from the leftmost plant, your friend
  from the rightmost.

- You start with a waterbottle with capacity1,
  your friend with a waterbottle with capacty2.

- You both start with capacity of 0, meaning you will need to
  fill up your water bottle at the beginning of the process.

- It is possible that you will meet in the middle on the
  same plant. In this case, if you and your friend together
  have enough water, you may water the plant, otherwise only
  one of you has to refill their bottle

- You must have the full amount of water required to water
  a plant to begin watering it, otherwise you must refill.

- Write a function refillsNeeded(p []int, c1 int, c2 int)
  that returns the amount of times you and your friend will
  have to refill your water bottles
--------------------------------------------------------------
*/

import "fmt"

func main() {
	run1 := refillsNeeded([]int{2, 4, 5, 1, 2}, 5, 7)
	fmt.Printf("refillsNeeded([]int{2, 4, 5, 1, 2}, 5, 7) - expected 3, got %d\n", run1)

	run2 := refillsNeeded([]int{2, 4, 5, 1, 1, 2}, 11, 50)
	fmt.Printf("refillsNeeded([]int{2, 4, 5, 1, 2}, 11, 50) - expected 2, got %d\n", run2)

	run3 := refillsNeeded([]int{2, 2, 2, 6, 2, 2, 2}, 6, 6)
	fmt.Printf("refillsNeeded([]int{2, 2, 2, 1, 2, 2, 2}, 6, 6) - expected 3, got %d\n", run3)
}

func refillsNeeded(plants []int, capacity1, capacity2 int) int {
	w1, w2, refs := 0, 0, 0

	// we only iterate through the first half
	// our friend looks at the other half
	for you := 0; you < len(plants)/2; you++ {
		if w1 < plants[you] {
			w1 = capacity1
			refs++
		}
		if w2 < plants[len(plants)-you-1] {
			w2 = capacity2
			refs++
		}
		w1 -= plants[you]
		w2 -= plants[len(plants)-you-1]
	}

	// if there is an odd number of plants
	// you and your friend meet in the middle
	if len(plants)%2 != 0 {
		if w1+w2 >= plants[len(plants)/2+1] {
			return refs
		}
		return refs + 1
	}

	return refs
}
