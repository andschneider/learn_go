package main

import "testing"
import "os/exec"

// exercise 1.3
// Experiment to measure the difference in running time between our potentially 
// inefficient versions and the one that uses strings.Join.


func BenchmarkBase(b *testing.B) {
	args := [3]string{"a", "b", "c"}
	for i := 1; 1 < b.N; i+=100 {
		exec.Command(echoJoin(), args)
	}
}
