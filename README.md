<!-- https://github.com/Zeronetsec/Woofind -->

[![version](https://img.shields.io/badge/Woofind-Version%201.0-blue.svg?maxAge=259200)]()
[![gover](https://img.shields.io/badge/Go-Version%201.26.1-blue.svg)]()
[![os](https://img.shields.io/badge/Supported%20OS-Linux-blue.svg)]()
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

# Woofind
Woofind is an all-in-one CLI toolkit for system exploration, security auditing, and data inspection. <br>
From decoding obscure strings to uncovering system misconfigurations, Woofind gives you deeper insight into your system.

## Features
- Inspect system and hardware information
- Analyze running processes and resource usage
- Detect misconfigurations and potential risks
- Decode and analyze encoded data automatically
- Extract insights from strings (entropy, hashes, metadata)
- And more

## Installation
```bash
git clone https://github.com/Zeronetsec/Woofind.git
cd Woofind
chmod +x install.sh
./install.sh

# for backup
./install.sh --backup
```

## Usage
```bash
woofind --misconfind <path>
woofind --dumpstring <string>
woofind --decode <string|file>
woofind --procinfo
woofind --sysinfo
```
And more commands.

## License
This project is licensed under the MIT License. <br>

<!-- Copyright (c) 2026 Zeronetsec -->