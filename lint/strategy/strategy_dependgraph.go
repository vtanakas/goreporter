package strategy

import (
	"goreporter/engine"
	"goreporter/linters/depend"
)

type StrategyDependGraph struct {
	Sync *engine.Synchronizer `inject:""`
}

func (s *StrategyDependGraph) GetName() string {
	return "DependGraph"
}

func (s *StrategyDependGraph) GetDescription() string {
	return "The dependency graph for all packages in the project helps you optimize the project architecture."
}

func (s *StrategyDependGraph) GetWeight() float64 {
	return 0.
}

// linterDependGraph is a function that builds the dependency graph of all packages in the
// project helps you optimize the project architecture.It will extract from the linter need
// to convert the data.The result will be saved in the r's attributes.
func (s *StrategyDependGraph) Compute(parameters engine.StrategyParameter) (summaries *engine.Summaries) {
	summaries = engine.NewSummaries()

	graph := depend.Depend(parameters.ProjectPath, parameters.ExceptPackages)
	summaries.Summaries["graph"] = engine.Summary{
		Name:        s.GetName(),
		Description: graph,
	}

	return
}

func (s *StrategyDependGraph) Percentage(summaries *engine.Summaries) float64 {
	return 0.
}
