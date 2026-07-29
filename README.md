<!-- https://github.com/Zeronetsec/Woofind -->

<div align="center">
    <img src="https://img.shields.io/badge/Woofind-Version%200.1-blue?style=square&logo=go&v=1" />
    <img src="https://img.shields.io/badge/Supported%20OS-Linux-blue?style=square&logo=linux&v=1" />
    <a href="LICENSE">
        <img src="https://img.shields.io/badge/License-GPL--3.0-blue?style=square&logo=github&v=1" />
    </a>
</div>

# Woofind
Woofind is a CLI toolkit for system exploration, security auditing, and data inspection.

## Features
- Analyze running processes and resource usage.
- Detect misconfigurations and potential risks.
- Decode and analyze encoded data automatically.
- Extract insights from strings (entropy, hashes, metadata).
- And more features.

## Disclaimer
Please read [.docs/disclaimer.md](.docs/disclaimer.md) before using this tool. </br>
Use this software at your own risk. </br>
The author is not responsible for any damage, data loss, or issues that may result from its use.

## Installation
Quick install:
```bash
git clone https://github.com/Zeronetsec/Woofind
bash Woofind/install.sh
```
For more detailed installation and uninstallation instructions, see [.docs/install_and_uninstall.md](.docs/install_and_uninstall.md).


## Usage Example
```bash
woofind --pattern-scan mydir/
woofind --capability /usr/sbin/ 'cap_sys_ptrace:cap_setuid:cap_net_raw' --threads 100
woofind --useraudit --passwd mypasswd --shadow myshadow
woofind --decode susfile.txt --limit unlimit --disable 'rot13:morse'
woofind --owner /usr/bin/ --force
```
And more commands.

<!-- Copyright (c) 2026 Zeronetsec -->