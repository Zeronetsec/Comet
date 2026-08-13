// https://github.com/Zeronetsec/Comet

package sqlscan

import (
    "bufio"
    "os"
    "strings"
    "path/filepath"
    "io/fs"
)

func LoadWordlist(
    path string,
    embeddedFS fs.FS,
) ([]string, error) {
    unique := make(map[string]bool)
    var payloads []string

    processScanner := func(scanner *bufio.Scanner) {
        for scanner.Scan() {
            line := strings.TrimSpace(scanner.Text())
            if line == "" || strings.HasPrefix(line, "#") {
                continue
            }

            if !unique[line] {
                unique[line] = true
                payloads = append(payloads, line)
            }
        }
    }

    if path == "embedded" {
        err := fs.WalkDir(
            embeddedFS,
            "wordlist/sqlscan",
            func(
                p string, d fs.DirEntry, err error,
            ) error {
                if err != nil {
                    return err
                }
                if !d.IsDir() && strings.HasSuffix(
                    strings.ToLower(d.Name()),
                    ".txt",
                ) {
                    file, err := embeddedFS.Open(p)
                    if err != nil {
                        return err
                    }
                    defer file.Close()
                    scanner := bufio.NewScanner(file)
                    processScanner(scanner)
                }
                return nil
            },
        )
        return payloads, err
    }

    info, err := os.Stat(path)
    if err != nil {
        return nil, err
    }

    if info.IsDir() {
        err = filepath.Walk(
            path,
            func(
                p string, i os.FileInfo, e error,
            ) error {
                if e != nil {
                    return e
                }

                if !i.IsDir() && strings.HasSuffix(
                    strings.ToLower(i.Name()),
                    ".txt",
                ) {
                    file, err := os.Open(p)
                    if err != nil {
                        return err
                    }
                    defer file.Close()
                    scanner := bufio.NewScanner(file)
                    processScanner(scanner)
                }
                return nil
            },
        )
    } else {
        file, err := os.Open(path)
        if err != nil {
            return nil, err
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)
        processScanner(scanner)
    }

    return payloads, err
}

// Copyright (c) 2026 Zeronetsec