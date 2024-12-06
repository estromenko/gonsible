package inventory

import (
	"os"
	"slices"

	"github.com/pelletier/go-toml/v2"
)

type InventoryGroup struct {
	Name  string   `toml:"name"`
	Hosts []string `toml:"hosts"`
}

type Inventory struct {
	Groups []InventoryGroup `toml:"groups"`
}

func New(inventoryPath string) (*Inventory, error) {
	var inventory Inventory
	rawData, err := os.ReadFile(inventoryPath)
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(rawData, &inventory); err != nil {
		return nil, err
	}
	return &inventory, err
}

func (i *Inventory) GetHostsByGroups(groupNames []string) []string {
	var hosts []string
	for _, group := range i.Groups {
		if !slices.Contains(groupNames, group.Name) {
			continue
		}
		for _, host := range group.Hosts {
			hosts = append(hosts, host)
		}
	}
	return hosts
}
