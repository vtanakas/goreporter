package strategy

import (
	"goreporter/engine"
	"strconv"
	"strings"

	"goreporter/linters/depth"
	"goreporter/utils"
)

type StrategyDepth struct {
	Sync            *engine.Synchronizer `inject:""`
	compBigThan3    int
	sumAverageDepth float64
	allDirs         map[string]string
}

func (s *StrategyDepth) GetName() string {
	return "Depth"
}

func (s *StrategyDepth) GetDescription() string {
	return "Computing all [.go] file's max depth."
}

func (s *StrategyDepth) GetWeight() float64 {
	return 0.05
}

// Compute all [.go] file's function maximum depth. It is an important indicator
// that allows developer to see whether a function needs to be splitted into smaller functions for readability purpose
func (s *StrategyDepth) Compute(parameters engine.StrategyParameter) (summaries *engine.Summaries) {
	summaries = engine.NewSummaries()

	s.allDirs = parameters.AllDirs

	sumProcessNumber := int64(10)
	processUnit := utils.GetProcessUnit(sumProcessNumber, len(s.allDirs))

	for pkgName, pkgPath := range s.allDirs {
		errors := make([]engine.Error, 0)
		depthResult, avg := depth.Depth(pkgPath)
		avgfloat, _ := strconv.ParseFloat(avg, 64)
		s.sumAverageDepth = s.sumAverageDepth + avgfloat
		for _, val := range depthResult {
			depthvalues := strings.Split(val, " ")
			if len(depthvalues) == 4 {
				comp, _ := strconv.Atoi(depthvalues[0])
				erroru := engine.Error{
					LineNumber:  comp,
					ErrorString: utils.AbsPath(depthvalues[3]),
				}
				if comp >= 3 {
					s.compBigThan3 = s.compBigThan3 + 1
				}
				errors = append(errors, erroru)
			}
		}
		summaries.Lock()
		summaries.Summaries[pkgName] = engine.Summary{
			Name:        pkgName,
			Errors:      errors,
			Description: avg,
		}
		summaries.Unlock()
		if sumProcessNumber > 0 {
			s.Sync.LintersProcessChans <- processUnit
			sumProcessNumber = sumProcessNumber - processUnit
		}
	}

	return
}

func (s *StrategyDepth) Percentage(summaries *engine.Summaries) float64 {
	return utils.CountPercentage(s.compBigThan3 + int(s.sumAverageDepth/float64(len(s.allDirs))) - 1)
}
