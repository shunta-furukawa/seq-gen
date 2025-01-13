package test

import (
	"testing"

	"github.com/shunta-furukawa/seq-gen/pkg/analyzer"
)

func TestGRPCServerAnalysis(t *testing.T) {
	// 解析対象のgRPCサーバファイルパス
	filename := "./grpcserver/server.go"

	analyzer := &analyzer.Analyzer{}
	if err := analyzer.AnalyzeFile(filename); err != nil {
		t.Fatalf("Failed to analyze file: %v", err)
	}

	calls := analyzer.GetCalls()
	if len(calls) == 0 {
		t.Fatalf("Expected calls to be detected, but got none")
	}

	for _, call := range calls {
		t.Logf("Detected call: Receiver=%s, Method=%s", call.Receiver, call.Method)
	}
}
