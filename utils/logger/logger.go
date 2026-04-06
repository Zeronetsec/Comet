// Comet Framework

package logger

import (
    "fmt"
    "os"
    "time"
    "path/filepath"
)

type Logger struct {
    FileName string
}

func NewLogger(fileName string) *Logger {
    return &Logger{FileName: fileName}
}

func (l *Logger) Log(tag string, message string) error {
    home, err := os.UserHomeDir()
    if err != nil {
        return fmt.Errorf(
            "Failed get $HOME: %w",
            err,
        )
    }

    basePath := filepath.Join(home, ".comet_log")
    now := time.Now()
    dateStr := now.Format("2006_January_02")
    timeStr := now.Format("15:04")

    var prefix string
    switch tag {
        case ":error": prefix = "[!]"
        case ":": prefix = "[+]"
        case ":info": prefix = "[*]"
        case ":?": prefix = "[?]"
        case ":-": prefix = "[-]"
        default: prefix = "[ ]"
    }

    logFolder := filepath.Join(basePath, l.FileName)
    if err := os.MkdirAll(logFolder, 0755); err != nil {
        return err
    }

    fileName := fmt.Sprintf(
        "%s_%s.log", l.FileName, dateStr,
    )

    fullPath := filepath.Join(logFolder, fileName)
    content := fmt.Sprintf(
        "%s - %s %s\n", timeStr, prefix, message,
    )

    f, err := os.OpenFile(
        fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
    )

    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.WriteString(content)
    return err
}

// Copyright (c) 2026 Zeronetsec