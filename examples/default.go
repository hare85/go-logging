package main

import logging "github.com/hare85/go-logging"

func main() {
	logger := logging.GetLogger("default")
	
	logger.Debug("DEBUG log - %s", "string")

	logger.Info("INFO log - %d", 1)

	test_map := map[string]int{}
	test_map["key1"] = 1
	test_map["key2"] = 2
	logger.Warn("WARN log - %x", test_map)

	logger.Error("ERROR log - %s")

	logger.Panic("PANIC log - %s", "This is panic log")
}
