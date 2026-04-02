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

path="$(
    cd -- "$(
        command dirname -- "${BASH_SOURCE[0]}"
    )" &> /dev/null && pwd
)"

if [[ "$1" == "--backup" ]]; then
    backup="true"
fi

function install() {
    local cmd="$1"
    local desc="$2"
    echo -e "\n${B}[*] ${N}${desc}"
    eval "${cmd}" >/dev/null
    local status=$?
    echo -e "    ${DG}└── ${N}exit: ${GG}${status}${N}"
}

function getinstall() {
    if command -v apt >/dev/null 2>&1; then
        installw="command apt install -y"
    elif command -v apk >/dev/null 2>&1; then
        installw="command apk add"
    elif command -v pacman >/dev/null 2>&1; then
        installw="command pacman -S --noconfirm"
    else
        exit 1
    fi

    echo -e "$1" | while IFS= read -r line; do
        [[ -z "$line" ]] && continue
        IFS="::" read -ra pkgs <<< "$line"
        for pkg in "${pkgs[@]}"; do
            pkg="$(echo -e "$pkg" | command xargs)"
            if eval "$installw $pkg" 2>/dev/null; then
                break
            fi
        done
    done
}

if [[ ! -d "$path" ]]; then
    echo -e "\n${R}[!] ${N}Folder: ${GG}${path} ${N}not found! \n"
    exit 1
fi

echo -e "\n${B}[*] ${N}Installing: ${GG}Comet${N}"

pack=(
    "golang"
    "zip"
)

for i in "${pack[@]}"; do
    install \
        "getinstall ${i} -y" \
        "Installing: ${GG}${i}${N}"
done

if [[ ! -d "$base" ]]; then
    install \
        "command mkdir -p ${base}" \
        "Create directory: ${GG}${base}${N}"
fi


if [[ "$backup" == "true" && -d "$base/comet" ]]; then
    cd "$base"
    install \
        "command zip -r comet_${bkdate}.bak.zip comet" \
        "Backup: ${GG}${base}/comet ${DG}=> ${GG}${base}/comet_${bkdate}.bak.zip${N}"
    cd
fi

if [[ -d "$base/comet" ]]; then
    install \
        "command rm -rf ${base}/comet" \
        "Removing: ${GG}old comet${N}"
fi

install \
    "command mv ${path} ${base}/comet" \
    "Moving: ${GG}${path} ${DG}=> ${GG}${base}/comet${N}"

lmeta="$base/comet/utils/listcmd/metadata"
smeta="$base/comet/utils/searchcmd/metadata"

install \
    "command cp -f ${lmeta}/*.json ${smeta}/ && command rm -f ${smeta}/placeholder.txt" \
    "Sync: ${GG}${lmeta}/*.json ${DG}=> ${GG}${smeta}/"

cd "$base/comet"
install \
    "command go mod tidy" \
    "Retidy: ${GG}comet${N}"

install \
    "command go build -v -o comet" \
    "Building: ${GG}comet${N}"
cd

install \
    "command ln -sf ${base}/comet/comet ${symlink}/comet" \
    "Symlink: ${GG}${base}/comet/comet ${DG}=> ${GG}${symlink}/comet${N}"

printf '\n'
if command -v comet &>/dev/null; then
    echo -e "${GG}[+] ${N}Comet installed!"
    echo -e "${GG}[+] ${N}Usage: ${GG}comet --help ${N}to show helper"
    printf '\n'
    exit 0
else
    echo -e "${R}[!] ${N}Failed installing comet!"
    printf '\n'
    exit 1
fi

# Copyright (c) 2026 Zeronetsec