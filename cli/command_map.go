// MAP and MAPB CLI Commands

package cli

import (
	"github.com/Jhaendar/pokedexcli/pokeapi"
)

func init() {
	var SharedMapState = &MapState{}

	mapFCLICommand := MapFCLICommand{
		CLICommandInfo{
			name:        "map",
			description: "Display 20 locations",
		},
		SharedMapState,
	}

	addCLICommand("map", &mapFCLICommand)

	mapbCLICommand := MapBCLICommand{
		CLICommandInfo{
			name:        "mapb",
			description: "Display previous 20 locations",
		},
		SharedMapState,
	}

	addCLICommand("mapb", &mapbCLICommand)
}

type MapState struct {
	Next     *string
	Previous *string
}

func (m *MapState) UpdateMapState(next, previous *string) {
	m.Next = next
	m.Previous = previous
}

type MapFCLICommand struct {
	CLICommandInfo
	MapState *MapState
}

func (m *MapFCLICommand) Execute() error {
	var url string
	if m.MapState.Next != nil {
		url = *m.MapState.Next
	}

	mapIndex, err := pokeapi.GetLocationAreaIndex(url)
	if err != nil {
		return err
	}

	m.MapState.UpdateMapState(mapIndex.Next, mapIndex.Previous)
	pokeapi.PrintResourceNames(mapIndex.Results)

	return nil
}

type MapBCLICommand struct {
	CLICommandInfo
	MapState *MapState
}

func (m *MapBCLICommand) Execute() error {
	var url string
	if m.MapState.Next != nil {
		url = *m.MapState.Next
	}
	mapIndex, err := pokeapi.GetLocationAreaIndex(url)
	if err != nil {
		return err
	}

	m.MapState.UpdateMapState(mapIndex.Next, mapIndex.Previous)
	pokeapi.PrintResourceNames(mapIndex.Results)

	return nil
}
