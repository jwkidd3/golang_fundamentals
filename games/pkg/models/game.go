package models

type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}

var Games []Game

func init() {

	Games = append(Games, Game{ID: "1", Name: "Hearthstone", Publisher: "Blizzard"})
	Games = append(Games, Game{ID: "2", Name: "Overwatch", Publisher: "Blizzard"})
	Games = append(Games, Game{ID: "3", Name: "Dota 2", Publisher: "Valve"})
	Games = append(Games, Game{ID: "4", Name: "GTA V", Publisher: "RockStar"})
	Games = append(Games, Game{ID: "5", Name: "DOOM Eternal", Publisher: "Bethesda"})

}
