package main

import (
	"os"

	_ "github.com/AiRISTAFlowInc/flow-studio-core/data/expression/script"
	"github.com/AiRISTAFlowInc/flow-studio-core/engine"
	"github.com/AiRISTAFlowInc/flow-studio-core/support/log"
)

var (
	cfgJson       string
	cfgEngine     string
	cfgCompressed bool
	flogoEngine   engine.Engine
)

func init() {
	log.SetLogLevel(log.RootLogger(), log.ErrorLevel)

	cfg, err := engine.LoadAppConfig(cfgJson, cfgCompressed)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}

	flogoEngine, err = engine.New(cfg, engine.ConfigOption(cfgEngine, cfgCompressed), engine.DirectRunner)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}
}
