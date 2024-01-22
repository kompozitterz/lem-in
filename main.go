package main

import (
	Interface "Lem-IN/Crea_interface"
	ia "Lem-IN/IA"
	func_cool "Lem-IN/func_utiles"
	"errors"

	"fmt"
)

func main() {

	// Début du setup du code
	content := func_cool.OpenFile()

	AllRoom := func_cool.FoundNameAndCoordonnees(content)

	Interface.CreaSallesInterface(AllRoom)
	// Fin de la configuration

	// Vérifier le nombre de fourmis
	if err := verifierNombreFourmis(AllRoom.Nombres_fourmis); err != nil {
		fmt.Println("Erreur :", err)
		return // Arrête le programme en cas d'erreur
	}

	AllRoom = ia.LesCheminsTrier(AllRoom)

	fmt.Println()
	fmt.Println("Nombres de fourmis:", AllRoom.Nombres_fourmis)

	fmt.Println("Etapes:", AllRoom.CheminsOptimaux)
	fmt.Println()
	ia.Printeur(&AllRoom)
}

func verifierNombreFourmis(nombresFourmis int) error {
	if nombresFourmis == 0 {
		return errors.New("Le nombre de fourmis ne peut pas être égal à zéro")
	}
	return nil
}
