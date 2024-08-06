package main

import (
	"flag"
	"fmt"

	"github.com/cihub/seelog"
)

func main() {
	configFilePath := flag.String("C", "conf/conf.yaml", "config file path")
	logConfigPath := flag.String("L", "conf/seelog.xml", "log Config file path")
	flag.Parse()

	logger, err := seelog.LoggerFromConfigAsFile(*logConfigPath)
	if err != nil {
		seelog.Critical("--err parsing seelog config file", err)
		return
	}

	seelog.ReplaceLogger(logger)
	if err != nil {
		seelog.Critical("--ReplaceLogger", err)
		return
	}
	defer seelog.Flush()

	fmt.Println(configFilePath)

}
