package main

import (
	"fmt"
	"log"
	"os"

	"github.com/seq-gen/pkg/analyzer"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <file.go>", os.Args[0])
	}

	filename := os.Args[1]
	analyzer := &analyzer.Analyzer{}

	if err := analyzer.AnalyzeFile(filename); err != nil {
		log.Fatalf("Error analyzing file: %v", err)
	}

	calls := analyzer.GetCalls()
	fmt.Println("@startuml")
	for _, call := range calls {
		fmt.Printf("%s -> %s: %s\n", "Client", call.Receiver, call.Method)
	}
	fmt.Println("@enduml")
}
