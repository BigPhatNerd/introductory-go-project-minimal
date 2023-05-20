package internal


type Service interface {
	AddPokemon(pokemon Pokemon) error
}

type PokemonService struct {
	repo Repository
}

func NewPokemonService(repo Repository) *PokemonService {
	return &PokemonService{
		repo: repo,
	}
}

func (s *PokemonService) AddPokemon(pokemon Pokemon) error {
	s.repo.AddPokemon(pokemon)
	return nil
}