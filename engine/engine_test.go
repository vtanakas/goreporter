// Copyright 2017 The GoReporter Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"goreporter/lint/strategy"
	"log"
	"sync"
	"testing"

	"github.com/facebookgo/inject"
	"github.com/golang/glog"
	"goreporter/engine/processbar"
)

func Test_Engine(t *testing.T) {
	synchronizer := &Synchronizer{
		LintersProcessChans:   make(chan int64, 20),
		LintersFinishedSignal: make(chan string, 10),
	}
	syncRW := &sync.RWMutex{}
	waitGW := &WaitGroupWrapper{}

	reporter := NewReporter("../../../wgliang/logcool", "foo", "foo", "baz")
	strategyCopyCheck := &strategy.StrategyCopyCheck{}
	strategyCountCode := &strategy.StrategyCountCode{}
	strategyCyclo := &strategy.StrategyCyclo{}
	strategyDeadCode := &strategy.StrategyDeadCode{}
	strategyDependGraph := &strategy.StrategyDependGraph{}
	strategyDepth := &strategy.StrategyDepth{}
	strategyImportPackages := &strategy.StrategyImportPackages{}
	strategyInterfacer := &strategy.StrategyInterfacer{}
	strategySimpleCode := &strategy.StrategySimpleCode{}
	strategySpellCheck := &strategy.StrategySpellCheck{}
	strategyUnitTest := &strategy.StrategyUnitTest{}

	if err := inject.Populate(
		reporter,
		synchronizer,
		strategyCopyCheck,
		strategyCountCode,
		strategyCyclo,
		strategyDeadCode,
		strategyDependGraph,
		strategyDepth,
		strategyImportPackages,
		strategyInterfacer,
		strategySimpleCode,
		strategySpellCheck,
		strategyUnitTest,
		syncRW,
		waitGW,
	); err != nil {
		log.Fatal(err)
	}

	reporter.AddLinters(strategyCopyCheck, strategyCountCode, strategyCyclo, strategyDeadCode, strategyDependGraph,
		strategyDepth, strategyImportPackages, strategyInterfacer, strategySimpleCode, strategySpellCheck, strategyUnitTest)

	go processbar.LinterProcessBar(synchronizer.LintersProcessChans, synchronizer.LintersFinishedSignal)

	if err := reporter.Report(); err != nil {
		glog.Errorln(err)
	}
}
