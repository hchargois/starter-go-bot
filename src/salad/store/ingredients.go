package store

import (
	"fmt"
	"math/rand"
	"bytes"
)

var INGREDIENTS = IngredientList{
	"vert": {
		"g1": 1,
		"g2": 1,
		"g3": 1,
		"g4": 1,
	},
	"bleu": {
		"b1": 1,
		"b2": 1,
		"b3": 1,
	},
}

type IngredientLister interface {
	Add(string, string)
	Remove(string, string)
}

type IngredientList map[string]map[string]int

func (il *IngredientList) Add(category, ingredient string) {
	cat, ok := (*il)[category]
	if !ok {
		(*il)[category] = map[string]int{ingredient: 1}
		return
	}
	cat[ingredient] = 1
}

func (il *IngredientList) Remove(category, ingredient string) {
	cat, ok := (*il)[category]
	if ! ok {
		return
	}
	delete(cat, ingredient)
	if len(cat) == 0 {
		delete(*il, category)
	}
}

func (il *IngredientList) GetRandom(category string) (string, bool) {
	cat, ok := (*il)[category]
	if ! ok {
		return "", false
	}
	idx := rand.Intn(len(cat))
	i := 0
	for ing := range(cat) {
		if i == idx {
			return ing, true
		}
		i++
	}
	panic("shouldn't happen")
}

func (il *IngredientList) List(category string) []string {
	var ret []string
	for ing := range((*il)[category]) {
		ret = append(ret, ing)
	}
	return ret
}

func (il *IngredientList) String() string {
	var buf bytes.Buffer
	for cat := range(*il) {
		buf.WriteString(fmt.Sprintf("%v:\n", cat))
		for _, ing := range(il.List(cat)) {
			buf.WriteString(fmt.Sprintf("  %v\n", ing))
		}
	}
	return buf.String()
}
