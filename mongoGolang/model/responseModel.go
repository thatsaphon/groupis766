package model

type MenuResponseModel struct {
	Menus   []map[string]interface{} `json:"menues"`
	Message string                   `json:"message"`
}
type RecipeResponseModel struct {
	Recipes []map[string]interface{} `json:"recipes"`
	Message string                   `json:"message"`
}
