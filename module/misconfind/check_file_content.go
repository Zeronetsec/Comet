// Comet Framework

package misconfind

import (
    "os"
    "bufio"
    "strings"
    "fmt"
    "comet/utils/color"
)

func checkFileContent(path string, patterns []string) {
    file, err := os.Open(path)
    if err != nil {
        return
    }
    defer file.Close()

    if isBinary(file) {
        return
    }

    file.Seek(0, 0)

    scanner := bufio.NewScanner(file)
    lineNum := 0

    for scanner.Scan() {
        lineNum++
        line := strings.ToLower(scanner.Text())

        for _, p := range patterns {
            if strings.Contains(line, p) {
                fmt.Printf(
                    "%s[Alert] %sPattern: %s%s %sin %s%s%s:%s%d%s\n",
                    color.R, color.N, color.GG, p, color.N, color.GG, path, color.DG, color.GG, lineNum, color.N,
                )
            }
        }
    }
}

// Copyright (c) 2026 Zeronetsec