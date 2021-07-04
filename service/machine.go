package service

type Machine struct {
	Outlets               Outlet                    `json:"outlets"`
	IngredientQuantityMap map[string]int            `json:"total_items_quantity"`
	Beverages             map[string]map[string]int `json:"beverages"`
}
