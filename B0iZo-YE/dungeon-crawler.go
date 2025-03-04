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
	/**
	/* Insérez votre code ici
	/*
	/*                        */
	currentIndex++
	fmt.Println("currentIndex", currentIndex)
	//return actions[currentIndex]
	return MoveUp
}
