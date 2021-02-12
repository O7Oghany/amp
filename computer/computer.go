package computer

import (
	"github.com/abdelmmu/amp/base"
	"github.com/abdelmmu/amp/model"
)

//Computer implements the method adn the necessary data  to interact with computer endpoint
//more info: https://api-docs.amp.cisco.com/api_resources/Computer?api_host=api.eu.amp.cisco.com&api_version=v1
type Computer struct {
	Proxy base.AMP
	Model model.Computer
}

func NewComputer(auth string) *Computer{
	amp := base.NewAMP(auth)
	if amp == nil {
		return nil
	}
	return &Computer{
		Proxy: *amp,
		Model: model.Computer{},
	}
}

func (c * Computer) GetComputers() {

}