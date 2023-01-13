package model

type Game struct {
	ID       string  `json:"id"`
	NAME     string  `json:"name"`
	CATEGORY string  `json:"category"`
	PRICE    float32 `json:"price"`
}

var InitialData = []Game{
	{
		ID:       "1",
		NAME:     "God of War",
		CATEGORY: "Open World",
		PRICE:    45.0,
	},
	{
		ID:       "2",
		NAME:     "World War",
		CATEGORY: "Fantasy",
		PRICE:    25.0,
	},
	{
		ID:       "3",
		NAME:     "BattleFiled",
		CATEGORY: "Shooter",
		PRICE:    15.0,
	},
}
