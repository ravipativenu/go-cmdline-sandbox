package cf

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const isWin = runtime.GOOS == "windows"
const CF_COMMAND = "cf"

func CfRun(args ...string) (string, error) {
	// Replace double quotes with escaped double quotes in each argument
	for i, arg := range args {
		args[i] = strings.ReplaceAll(arg, `"`, `\"`)
	}

	// Construct the command line
	//cmdLine := fmt.Sprintf("%s %s", CF_COMMAND, strings.Join(args, " "))

	// Execute the command
	cmd := exec.Command("cf")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	fmt.Println(string(output))
	return strings.TrimSpace(string(output)), nil
}
