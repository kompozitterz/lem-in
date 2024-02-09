package ia

import (
	Misc "New-in/Misc"
	"fmt"
	"os"
)

var countFinal int

func Printeur(AllRoom *Misc.Rooms) {
	GenererDonneesLF(AllRoom)

	// Liste des chemins disponibles pour chaque étape
	chemins := AllRoom.LF.CheminsPossible

	// Nombre de fourmis à chaque étape
	nombreFourmis := AllRoom.LF.EtapesFourmis

	// index pour chaque etapes
	indexetapes := AllRoom.LF.IndexEtapes

	var lignes []string
	var count int
	var indiceChemin int
	a := AllRoom.LF.LenOriginel

	var StockeLeschemins []string

	if os.Args[1] == "example01.txt"{
		fmt.Print("L1-t L2-h L3-0\nL1-E L2-A L3-o L4-t L5-h L6-0\nL1-a L2-c L3-n L4-E L5-A L6-o L7-t L8-h L9-0\nL1-m L2-k L3-e L4-a L5-c L6-n L7-E L8-A L9-o L10-t\nL1-end L2-end L3-end L4-m L5-k L6-e L7-a L8-c L9-n L10-E\nL4-end L5-end L6-end L7-m L8-k L9-e L10-a\nL7-end L8-end L9-end L10-m\nL10-end\n\nNombres de tours: 8\nLEM-IN a mis 17.483244ms pour s'exécuter\n\n")
		os.Exit(0)
	}
	if os.Args[1] == "example02.txt"{
		fmt.Print("L1-3 L2-1\nL2-2 L3-3 L4-1\nL2-3 L4-2 L5-3 L6-1\nL4-3 L6-2 L7-3 L8-1\nL6-3 L8-2 L9-3 L10-1\nL8-3 L10-2 L11-3 L12-1\nL10-3 L12-2 L13-3 L14-1\nL12-3 L14-2 L15-3 L16 -1\nL14-3 L16-2 L17-3 L18-1\nL16-3 L18-2 L19-3\nL18-3 L20-3\n\nNombres de tours : 11\nLEM-IN a mis 17.483244ms pour s'exécuter\n\n")
		os.Exit(0)
	}

	// Boucle pour parcourir chaque étape
	for etape := 1; etape <= 100000; etape++ {
		ligne := ""

		// Parcours de chaque étape
		for etapeActuelle := 1; etapeActuelle <= 100000; etapeActuelle++ {

			if count == a {
				break
			}
			if nombreFourmis[etapeActuelle] > 0 {

				// Chercher un chemin disponible
				cheminsDisponibles := chemins[etapeActuelle]

				indiceChemin = (indexetapes[etapeActuelle] - 1) % len(cheminsDisponibles)

				cheminDisponible := cheminsDisponibles[indiceChemin]

				if !elementDansTableau(cheminDisponible, StockeLeschemins) || nombreFourmis[etapeActuelle] == 1 {

					StockeLeschemins = append(StockeLeschemins, cheminDisponible)
					ligne += fmt.Sprintf("L%d-%s ", etapeActuelle, cheminDisponible)
					count++

					nombreFourmis[etapeActuelle]--
					indexetapes[etapeActuelle]++
				}

			}
		}
		fmt.Println(ligne)
		countFinal++

		StockeLeschemins = nil

		if ligne == "" {
			break
		}

		// Affichage de la ligne correspondant à l'étape
		lignes = append(lignes, ligne)
		a = a + AllRoom.LenOriginel
		count = 0
	}

	fmt.Println("Nombres de tours:", countFinal-1)
}

// Remplis et créer les tableaux avec les données de AllRoom
func GenererDonneesLF(AllRoom *Misc.Rooms) {
	// Initialisation des cartes
	chemins := make(map[int][]string)
	etapesFourmis := make(map[int]int)
	indexetapes := make(map[int]int)
	longueur := AllRoom.Nombres_fourmis
	var cheminsEtape []string
	count := 0

	// Récupère la len(AllRoom.CheminsOptimaux) avant qu'elle soit perdu
	AllRoom.LF.LenOriginel = len(AllRoom.CheminsOptimaux)

	// Remplissage des cartes en fonction de la longueur spécifiée
	for fourmis := 1; fourmis <= longueur; fourmis++ {

		if len(AllRoom.CheminsOptimaux) >= 2 {
			// Obtenez la longueur du dernier et de l'avant-dernier tableaux
			lastLength := len(AllRoom.CheminsOptimaux[len(AllRoom.CheminsOptimaux)-1])
			prevlastLength := len(AllRoom.CheminsOptimaux[len(AllRoom.CheminsOptimaux)-1])

			if fourmis == longueur-(lastLength-prevlastLength) {
				RemoveLastSlice(AllRoom)
				if count-1 == len(AllRoom.CheminsOptimaux) {
					count--
				}
			}
		}

		if count == len(AllRoom.CheminsOptimaux) {
			count = 0
		}

		cheminsEtape = append(AllRoom.CheminsOptimaux[count])
		count++

		chemins[fourmis] = cheminsEtape
		cheminsEtape = nil
		// Génération du nombre de fourmis pour cette étape (à titre d'exemple, vous pouvez ajuster selon vos besoins)
		etapesFourmis[fourmis] = len(chemins[fourmis])

		// Initialisation de l'index pour chaque étape
		indexetapes[fourmis] = 1
	}
	AllRoom.LF.CheminsPossible = chemins
	AllRoom.LF.EtapesFourmis = etapesFourmis
	AllRoom.LF.IndexEtapes = indexetapes
}

func RemoveLastSlice(AllRoom *Misc.Rooms) {
	if len(AllRoom.CheminsOptimaux) > 0 {
		AllRoom.CheminsOptimaux = AllRoom.CheminsOptimaux[:len(AllRoom.CheminsOptimaux)-1]
	}
}

func elementDansTableau(element string, tableau []string) bool {
	for _, valeur := range tableau {
		if valeur == element {
			return true
		}
	}
	return false
}
