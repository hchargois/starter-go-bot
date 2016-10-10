package command

import (
	"salad/store"
	"strings"
	"fmt"
	"strconv"
	"bytes"
)

func ParseCommandLine(cmd string) (string, []string) {
	f := strings.Fields(cmd)
	if len(f) == 0 {
		return "help", []string{}
	}
	return f[0], f[1:]
}


var COMMANDS = map[string]func([]string) string {
	"help": HelpCmd,
	"aide": HelpCmd,
	"surprise": SaladCmd,
	"ingredients": IngredientsCmd,
	"ajouter": AddIngredientCmd,
	"enlever": RemoveIngredientCmd,
}


func HelpCmd(args []string) string {
	return "Usage:\n" +
	"surprise [[n] catégorie]...: crée une salade surprise avec les catégories d'ingrédients demandées\n" +
	"  ex. surprise base 2 vert 2 bleu sauce extra\n" +
	"ingredients: liste les ingrédients disponibles\n" +
	"ajouter <catégorie> <ingrédient>: ajoute un ingrédient à la catégorie\n" +
	"enlever <catégorie> <ingrédient>: l'inverse de la commande précédente"
}


func SaladCmd(args []string) string {
	st := store.GetStore()
	if len(args) == 0 {
		return "Empty salad :("
	}

	var buf bytes.Buffer
	multiplier := 1
	for _, arg := range(args) {
		i, err := strconv.Atoi(arg)
		if err == nil {
			multiplier = i
			continue
		}

		for i:=0; i<multiplier; i++ {
			randIng, _ := st.Ingredients().GetRandom(arg)
			buf.WriteString(fmt.Sprintf("%v ", randIng))
		}

		multiplier = 1
	}

	return buf.String()
}


func IngredientsCmd(args []string) string {
	ings := store.GetStore().Ingredients()
	return ings.String()
}


func AddIngredientCmd(args []string) string {
	st := store.GetStore()
	ings := st.Ingredients()

	if len(args) < 2 {
		return "Usage: ajouter <categorie> <ingredient>"
	}

	ing := strings.Join(args[1:], " ")

	ings.Add(args[0], ing)
	st.Save()
	return fmt.Sprintf("Ingrédient %v ajouté dans la catégorie %v !", ing, args[0])
}


func RemoveIngredientCmd(args []string) string {
	st := store.GetStore()
	ings := st.Ingredients()

	if len(args) < 2 {
		return "Usage: enlever <categorie> <ingredient>"
	}

	ing := strings.Join(args[1:], " ")

	ings.Remove(args[0], ing)
	st.Save()
	return fmt.Sprintf("Ingrédient %v enlevé de la catégorie %v !", ing, args[0])
}


func ExecuteCommandLine(cmdLine string) string {
	cmdVerb, args := ParseCommandLine(cmdLine)
	cmdFunc, ok := COMMANDS[cmdVerb]
	if !ok {
		return "Lapin compris!"
	}
	return cmdFunc(args)
}
