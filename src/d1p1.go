package main

func numPotionsNeeded(monsters string) int {
	potions := 0
	for _, mtype := range monsters {
		switch mtype {
    case 'B':
      {
        potions += 1
      }
    case 'C':
      {
        potions += 3
      }
		}
	}

	return potions
}

func main() {

}
