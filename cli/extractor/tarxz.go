package extractor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func extractTarXzExternal(path string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("file does not exist: %w", err)
	}

	if !strings.HasSuffix(strings.ToLower(path), ".tar.xz") {
		return fmt.Errorf("file is not a .tar.xz archive")
	}

	tarCmd := "tar"
	if runtime.GOOS == "windows" {
		systemTar := filepath.Join(os.Getenv("SystemRoot"), "System32", "tar.exe")
		if _, err := os.Stat(systemTar); err == nil {
			tarCmd = systemTar
		} else {
			if pathTar, err := exec.LookPath("tar"); err == nil {
				tarCmd = pathTar
			} else {
				return fmt.Errorf("tar not found on Windows system. Please install tar (e.g., via Git Bash, WSL, or Windows 10+)")
			}
		}
	}

	if _, err := exec.LookPath(tarCmd); err != nil {
		return fmt.Errorf("tar not found on system: %w", err)
	}

	destDir := filepath.Dir(path)

	cmd := exec.Command(tarCmd, "-xJf", path, "-C", destDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tar.xz using system tar: %w", err)
	}

	return nil
}
