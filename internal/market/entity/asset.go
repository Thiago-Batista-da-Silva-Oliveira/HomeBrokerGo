package entity

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

func NewAsset(id string, name string, marketVolumn int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolumn,
	}
}
