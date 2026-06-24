<!-- https://github.com/Zeronetsec/Woofind -->

# Installation
`install.sh` optional option:
- `--backup`

Use `--backup` to create a backup of the existing Woofind installation before replacing it.

## Termux & Linux (root)
```bash
git clone https://github.com/Zeronetsec/Woofind
cd Woofind
chmod +x install.sh
./install.sh
```

## Linux (user)
```bash
git clone https://github.com/Zeronetsec/Woofind
cd Woofind
chmod +x install.sh
sudo ./install.sh
```

## Uninstallation
```bash
export prefix="${PREFIX:-/usr}"
rm -f "${prefix}/bin/woofind"
rm -rf "${prefix}/opt/woofind"
```

<!-- Copyright (c) 2026 Zeronetsec -->