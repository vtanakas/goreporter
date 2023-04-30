package strategy

import (
	"goreporter/engine"
	"goreporter/lint/linters/unittest"
	"goreporter/utils"
)

type StrategyImportPackages struct {
	Sync *engine.Synchronizer `inject:""`
}

func (s *StrategyImportPackages) GetName() string {
	return "ImportPackages"
}

func (s *StrategyImportPackages) GetDescription() string {
	return "Check the project variables, functions, etc. naming spelling is wrong."
}

func (s *StrategyImportPackages) GetWeight() float64 {
	return 0.
}

// linterImportPackages is a function that scan the project contains all the
// package lists.It will extract from the linter need to convert
// the data.The result will be saved in the r's attributes.
func (s *StrategyImportPackages) Compute(parameters engine.StrategyParameter) (summaries *engine.Summaries) {
	summaries = engine.NewSummaries()

	importPkgs := unittest.GoListWithImportPackages(parameters.ProjectPath)
	for i := 0; i < len(importPkgs); i++ {
		summaries.Lock()
		summaries.Summaries[importPkgs[i]] = engine.Summary{Name: importPkgs[i]}
		summaries.Unlock()
	}
	return
}

func (s *StrategyImportPackages) Percentage(summaries *engine.Summaries) float64 {
	summaries.RLock()
	defer summaries.RUnlock()
	return utils.CountPercentage(len(summaries.Summaries))
}
