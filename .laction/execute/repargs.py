import glob
import json
import os
import sys
from pathlib import Path

script_dir = Path(__file__).resolve().parent
project_root = script_dir.parent.parent

list_dir = project_root / "module" / "list" / "list"
metafile = project_root / "module" / "helper" / "metadata" / "list.json"

def update_list_metadata():
    if not os.path.exists(list_dir):
        print(f"\x1b[1;31m[!] \x1b[0mFolder: \x1b[0;32m{list_dir} \x1b[0mnot found!")
        sys.exit(1)

    if not os.path.exists(metafile):
        print(f"\x1b[1;31m[!] \x1b[0mFile: \x1b[0;32m{metafile}\x1b[0m")
        sys.exit(1)

    txt_files = glob.glob(
        os.path.join(list_dir, "*.txt"),
    )

    list_names = [
        os.path.splitext(
            os.path.basename(f),
        )[0] for f in txt_files
    ]
    list_names.sort()

    if not list_names:
        print(f"\x1b[0;33m[!] \x1b[0mFile: \x1b[0;32m*.txt \x1b[0mnot found in \x1b[0;32m{list_dir}\x1b[0m")
        return

    args_value = f"<{'|'.join(list_names)}>"
    with open(metafile, "r", encoding="utf-8") as f:
        try:
            data = json.load(f)
        except json.JSONDecodeError as e:
            print(f"\x1b[0;31m[!] \x1b[0mFailed to read json file: \x1b[0;32m{e}\x1b[0m")
            sys.exit(1)

    data["Args"] = args_value
    with open(metafile, "w", encoding="utf-8") as f:
        json.dump(data, f, indent=4)
        f.write("\n")

    print(f"\x1b[0;32m[+] \x1b[0mUpdate: \x1b[0;32m{args_value} \x1b[1;90m-> \x1b[0;32m{metafile}\x1b[0m")

if __name__ == "__main__":
    update_list_metadata()