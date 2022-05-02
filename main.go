package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"sigs.k8s.io/yaml"
)

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to construct logger: %v", err)
	}
	defer l.Sync()

	scanner := bufio.NewScanner(os.Stdin)
	var buffer []byte
	for scanner.Scan() {
		buffer = append(buffer, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		l.Fatal("failed to read from stdin", zap.Error(err))
	}

	converted, err := yaml.JSONToYAML(buffer)
	if err != nil {
		l.Fatal("failed to convert input from json to yaml", zap.Error(err))
	}
	fmt.Println(string(converted))
}
