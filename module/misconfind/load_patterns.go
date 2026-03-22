// Comet Framework

package misconfind

import (
    "bufio"
    "io"
    "strings"
    "io/fs"
)

func loadPatterns(fsys fs.FS, path string) ([]string, error) {
    var reader io.Reader

    if f, err := fsys.Open(path); err == nil {
        reader = f
    }

    var patterns []string
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        patterns = append(patterns, strings.ToLower(line))
    }

    return patterns, scanner.Err()
}

// Copyright (c) 2026 Zeronetsec