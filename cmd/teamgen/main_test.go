package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

var binName = "teamgen"

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build %s: %s\n", binName, err)
		os.Exit(1)
	}
	fmt.Println("Running tests...")
	result := m.Run()
	fmt.Println("Cleaning up")
	os.Exit(result)
}

func TestTeamgenCLI(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	t.Run("BaseRun", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		fmt.Printf("First line: %s\n", output[0])
		matched_0, err := regexp.MatchString(`[a-zA-Z]+\s+[a-zA-Z]+\s+\[[FM]\]\s+[2-9A-F]{6}\s+Age:\s+[1-5][0-9]\s+human`, output[0])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_0 {
			t.Errorf("Did not find match")
		}

		fmt.Printf("Second line: %s\n", output[1])
		fmt.Printf("Third line: %s\n", output[2])
		//matched_1, err := regexp.MatchString(`[a-zA-Z]+\s+[a-zA-Z]+\s+\[[FM]\]\s+[2-9A-F]{6}\s+Age:\s+[1-5][0-9]\s+human`, output[0])
		//if err != nil { t.Fatal(err) }
		//if !matched_0 { t.Errorf("Did not find match") }
	})

  t.Run("TestGenderF", func(t *testing.T) {
    cmd := exec.Command(cmdPath, "-gender", "F")
    out, err := cmd.CombinedOutput()
    if err != nil {
      t.Fatal(err)
    }
    output := strings.Split(string(out), "\n")
		fmt.Printf("First line: %s\n", output[0])
    matched_0, err := regexp.MatchString(`\[F\]`, output[0])
    if err != nil {
      t.Fatal(err)
    }
    if !matched_0 {
      t.Error("Did not find match")
    } 
  })


}
