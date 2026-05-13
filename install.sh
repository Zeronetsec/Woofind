#!/usr/bin/env bash
# https://github.com/Zeronetsec/Woofind

N='\033[0m'
R='\033[1;31m'
B='\033[1;34m'
GG='\033[0;32m'
DG='\033[1;90m'

base="${PREFIX}/opt"
symlink="${PREFIX}/bin"
bkdate="$(command date +%Y_%b_%d_%H_%M_%S)"

path="$(
    cd -- "$(
        command dirname -- "${BASH_SOURCE[0]}"
    )" &> /dev/null && pwd
)"

if [[ "${1}" == "--backup" ]]; then
    backup="true"
fi

function install() {
    local cmd="${1}"
    local desc="${2}"
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

    echo -e "${1}" | while IFS= read -r line; do
        [[ -z "${line}" ]] && continue
        IFS="::" read -ra pkgs <<< "${line}"
        for pkg in "${pkgs[@]}"; do
            pkg="$(echo -e "${pkg}" | command xargs)"
            if eval "${installw} ${pkg}" 2>/dev/null; then
                break
            fi
        done
    done
}

if [[ ! -d "${path}" ]]; then
    echo -e "\n${R}[!] ${N}Folder: ${GG}${path} ${N}not found! \n"
    exit 1
fi

echo -e "\n${B}[*] ${N}Installing: ${GG}Woofind${N}"

pack=(
    "golang"
    "zip"
)

for i in "${pack[@]}"; do
    install \
        "getinstall ${i} -y" \
        "Installing: ${GG}${i}${N}"
done

if [[ ! -d "${base}" ]]; then
    install \
        "command mkdir -p ${base}" \
        "Create directory: ${GG}${base}${N}"
fi


if [[ "${backup}" == "true" && -d "${base}/woofind" ]]; then
    cd "${base}"
    install \
        "command zip -r woofind_${bkdate}.bak.zip woofind" \
        "Backup: ${GG}${base}/woofind ${DG}=> ${GG}${base}/woofind_${bkdate}.bak.zip${N}"
    cd
fi

if [[ -d "${base}/woofind" ]]; then
    install \
        "command rm -rf ${base}/woofind" \
        "Removing: ${GG}old woofind${N}"
fi

install \
    "command mv ${path} ${base}/woofind" \
    "Moving: ${GG}${path} ${DG}=> ${GG}${base}/woofind${N}"

cd "${base}/woofind"
install \
    "command go mod tidy" \
    "Retidy: ${GG}woofind${N}"

install \
    "command go build -v -o woofind" \
    "Building: ${GG}woofind${N}"
cd

install \
    "command ln -sf ${base}/woofind/woofind ${symlink}/woofind" \
    "Symlink: ${GG}${base}/woofind/woofind ${DG}=> ${GG}${symlink}/woofind${N}"

printf '\n'
if command -v woofind &>/dev/null; then
    echo -e "${GG}[+] ${N}Woofind installed!"
    echo -e "${GG}[+] ${N}Usage: ${GG}woofind --help ${N}to show helper"
    printf '\n'
    exit 0
else
    echo -e "${R}[!] ${N}Failed installing woofind!"
    printf '\n'
    exit 1
fi

# Copyright (c) 2026 Zeronetsec