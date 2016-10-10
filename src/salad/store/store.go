package store

import (
	"encoding/json"
	"io/ioutil"
)

var store *StoreImpl

type Store interface {
	Ingredients() *IngredientList
	SetIngredients(*IngredientList)
	Save()
}

type StoreImpl struct {
	ingredients *IngredientList
}

func InitStore() Store {
	store = &StoreImpl{ingredients: &IngredientList{}}
	return store
}

func GetStore() Store {
	return store
}

func (st *StoreImpl) Ingredients() *IngredientList {
	return st.ingredients
}

func (st *StoreImpl) SetIngredients(ings *IngredientList) {
	st.ingredients = ings
}

type IngredientsJsonRepr map[string][]string

type StoreJsonRepr struct {
	Ingredients IngredientsJsonRepr
}

func (st *StoreImpl) Save() {
	ijr := IngredientsJsonRepr{}
	sjr := StoreJsonRepr{Ingredients: ijr}
	ingredients := *st.Ingredients()
	for cat := range(ingredients) {
		ijr[cat] = ingredients.List(cat)
	}
	b, _ := json.MarshalIndent(sjr, "", "  ")

	ioutil.WriteFile("data.json", b, 0644)
}

func Load() Store {
	st := InitStore()

	b, err := ioutil.ReadFile("data.json")
	if err != nil {
		return st
	}

	var sjr StoreJsonRepr
	json.Unmarshal(b, &sjr)

	ingredients := st.Ingredients()
	for cat, ings := range(sjr.Ingredients) {
		for _, ing := range(ings) {
			ingredients.Add(cat, ing)
		}
	}
	return st
}
