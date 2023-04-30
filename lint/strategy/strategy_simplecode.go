package strategy

import (
	"goreporter/engine"
	"strconv"
	"strings"

	"goreporter/lint/linters/simplecode"
	"goreporter/utils"
)

type StrategySimpleCode struct {
	Sync *engine.Synchronizer `inject:""`
}

func (s *StrategySimpleCode) GetName() string {
	return "Simple"
}

func (s *StrategySimpleCode) GetDescription() string {
	return "All golang code hints that can be optimized and give suggestions for changes."
}

func (s *StrategySimpleCode) GetWeight() float64 {
	return 0.05
}

func (s *StrategySimpleCode) Compute(parameters engine.StrategyParameter) (summaries *engine.Summaries) {
	summaries = engine.NewSummaries()

	simples := simplecode.Simple(parameters.AllDirs, parameters.ExceptPackages)
	sumProcessNumber := int64(10)
	processUnit := utils.GetProcessUnit(sumProcessNumber, len(simples))
	for _, simpleTip := range simples {
		simpleTips := strings.Split(simpleTip, ":")
		if len(simpleTips) == 4 {
			packageName := utils.PackageNameFromGoPath(simpleTips[0])
			line, _ := strconv.Atoi(simpleTips[1])
			erroru := engine.Error{
				LineNumber:  line,
				ErrorString: utils.AbsPath(simpleTips[0]) + ":" + strings.Join(simpleTips[1:], ":"),
			}
			summaries.Lock()
			if summarie, ok := summaries.Summaries[packageName]; ok {
				summarie.Errors = append(summarie.Errors, erroru)
				summaries.Summaries[packageName] = summarie
			} else {
				summarie := engine.Summary{
					Name:   packageName,
					Errors: make([]engine.Error, 0),
				}
				summarie.Errors = append(summarie.Errors, erroru)
				summaries.Summaries[packageName] = summarie
			}
			summaries.Unlock()
		}
		if sumProcessNumber > 0 {
			s.Sync.LintersProcessChans <- processUnit
			sumProcessNumber = sumProcessNumber - processUnit
		}
	}

	return summaries
}

func (s *StrategySimpleCode) Percentage(summaries *engine.Summaries) float64 {
	summaries.RLock()
	defer summaries.RUnlock()
	return utils.CountPercentage(len(summaries.Summaries))
}
