import os
import re
import sys
from pathlib import Path

script_dir = Path(__file__).resolve().parent
project_root = script_dir.parent.parent

input_file = project_root / "module" / "capability" / "cap_map.go"
output_file = project_root / "module" / "list" / "list" / "caplists.txt"

def generate_caplist():
    if not os.path.exists(input_file):
        print(f"\x1b[1;31m[!] \x1b[0mFile: \x1b[0;32m{input_file} \x1b[0mnot found!")
        sys.exit(1)

    with open(input_file, "r", encoding="utf-8") as f:
        content = f.read()

    pattern = r'"([^"]+)":\s*(\d+)'
    matches = re.findall(pattern, content)

    if not matches:
        print("\x1b[0;33m[!] \x1b[0mNothing capability found!")
        return

    lines = []
    for cap, num in matches:
        line = f"\x1b[1;90m* \x1b[0;32m{num} \x1b[1;90m-> \x1b[0;32m{cap}\x1b[0m"
        lines.append(line)

    with open(output_file, "w", encoding="utf-8") as f:
        f.write("\n".join(lines) + "\n")

    print(f"\x1b[0;32m[+] \x1b[0mGenerated: \x1b[0;32m{len(lines)} \x1b[1;90m-> \x1b[0;32m{output_file}\x1b[0m")

if __name__ == "__main__":
    generate_caplist()