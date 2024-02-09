package ia

import (
	Misc "New-in/Misc"
	"fmt"
	"os"
	"strings"
)

var (
	ChangeLimite bool
	EndIsNothing bool
)

// Trouve tout les chemins possibles
func LesCheminsTrier(AllRoom Misc.Rooms) Misc.Rooms {
	var startRooms []string
	var endroom []string
	startRooms, endroom = Misc.Trouversallefinetdebut(AllRoom)
	EndIsNothing = true

	for a := 0; a < len(endroom); a++ {
		split := strings.Split(endroom[a], "-")
		if isEndRoom(AllRoom, split[0]) || isEndRoom(AllRoom, split[1]) {
			EndIsNothing = false
		}
	}

	if EndIsNothing{
		fmt.Println()
		fmt.Println("C'est la fin... car il n'y a pas de fin mdr : Aucun chemins n'as de salle de type <fin>")
		os.Exit(12)
	}

	// Créer des chemins en fct du nombre de salles de départs et ainsi créer des chemins pour toutes les salles de départs
	for _, startRoom := range startRooms {
		usedRooms := make(map[string]bool)
		pathstart := []string{startRoom}
		AllRoom = backtrack(pathstart, endroom, AllRoom.Chemins, usedRooms, startRooms)
	}
	CheminFinEtStart(&AllRoom)

	OptimisationsDesChemins(&AllRoom)

	SuppressionsDesCheminsCroises1(&AllRoom)

	fmt.Println()
	//fmt.Println("Chemins possibles sans croissements :", AllRoom.CheminsOptimaux)
	SimplifyPaths(&AllRoom)

	return AllRoom
}

// Va chercher tout les chemins possibles c'est la func la plus importante !!!
func backtrack(pathstart []string, endRooms []string, rooms []string, used map[string]bool, startRooms []string) Misc.Rooms {
	EndIsNothing = false
	if len(pathstart) >= 1 {
		for _, endRoom := range endRooms {
			if pathstart[len(pathstart)-1] == endRoom {
				// Si le chemin a atteint "end", ajoutez-le à la liste des chemins
				Misc.AllRoom.CheminsOptimaux = append(Misc.AllRoom.CheminsOptimaux, append([]string(nil), pathstart...))
				return Misc.AllRoom
			} else if pathstart[0] == endRoom {

				Misc.AllRoom.CheminsOptimaux = append(Misc.AllRoom.CheminsOptimaux, append([]string(nil), pathstart...))
				return Misc.AllRoom

			}
		}
	}

	lastRoom := pathstart[len(pathstart)-1]

	for _, room := range rooms {
		// Vérifiez si la salle n'a pas déjà été utilisée et qu'elle n'est pas déjà dans le chemin
		if !used[room] && !contains(startRooms, room) && room != lastRoom {
			if LogicForTravel(Misc.AllRoom, room, pathstart) {
				// Ajoutez la salle au chemin
				if len(pathstart) < 7 && ChangeLimite == false {
					pathstart = append(pathstart, room)
					used[room] = true

					// Récursivement, explorez le chemin suivant
					backtrack(pathstart, endRooms, rooms, used, startRooms)

					// Retirez la salle du chemin pour explorer d'autres possibilités
					pathstart = pathstart[:len(pathstart)-1]
					used[room] = false
				} else if ChangeLimite == true {
					pathstart = append(pathstart, room)
					used[room] = true

					// Récursivement, explorez le chemin suivant
					backtrack(pathstart, endRooms, rooms, used, startRooms)

					// Retirez la salle du chemin pour explorer d'autres possibilités
					pathstart = pathstart[:len(pathstart)-1]
					used[room] = false
				}
			}
		}
	}

	return Misc.AllRoom
}

// Voit si il contient une salle spécifique (ici les salles de départ)
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func CheminFinEtStart(AllRoom *Misc.Rooms) {
	for i := 0; i < len(AllRoom.Chemins); i++ {

		split := strings.Split(AllRoom.Chemins[i], "-")

		if (isStartRoom(Misc.AllRoom, split[0]) || isStartRoom(Misc.AllRoom, split[1])) && (isEndRoom(Misc.AllRoom, split[0]) || isEndRoom(Misc.AllRoom, split[1])) {

			spliteur := strings.Split(AllRoom.Chemins[i], " ")
			AllRoom.CheminsOptimaux = append(AllRoom.CheminsOptimaux, spliteur)
		}

	}
}
