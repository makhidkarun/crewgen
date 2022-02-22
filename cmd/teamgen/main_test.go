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
		matched_0, err := regexp.MatchString(`[a-zA-Z]+\s+[a-zA-Z]+\s+\[[FM]\]\s+[2-9A-F]{6}\s+Age:\s+[1-5][0-9]\s+human`, output[0])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_0 {
			t.Errorf("Did not find match")
		}
	})

	t.Run("TestGenderF", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-gender", "F")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_0, err := regexp.MatchString(`\[F\]`, output[0])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_0 {
			t.Error("Did not find match to [F]")
		}
	})

	t.Run("TestGenderM", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-gender", "M")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_0, err := regexp.MatchString(`\[M\]`, output[0])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_0 {
			t.Error("Did not find match to [M]")
		}
	})

	t.Run("TestGenderOdd", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-gender", "G")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_0, err := regexp.MatchString(`\[[FM]\]`, output[0])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_0 {
			t.Error("Did not find match to [M] or [F]")
		}
	})

	t.Run("TestTerms1", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-terms", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_1, err := regexp.MatchString(`^1 term`, output[1])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_1 {
			t.Error("Did not find match to 1 term")
		}
	})

	t.Run("TestTerms10", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-terms", "10")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_1, err := regexp.MatchString(`^[1-5] term`, output[1])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_1 {
			t.Error("Did not find match to random term less than 10")
		}
	})

	t.Run("TestAgeOneTerm", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-terms", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_1, err := regexp.MatchString(`Age: 2[2-5]`, output[0])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_1 {
			t.Error("Age not correct for 1 term")
		}
	})

	t.Run("TestCareerNavy", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-career", "Navy")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_1, err := regexp.MatchString(`Navy`, output[1])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_1 {
			t.Error("Career not Navy")
		}
	})

	t.Run("TestCareerMerchant", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-career", "Merchant")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_1, err := regexp.MatchString(`Merchant`, output[1])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_1 {
			t.Error("Career not Merchant")
		}
	})

	t.Run("TestCareerMarines", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-career", "Marines")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_1, err := regexp.MatchString(`Marines`, output[1])
		if err != nil {
			t.Fatal(err)
		}
		if !matched_1 {
			t.Error("Career not Marines")
		}
	})

	t.Run("TestJobPilot", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-job", "pilot", "-terms", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_2, err := regexp.MatchString(`Pilot-[1-9]`, output[2])
		if !matched_2 {
			t.Error("Does not have Pilot skill")
		}
	})

	t.Run("TestJobMedic", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-job", "medic", "-terms", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_2, err := regexp.MatchString(`Medical-[1-9]`, output[2])
		if !matched_2 {
			t.Error("Does not have Medic skill")
		}
	})

	t.Run("TestJobEngineer", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-job", "engineer", "-terms", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_2, err := regexp.MatchString(`Engineering-[1-9]`, output[2])
		if !matched_2 {
			t.Error("Does not have Engineering skill")
		}
	})

	t.Run("TestJobNavigator", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-job", "navigator", "-terms", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_2, err := regexp.MatchString(`Navigation-[1-9]`, output[2])
		if !matched_2 {
			t.Error("Does not have Navigation skill")
		}
	})

	t.Run("TestJobGunner", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-job", "gunner", "-terms", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_2, err := regexp.MatchString(`Gunnery-[1-9]`, output[2])
		if !matched_2 {
			t.Error("Does not have Gunnery skill")
		}
	})

	t.Run("TestJobSteward", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-job", "steward", "-terms", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(string(out), "\n")
		matched_2, err := regexp.MatchString(`Steward-[1-9]`, output[2])
		if !matched_2 {
			t.Error("Does not have Steward skill")
		}
	})

}
