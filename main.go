package main

import (
	"fmt"
	"time"
	ia "New-in/IA"
	Misc "New-in/Misc"
)

func main() {
	début := time.Now()

	// D'abord on ouvre le fichiers et récupère/trie les données.
	content := Misc.OpenFile()
	AllRoom := Misc.FoundNameAndCoordonnees(content)

	//Les chemins sont plus relous à trier entre les possibles, les impossibles et les optis.
	AllRoom = ia.LesCheminsTrier(AllRoom)

	if ia.EndIsNothing {
		return
	}

	ia.Printeur(&AllRoom)
	fin := time.Now()
	tempsDExecution := fin.Sub(début)

	fmt.Printf("LEM-IN a mis %v pour s'exécuter.\n", tempsDExecution)
	fmt.Println()
}
