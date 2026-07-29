import os
import re
import sys
from pathlib import Path

script_dir = Path(__file__).resolve().parent
project_root = script_dir.parent.parent

input_file = project_root / "module" / "decode" / "brute.go"
output_file = project_root / "module" / "list" / "list" / "decoders.txt"

def generate_decoderlist():
    if not os.path.exists(input_file):
        print(f"\x1b[1;31m[!] \x1b[0mFile: \x1b[0;32m{input_file} \x1b[0mnot found!")
        sys.exit(1)

    with open(input_file, "r", encoding="utf-8") as f:
        content = f.read()

    block_match = re.search(
        r"decoders\s*:=\s*\[\]struct[\s\S]*?\{([\s\S]*?)\n\t*\}", content
    )
    block_content = block_match.group(1) if block_match else content

    pattern = r'\{\s*"([^"]+)"'
    decoders = re.findall(pattern, block_content)

    if not decoders:
        print("\x1b[0;33m[!] \x1b[0mNothing decoders found!")
        return

    lines = []
    for name in decoders:
        line = f"\x1b[1;90m* \x1b[0;32m{name}\x1b[0m"
        lines.append(line)

    with open(output_file, "w", encoding="utf-8") as f:
        f.write("\n".join(lines) + "\n")

    print(f"\x1b[0;32m[+] \x1b[0mGenerated: \x1b[0;32m{len(lines)} \x1b[1;90m-> \x1b[0;32m{output_file}\x1b[0m")

if __name__ == "__main__":
    generate_decoderlist()