package pokeapi

type Resource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaIndex struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Resource `json:"results"`
}

type LocationArea struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	Location          Resource           `json:"location"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Resource `json:"pokemon"`
}

// Height, Weight, stats, Types
type Pokemon struct {
	// Abilities []struct {
	// 	Ability  Resource `json:"ability,omitempty"`
	// 	IsHidden bool     `json:"is_hidden,omitempty"`
	// 	Slot     int      `json:"slot,omitempty"`
	// } `json:"abilities,omitempty"`
	BaseExperience int `json:"base_experience,omitempty"`
	// Cries          struct {
	// 	Latest string `json:"latest,omitempty"`
	// 	Legacy string `json:"legacy,omitempty"`
	// } `json:"cries,omitempty"`
	// Forms       []Resource `json:"forms,omitempty"`
	// GameIndices []struct {
	// 	GameIndex int      `json:"game_index,omitempty"`
	// 	Version   Resource `json:"version,omitempty"`
	// } `json:"game_indices,omitempty"`
	Height int `json:"height,omitempty"`
	// HeldItems              []any  `json:"held_items,omitempty"`
	ID int `json:"id,omitempty"`
	// IsDefault              bool   `json:"is_default,omitempty"`
	// LocationAreaEncounters string `json:"location_area_encounters,omitempty"`
	// Moves                  []struct {
	// 	Move                Resource `json:"move,omitempty"`
	// 	VersionGroupDetails []struct {
	// 		LevelLearnedAt  int      `json:"level_learned_at,omitempty"`
	// 		MoveLearnMethod Resource `json:"move_learn_method,omitempty"`
	// 		VersionGroup    Resource `json:"version_group,omitempty"`
	// 	} `json:"version_group_details,omitempty"`
	// } `json:"moves,omitempty"`
	Name string `json:"name,omitempty"`
	// Order         int      `json:"order,omitempty"`
	// PastAbilities []any    `json:"past_abilities,omitempty"`
	// PastTypes     []any    `json:"past_types,omitempty"`
	Species Resource `json:"species,omitempty"`
	// Sprites       struct {
	// 	BackDefault      string `json:"back_default,omitempty"`
	// 	BackFemale       any    `json:"back_female,omitempty"`
	// 	BackShiny        string `json:"back_shiny,omitempty"`
	// 	BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
	// 	FrontDefault     string `json:"front_default,omitempty"`
	// 	FrontFemale      any    `json:"front_female,omitempty"`
	// 	FrontShiny       string `json:"front_shiny,omitempty"`
	// 	FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 	Other            struct {
	// 		DreamWorld struct {
	// 			FrontDefault string `json:"front_default,omitempty"`
	// 			FrontFemale  any    `json:"front_female,omitempty"`
	// 		} `json:"dream_world,omitempty"`
	// 		Home struct {
	// 			FrontDefault     string `json:"front_default,omitempty"`
	// 			FrontFemale      any    `json:"front_female,omitempty"`
	// 			FrontShiny       string `json:"front_shiny,omitempty"`
	// 			FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 		} `json:"home,omitempty"`
	// 		OfficialArtwork struct {
	// 			FrontDefault string `json:"front_default,omitempty"`
	// 			FrontShiny   string `json:"front_shiny,omitempty"`
	// 		} `json:"official-artwork,omitempty"`
	// 		Showdown struct {
	// 			BackDefault      string `json:"back_default,omitempty"`
	// 			BackFemale       any    `json:"back_female,omitempty"`
	// 			BackShiny        string `json:"back_shiny,omitempty"`
	// 			BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
	// 			FrontDefault     string `json:"front_default,omitempty"`
	// 			FrontFemale      any    `json:"front_female,omitempty"`
	// 			FrontShiny       string `json:"front_shiny,omitempty"`
	// 			FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 		} `json:"showdown,omitempty"`
	// 	} `json:"other,omitempty"`
	// 	Versions struct {
	// 		GenerationI struct {
	// 			RedBlue struct {
	// 				BackDefault      string `json:"back_default,omitempty"`
	// 				BackGray         string `json:"back_gray,omitempty"`
	// 				BackTransparent  string `json:"back_transparent,omitempty"`
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontGray        string `json:"front_gray,omitempty"`
	// 				FrontTransparent string `json:"front_transparent,omitempty"`
	// 			} `json:"red-blue,omitempty"`
	// 			Yellow struct {
	// 				BackDefault      string `json:"back_default,omitempty"`
	// 				BackGray         string `json:"back_gray,omitempty"`
	// 				BackTransparent  string `json:"back_transparent,omitempty"`
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontGray        string `json:"front_gray,omitempty"`
	// 				FrontTransparent string `json:"front_transparent,omitempty"`
	// 			} `json:"yellow,omitempty"`
	// 		} `json:"generation-i,omitempty"`
	// 		GenerationIi struct {
	// 			Crystal struct {
	// 				BackDefault           string `json:"back_default,omitempty"`
	// 				BackShiny             string `json:"back_shiny,omitempty"`
	// 				BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
	// 				BackTransparent       string `json:"back_transparent,omitempty"`
	// 				FrontDefault          string `json:"front_default,omitempty"`
	// 				FrontShiny            string `json:"front_shiny,omitempty"`
	// 				FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
	// 				FrontTransparent      string `json:"front_transparent,omitempty"`
	// 			} `json:"crystal,omitempty"`
	// 			Gold struct {
	// 				BackDefault      string `json:"back_default,omitempty"`
	// 				BackShiny        string `json:"back_shiny,omitempty"`
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontTransparent string `json:"front_transparent,omitempty"`
	// 			} `json:"gold,omitempty"`
	// 			Silver struct {
	// 				BackDefault      string `json:"back_default,omitempty"`
	// 				BackShiny        string `json:"back_shiny,omitempty"`
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontTransparent string `json:"front_transparent,omitempty"`
	// 			} `json:"silver,omitempty"`
	// 		} `json:"generation-ii,omitempty"`
	// 		GenerationIii struct {
	// 			Emerald struct {
	// 				FrontDefault string `json:"front_default,omitempty"`
	// 				FrontShiny   string `json:"front_shiny,omitempty"`
	// 			} `json:"emerald,omitempty"`
	// 			FireredLeafgreen struct {
	// 				BackDefault  string `json:"back_default,omitempty"`
	// 				BackShiny    string `json:"back_shiny,omitempty"`
	// 				FrontDefault string `json:"front_default,omitempty"`
	// 				FrontShiny   string `json:"front_shiny,omitempty"`
	// 			} `json:"firered-leafgreen,omitempty"`
	// 			RubySapphire struct {
	// 				BackDefault  string `json:"back_default,omitempty"`
	// 				BackShiny    string `json:"back_shiny,omitempty"`
	// 				FrontDefault string `json:"front_default,omitempty"`
	// 				FrontShiny   string `json:"front_shiny,omitempty"`
	// 			} `json:"ruby-sapphire,omitempty"`
	// 		} `json:"generation-iii,omitempty"`
	// 		GenerationIv struct {
	// 			DiamondPearl struct {
	// 				BackDefault      string `json:"back_default,omitempty"`
	// 				BackFemale       any    `json:"back_female,omitempty"`
	// 				BackShiny        string `json:"back_shiny,omitempty"`
	// 				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontFemale      any    `json:"front_female,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 			} `json:"diamond-pearl,omitempty"`
	// 			HeartgoldSoulsilver struct {
	// 				BackDefault      string `json:"back_default,omitempty"`
	// 				BackFemale       any    `json:"back_female,omitempty"`
	// 				BackShiny        string `json:"back_shiny,omitempty"`
	// 				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontFemale      any    `json:"front_female,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 			} `json:"heartgold-soulsilver,omitempty"`
	// 			Platinum struct {
	// 				BackDefault      string `json:"back_default,omitempty"`
	// 				BackFemale       any    `json:"back_female,omitempty"`
	// 				BackShiny        string `json:"back_shiny,omitempty"`
	// 				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontFemale      any    `json:"front_female,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 			} `json:"platinum,omitempty"`
	// 		} `json:"generation-iv,omitempty"`
	// 		GenerationV struct {
	// 			BlackWhite struct {
	// 				Animated struct {
	// 					BackDefault      string `json:"back_default,omitempty"`
	// 					BackFemale       any    `json:"back_female,omitempty"`
	// 					BackShiny        string `json:"back_shiny,omitempty"`
	// 					BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
	// 					FrontDefault     string `json:"front_default,omitempty"`
	// 					FrontFemale      any    `json:"front_female,omitempty"`
	// 					FrontShiny       string `json:"front_shiny,omitempty"`
	// 					FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 				} `json:"animated,omitempty"`
	// 				BackDefault      string `json:"back_default,omitempty"`
	// 				BackFemale       any    `json:"back_female,omitempty"`
	// 				BackShiny        string `json:"back_shiny,omitempty"`
	// 				BackShinyFemale  any    `json:"back_shiny_female,omitempty"`
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontFemale      any    `json:"front_female,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 			} `json:"black-white,omitempty"`
	// 		} `json:"generation-v,omitempty"`
	// 		GenerationVi struct {
	// 			OmegarubyAlphasapphire struct {
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontFemale      any    `json:"front_female,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 			} `json:"omegaruby-alphasapphire,omitempty"`
	// 			XY struct {
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontFemale      any    `json:"front_female,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 			} `json:"x-y,omitempty"`
	// 		} `json:"generation-vi,omitempty"`
	// 		GenerationVii struct {
	// 			Icons struct {
	// 				FrontDefault string `json:"front_default,omitempty"`
	// 				FrontFemale  any    `json:"front_female,omitempty"`
	// 			} `json:"icons,omitempty"`
	// 			UltraSunUltraMoon struct {
	// 				FrontDefault     string `json:"front_default,omitempty"`
	// 				FrontFemale      any    `json:"front_female,omitempty"`
	// 				FrontShiny       string `json:"front_shiny,omitempty"`
	// 				FrontShinyFemale any    `json:"front_shiny_female,omitempty"`
	// 			} `json:"ultra-sun-ultra-moon,omitempty"`
	// 		} `json:"generation-vii,omitempty"`
	// 		GenerationViii struct {
	// 			Icons struct {
	// 				FrontDefault string `json:"front_default,omitempty"`
	// 				FrontFemale  any    `json:"front_female,omitempty"`
	// 			} `json:"icons,omitempty"`
	// 		} `json:"generation-viii,omitempty"`
	// 	} `json:"versions,omitempty"`
	// } `json:"sprites,omitempty"`
	Stats []struct {
		BaseStat int      `json:"base_stat,omitempty"`
		Effort   int      `json:"effort,omitempty"`
		Stat     Resource `json:"stat,omitempty"`
	} `json:"stats,omitempty"`
	Types []struct {
		Slot int      `json:"slot,omitempty"`
		Type Resource `json:"type,omitempty"`
	} `json:"types,omitempty"`
	Weight int `json:"weight,omitempty"`
}
