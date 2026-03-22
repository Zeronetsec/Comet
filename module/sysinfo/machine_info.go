// Comet Framework

package sysinfo

import (
    "fmt"
    "os"
    "runtime"
    "os/user"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/host"
    "comet/utils/color"
)

func MachineInfo() {
    usr, err := user.Current()
    username := os.Getenv("USER")
    virmem, _ := mem.VirtualMemory()
    cpuinfo, _ := cpu.Info()
    up, _ := host.Uptime()
    hostname, _ := os.Hostname()
    ip := execShell("hostname", "-i")
    machine := execShell("uname", "--machine")
    krn := execShell("uname", "--kernel-name")
    krs := execShell("uname", "--kernel-release")
    krv := execShell("uname", "--kernel-version")
    whoami := execShell("whoami")

    if username == "" && err == nil {
        username = usr.Username
    }

    fmt.Printf(
        "%s[*] %sSystem information\n",
        color.B, color.N,
    )

    fmt.Println()

    if err != nil {
        fmt.Printf(
            "%sUser: %sunknown%s\n",
            color.N, color.GG, color.N,
        )
    } else {
        fmt.Printf(
            "%sUser: %s%s (%s)%s\n",
            color.N, color.GG, username, whoami, color.N,
        )

        fmt.Printf(
            "%sUID: %s%s%s\n",
            color.N, color.GG, usr.Uid, color.N,
        )

        fmt.Printf(
            "%sGID: %s%s%s\n",
            color.N, color.GG, usr.Gid, color.N,
        )
    }

    fmt.Printf(
        "%sHostname: %s%s%s\n",
        color.N, color.GG, hostname, color.N,
    )

    fmt.Printf(
        "%sOS: %s%s%s\n",
        color.N, color.GG, runtime.GOOS, color.N,
    )

    fmt.Printf(
        "%sArch: %s%s%s\n",
        color.N, color.GG, runtime.GOARCH, color.N,
    )

    fmt.Printf(
        "%sGo Ver: %s%s%s\n",
        color.N, color.GG, runtime.Version(), color.N,
    )

    fmt.Printf(
        "%sShell: %s%s%s\n",
        color.N, color.GG, os.Getenv("SHELL"), color.N,
    )

    fmt.Printf(
        "%sTerm: %s%s%s\n",
        color.N, color.GG, os.Getenv("TERM"), color.N,
    )

    fmt.Printf(
        "%sHome: %s%s%s\n",
        color.N, color.GG, os.Getenv("HOME"), color.N,
    )

    fmt.Printf(
        "%sMemory: %s%d%s%d%s%s\n",
        color.N, color.GG, virmem.Used/1024/1024, "MB /", virmem.Total/1024/1024, "MB", color.N,
    )

    fmt.Printf(
        "%sCPU: %s%s%s\n",
        color.N, color.GG, cpuinfo[0].ModelName, color.N,
    )

    fmt.Printf(
        "%sUptime: %s%d%s%s\n",
        color.N, color.GG, up, "seconds", color.N,
    )

    fmt.Printf(
        "%sLocal IP: %s%s%s\n",
        color.N, color.GG, ip, color.N,
    )

    fmt.Printf(
        "%sMachine: %s%s%s\n",
        color.N, color.GG, machine, color.N,
    )

    fmt.Printf(
        "%sKernel Name: %s%s%s\n",
        color.N, color.GG, krn, color.N,
    )

    fmt.Printf(
        "%sKernel Rel: %s%s%s\n",
        color.N, color.GG, krs, color.N,
    )

    fmt.Printf(
        "%sKernel Ver: %s%s%s\n",
        color.N, color.GG, krv, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec