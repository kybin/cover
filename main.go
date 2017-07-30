package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	tmpf := filepath.Join(os.TempDir(), "coverage.out")

	cmd := exec.Command("go", "test", "-coverprofile", tmpf)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	out := strings.TrimSpace(string(bs))
	fmt.Println(out)

	if strings.HasSuffix(out, "[no test files]") {
		return
	}

	cmd = exec.Command("go", "tool", "cover", "-html", tmpf)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}
