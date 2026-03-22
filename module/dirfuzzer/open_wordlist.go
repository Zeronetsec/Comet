// Comet Framework

package dirfuzzer

import (
    "io"
    "os"
    "bufio"
    "strings"
    "io/fs"
)

func openWordlist(fsys fs.FS, path string) (*bufio.Scanner, func(), int, error) {
    var reader io.Reader
    var closer func() = func() {}

    if fsys != nil {
        if f, err := fsys.Open(path); err == nil {
            reader = f
        }
    }

    if reader == nil {
        file, err := os.Open(path)
        if err != nil {
            return nil, nil, 0, err
        }

        reader = file
        closer = func() {
            file.Close()
        }
    }

    scanner := bufio.NewScanner(reader)
    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    total := len(lines)
    newReader := strings.NewReader(strings.Join(lines, "\n"))
    return bufio.NewScanner(newReader), closer, total, nil
}

// Copyright (c) 2026 Zeronetsec