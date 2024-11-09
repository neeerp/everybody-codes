package main

func numPotionsNeeded(monsters string) int {
	potions := 0
	for _, mtype := range monsters {
    potions += potionsNeededFor(mtype)
	}

	return potions
}

func numPotionsNeededP2(monsters string) int {
	potions := 0
  for i := 0; i < len(monsters); {
    potions += potionsNeededFor(rune(monsters[i]))
    potions += potionsNeededFor(rune(monsters[i+1]))

    

    if monsters[i] != 'x' && monsters[i+1] != 'x' {
      potions += 2
    }

    i += 2
  }

	return potions
}

func numPotionsNeededP3(monsters string) int {
	potions := 0
  for i := 0; i < len(monsters); {
    potions += potionsNeededFor(rune(monsters[i]))
    potions += potionsNeededFor(rune(monsters[i+1]))
    potions += potionsNeededFor(rune(monsters[i+2]))

    
    monstersInGroup := 0
    if monsters[i] != 'x' {
      monstersInGroup += 1
    }

    if monsters[i+1] != 'x' {
      monstersInGroup += 1
    }

    if monsters[i+2] != 'x' {
      monstersInGroup += 1
    }

    switch monstersInGroup {
      case 2: potions += 2
      case 3: potions += 6
    }

    i += 3
  }

	return potions
}

func potionsNeededFor(mtype rune) int {
  switch mtype {
    case 'A': return 0
    case 'B': return 1
    case 'C': return 3
    case 'D': return 5
    default:  return 0
  }
}
