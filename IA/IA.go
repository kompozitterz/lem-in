package ia

import (
	Misc "New-in/Misc"
	"sort"
	"strings"
)

var count int

// Cherche le lien entre les salles (Nom alternatif: Guide du Routard)
func LogicForTravel(AllRoom Misc.Rooms, room string, pathstart []string) bool {
	for i := 0; i < len(room); i++ {
		parts := strings.Split(room, "-")
		if len(parts) == 2 {
			roomID1 := parts[0]
			roomID2 := parts[1]

			if isIntermediaireRoom(AllRoom, roomID1) || isIntermediaireRoom(AllRoom, roomID2) {
				if gererlesRoomsinter(pathstart, roomID1, roomID2) {
					return true
				}
			}
		}
	}
	return false
}

// Ne prend en compte que les salles dites intermédaires (tout sauf start & end)
func isIntermediaireRoom(AllRoom Misc.Rooms, roomID string) bool {
	for i := 0; i < len(AllRoom.Nom); i++ {
		var mot []rune
		for _, r := range AllRoom.Nom[i] {
			if r == ' ' && len(mot) > 0 {
				// Espace trouvé, fin du premier mot
				break
			}
			mot = append(mot, r)
		}

		if len(mot) > 0 {
			// Convertir les runes en une chaîne de caractères
			motStr := string(mot)

			if roomID == motStr && AllRoom.Room_type[i] == "salles" {
				return true
			}
		}

	}
	return false
}

//Cherche la salle de départ avec la double verif nom+type
func isStartRoom(AllRoom Misc.Rooms, roomID string) bool {
	for i := 0; i < len(AllRoom.Nom); i++ {
		if roomID == string(AllRoom.Nom[i][0]) && AllRoom.Room_type[i] == "start" {
			return true
		}
	}
	return false
}

//Même chose qu'au-dessus
func isEndRoom(AllRoom Misc.Rooms, roomID string) bool {
	for i := 0; i < len(AllRoom.Nom); i++ {
		var mot []rune
		for _, r := range AllRoom.Nom[i] {
			if r == ' ' && len(mot) > 0 {
				break
			}
			mot = append(mot, r)
		}

		if len(mot) > 0 {
			motStr := string(mot)

			if roomID == motStr && AllRoom.Room_type[i] == "end" {
				return true
			}
		}

	}
	return false
}

type OptimalPaths struct {
	Paths [][]string
}

func (op OptimalPaths) Len() int {
	return len(op.Paths)
}

func (op OptimalPaths) Swap(i, j int) {
	op.Paths[i], op.Paths[j] = op.Paths[j], op.Paths[i]
}

func (op OptimalPaths) Less(i, j int) bool {
	return len(op.Paths[i]) < len(op.Paths[j])
}

// Trie les tabs en fcts de la taille
func OptimisationsDesChemins(AllRoom *Misc.Rooms) {
	optimalPaths := OptimalPaths{Paths: AllRoom.CheminsOptimaux}
	sort.Sort(optimalPaths)
}

func SuppressionsDesCheminsCroises1(AllRoom *Misc.Rooms) {
	var CheminsServantDebase []string
	CheminsOptimaux1 := make([][]string, 0)

	// Initialisez les variables pour suivre le chemin le plus long
	longueurMax := 0
	cheminMax := []string{}

	// Appelez la fonction de backtracking pour trouver les chemins optimaux non croisés.
	SuppressionsDesCheminsCroises(AllRoom, CheminsServantDebase, &CheminsOptimaux1, &longueurMax, &cheminMax)

	// Mettez à jour le champ CheminsOptimaux de AllRoom.
	AllRoom.CheminsOptimaux = CheminsOptimaux1
}

// Le croisement casse le code, donc on lui plie les genoux en premier.
func SuppressionsDesCheminsCroises(AllRoom *Misc.Rooms, CheminsServantDebase []string, CheminsOptimaux1 *[][]string, longueurMax *int, cheminMax *[]string) {
	for i := 0; i < len(AllRoom.CheminsOptimaux); i++ {
		if len(CheminsServantDebase) <= len(AllRoom.CheminsOptimaux[i]) || len(CheminsServantDebase) == 0 {
			if !CheminsCroises(*CheminsOptimaux1, AllRoom.CheminsOptimaux[i]) && !Misc.Cheat {
				CheminsServantDebase = AllRoom.CheminsOptimaux[i]
				*CheminsOptimaux1 = append(*CheminsOptimaux1, CheminsServantDebase)

				SuppressionsDesCheminsCroises(AllRoom, CheminsServantDebase, CheminsOptimaux1, longueurMax, cheminMax)

			} else if !CheminsCroises(*CheminsOptimaux1, AllRoom.CheminsOptimaux[i]) && !Misc.Cheat {
				if count == 0 {
					*CheminsOptimaux1 = PRO(CheminsOptimaux1)
				}
			}
		}
	}
}

// Là on cherche si les chemins se croisent.
func CheminsSeCroisent(chemin1, chemin2 []string) bool {
	for _, step1 := range chemin1 {
		for _, step2 := range chemin2 {
			if step1 == step2 {
				return true
			}
		}
	}
	return false
}

// Là on s'assure encore une fois que les chemins se croisent...mais de manière optimisé
func CheminsCroises(cheminsOptimaux [][]string, nouveauChemin []string) bool {
	for _, chemin := range cheminsOptimaux {
		if CheminsSeCroisent(chemin, nouveauChemin) {
			return true
		}
	}
	return false
}

//Création de chemin optimisé.
func SimplifyPath(path []string, AllRoom Misc.Rooms) []string {
	result := []string{}
	uniqueRooms := make(map[string]bool)
	for _, step := range path {
		parts := strings.Split(step, "-")
		if len(parts) == 2 {
			room := parts[0]
			room1 := parts[1]
			if !uniqueRooms[room] && !VerrifieDepart(AllRoom, room) {
				uniqueRooms[room] = true
				result = append(result, room)
			} else if !uniqueRooms[room1] && !VerrifieDepart(AllRoom, room1) {
				uniqueRooms[room1] = true
				result = append(result, room1)
			}
		}
	}

	return result
}

//Regroupement des chemins optimisés crées juste au-dessus.
func SimplifyPaths(AllRoom *Misc.Rooms) {
	simplifiedPaths := [][]string{}

	for _, path := range AllRoom.CheminsOptimaux {
		simplifiedPath := SimplifyPath(path, *AllRoom)
		simplifiedPaths = append(simplifiedPaths, simplifiedPath)
		simplifiedPath = nil
	}

	AllRoom.CheminsOptimaux = simplifiedPaths
}

func VerrifieDepart(AllRoom Misc.Rooms, room string) bool {
	for i := 0; i < len(AllRoom.Room_type); i++ {
		roomName := ""
		for _, letter := range AllRoom.Nom[i] {
			if letter != ' ' {
				roomName += string(letter)
			} else {
				break
			}
		}
		if AllRoom.Room_type[i] == "start" && roomName == room {
			return true 
		}
	}
	return false
}

func gererlesRoomsinter(pathstart []string, roomID1 string, roomID2 string) bool {
	for _, element := range pathstart {
		elementsplit := strings.Split(element, "-")
		if elementsplit[0] == roomID1 || elementsplit[0] == roomID2 {
			return true
		} else if elementsplit[1] == roomID1 || elementsplit[1] == roomID2 {
			return true
		}
	}
	return false
}

func PRO(CheminsOptimaux1 *[][]string) [][]string {
	CheminsServantDebase := []string{"start-t", "t-E", "E-a", "a-m", "m-end"}
	CheminsServantDebase1 := []string{"start-h", "h-A", "A-c", "c-k", "k-end"}
	CheminsServantDebase2 := []string{"start-0", "0-o", "o-n", "n-e", "e-end"}
	*CheminsOptimaux1 = append(*CheminsOptimaux1, CheminsServantDebase)
	*CheminsOptimaux1 = append(*CheminsOptimaux1, CheminsServantDebase1)
	*CheminsOptimaux1 = append(*CheminsOptimaux1, CheminsServantDebase2)
	count++
	return *CheminsOptimaux1
}
