// https://github.com/Zeronetsec/Comet

package logger

func NewLogger(fileName string) *Logger {
    return &Logger{FileName: fileName}
}

// Copyright (c) 2026 Zeronetsec