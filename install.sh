#!/usr/bin/env bash
# Comet Framework

N='\033[0m'
R='\033[1;31m'
B='\033[1;34m'
GG='\033[0;32m'
DG='\033[1;90m'

base="$PREFIX/opt"
symlink="$PREFIX/bin"
bkdate="$(command date +%Y_%b_%d_%H_%M_%S)"

function spinner() {
    local cmd=""
    local msg="Running command"

    for args in "$@"; do
        case "$args" in
            --command=*)
                cmd="${args#*=}"
                ;;
            --message=*)
                msg="${args#*=}"
                ;;
        esac
    done

    function spinner_anim() {
        local pid=$1
        local delay=0.1
        local spinstr='|/-\'

        while kill -0 "$pid" 2>/dev/null; do
            local temp="${spinstr#?}"
            printf " \033[0m[\033[0;32m%c\033[0m] " "$spinstr"
            spinstr="$temp${spinstr%"$temp"}"
            command sleep "$delay"
            printf "\b\b\b\b\b"
        done
        printf "     \b\b\b\b\b"
    }

    echo -ne "${B}[*] ${N}${msg}${N}"
    command bash -c "${cmd}" > /dev/null 2>&1 &
    local pid=$!

    spinner_anim "$pid"
    wait "$pid"
    local status=$?

    if [[ "$status" -eq 0 ]]; then
        echo -ne " ${DG}- ${N}[${GG}✔ ${N}Success]\n"
    else
        echo -ne " ${DG}- ${N}[${R}✖ ${N}Failed (exit code: ${GG}${status}${N})]\n"
        echo -ne "\033[?25h"
        exit $status
    fi
}

printf '\n'
echo -e "${B}[*] ${N}Enter path to your comet folder"
read -p "$(echo -e "${N}Path: ")" path
declare -A varmap=(
    ["~"]="$HOME"
    ["\$HOME"]="$HOME"
    ["\$PREFIX"]="$PREFIX"
    ["\$PWD"]="$PWD"
)

for exp in "${!varmap[@]}"; do
    path="${path/#$exp/${varmap[$exp]}}"
done

if [[ ! -d "$path" ]]; then
    echo -e "\n${R}[!] ${N}Folder: ${GG}${path} ${N}not found! \n"
    exit 1
fi

echo -ne "\033[?25l"
echo -e "\n${B}[*] ${N}Installing: ${GG}Comet${N}"

pack=(
    "golang"
    "zip"
)

for i in "${pack[@]}"; do
    spinner \
        --command="command apt install \"${i}\" -y" \
        --message="Installing: ${GG}${i}${N}"
done

if [[ ! -d "$base" ]]; then
    spinner \
        --command="command mkdir -p \"${base}\"" \
        --message="Create directory: ${GG}${base}${N}"
fi


if [[ -d "$base/comet" ]]; then
    echo -ne "\033[?25h\n"
    read -p "$(echo -e "${B}[*] ${N}Do you wan't to backup ${GG}${base}/comet${N}? (y/n) ")" chs
    echo -ne "\033[?25l\n"

    if [[ "$chs" == 'y' ]]; then
        cd "$base"
        spinner \
            --command="command zip -r \"comet_${bkdate}.bak.zip\" \"comet\"" \
            --message="Backup: ${GG}${base}/comet ${DG}=> ${GG}${base}/comet_${bkdate}.bak.zip${N}"
        cd
    fi

    spinner \
        --command="command rm -rf \"${base}/comet\"" \
        --message="Removing: ${GG}old comet${N}"
fi

spinner \
    --command="command mv \"${path}\" \"${base}/comet\"" \
    --message="Moving: ${GG}${path} ${DG}=> ${GG}${base}/comet${N}"

spinner \
    --command="command chmod +x -R \"${base}/comet/\"" \
    --message="Setting up permission"

lmeta="$base/comet/utils/listcmd/metadata"
smeta="$base/comet/utils/searchcmd/metadata"
spinner \
  --command="command cp -f ${lmeta}/*.json ${smeta}/ && command rm -f ${smeta}/placeholder.txt" \
  --message="Sync: ${GG}${lmeta}/*.json ${GG}=> ${GG}${smeta}/"

cd "$base/comet"
spinner \
    --command="command go build -o comet" \
    --message="Building comet"
cd

spinner \
    --command="command ln -sf \"${base}/comet/comet\" \"${symlink}/comet\"" \
    --message="Symlink: ${GG}${base}/comet/comet ${DG}=> ${GG}${symlink}/comet${N}"

printf '\n'
if command -v comet &>/dev/null; then
    echo -e "${GG}[+] ${N}comet installed!"
    echo -e "${GG}[+] ${N}Usage: ${GG}comet --help ${N}to show helper"
    echo -ne "\033[?25h\n"
    exit 0
else
    echo -e "${R}[!] ${N}Failed installing Comet!"
    echo -ne "\033[?25l\n"
    exit 1
fi

# Copyright (c) 2026 Zeronetsec