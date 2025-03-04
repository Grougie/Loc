package main

// N'éditez pas ce fichier; éditez le fichier dungeon-crawler.go

import (
	// Si votre éditeur souligne cette ligne, vous pouvez l'ignorer!
	"syscall/js"
)

// Fonction assurant le lien avec la page Web
func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		input := args[0].String()
		action := NextMove(input)
		return action
	})
	return jsonFunc
}

// La fonction main ne doit pas être modifiée
func main() {
	js.Global().Set("next_move", jsonWrapper())
	<-make(chan struct{})
}
