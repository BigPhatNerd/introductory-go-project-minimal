package internal

type Repository interface {
	AddPokemon(pokemon Pokemon)
}

type PokemonRepository struct {
	pokedex []Pokemon
}

func NewPokemonRepository() *PokemonRepository {
	return &PokemonRepository{
		pokedex: []Pokemon{}
	}
}

func (r *PokemonRepository) AddPokemon(pokemon Pokemon){
	r.pokedex = append(r.pokedex, pokemon)
}

func SavePokedexToFile(filename string, repo *PokemonRepository) error {
	pokedexData, err := json.Marshal(repo.pokedex)
	if err != nil {
		return fmt.Errorf("failed to marshal pokedex data: %w", err)
	}
	err = ioutil.WriteFile(filename, pokedexData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write pokedex file: %w", err)
	}
	return nil
}
