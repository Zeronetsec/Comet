// https://github.com/Zeronetsec/Comet

package corsscan

import (
    "os"
    "bufio"
    "strings"
)

func readTarget(input string) ([]string, error) {
    var targets []string
    info, err := os.Stat(input)

    if err == nil && !info.IsDir() {
        file, err := os.Open(input)
        if err != nil {
            return nil, err
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            t := strings.TrimSpace(scanner.Text())
            if t != "" {
                if !strings.HasPrefix(t, "http") {
                    t = "http://" + t
                }
                targets = append(targets, t)
            }
        }
    } else {
        t := input
        if !strings.HasPrefix(t, "http") {
            t = "http://" + t
        }
        targets = append(targets, t)
    }

    return targets, nil
}

// Copyright (c) 2026 Zeronetsec