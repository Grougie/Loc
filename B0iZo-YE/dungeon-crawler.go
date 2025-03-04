package main
import "fmt"

//  --Variables utilitaires décrivant les commandes

// Chaque déplacement d'une case coûte un point de fatigue
var MoveUp = uint32(0)
var MoveLeft = uint32(1)
var MoveDown = uint32(2)
var MoveRight = uint32(3)

// La téléportation ramène l'héroïne à sa position initiale dans la salle, mais coûte 10 points de fatigue
var Teleport = uint32(4)

// Les sauts vous déplacent de deux cases mais coûtent 4 points de fatigue
var JumpUp = uint32(5)
var JumpLeft = uint32(6)
var JumpDown = uint32(7)
var JumpRight = uint32(8)

var currentIndex = -1
var actions = []uint32{0, 0, 0, 1, 2, 1, 0, 0, 4, 1, 1, 1}


func NextMove(room string) uint32 {
	roomArray := []rune(room)
	width := 12

    // Trouver la position de 'P' dans roomArray
    var pos int
    found := false
    for i, r := range roomArray {
        if r == 'P' {
            pos = i
            found = true
            break
        }
    }

    if !found {
        fmt.Println("P non trouvé dans le tableau")
        return 4
    }

    // Vérifier les cases autour
    fmt.Println("Position de P :", pos)

    // Vérifier Haut
	if pos-width >= 0 {
		fmt.Println("Haut :", string(roomArray[pos-width]))
		return 0
	} else if pos+width < len(roomArray) {
		fmt.Println("Bas :", string(roomArray[pos+width]))
		return 2
	} else if pos%width != 0 {
		fmt.Println("Gauche :", string(roomArray[pos-1]))
		return 1
	} else if (pos+1)%width != 0 {
		fmt.Println("Droite :", string(roomArray[pos+1]))
		return 3
	}

	return 4
}
