package pipeline

import (
	"log/slog"
	"os"
	"sync"

	"github.com/estromenko/gonsible/internal/inventory"
	"github.com/estromenko/gonsible/internal/ssh"
	"github.com/pelletier/go-toml/v2"
)

type Step struct {
	Description string `toml:"description"`
	Cmd         string `toml:"cmd"`
}

type Pipeline struct {
	Name   string   `toml:"name"`
	Groups []string `toml:"groups"`
	Steps  []Step   `toml:"steps"`
}

func New(pipelinePath string) (*Pipeline, error) {
	source, err := os.ReadFile(pipelinePath)
	if err != nil {
		return nil, err
	}

	var pipeline Pipeline
	err = toml.Unmarshal(source, &pipeline)

	return &pipeline, err
}

func (p *Pipeline) Execute(invent *inventory.Inventory) error {
	hosts := invent.GetHostsByGroups(p.Groups)

	var wg sync.WaitGroup

	for _, host := range hosts {
		wg.Add(1)
		go func() {
			slog.Info("start-pipeline", "name", p.Name)

			defer wg.Done()
			for _, step := range p.Steps {
				output, err := ssh.Execute(step.Cmd, host)
				if err != nil {
					slog.Error("error", "host", host, "pipeline", p.Name, "step", step.Description, "error", err, "output", string(output))
					break
				}
				slog.Info("execute-step", "host", host, "pipeline", p.Name, "step", step.Description, "output", string(output))
			}
		}()
	}

	wg.Wait()

	return nil
}
