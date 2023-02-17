import os


def get_lines_amount(filepath: str) -> int:
    with open(filepath, "r", encoding="utf8") as f:
        # return len([l for l in f.read().splitlines() if l])
        return len(f.read().splitlines())


def main():
    total = 0

    total += get_lines_amount("cmd/twitter-tools/main.go")
    
    root = "internal"
    dirs = os.listdir(root)
    for dir in dirs:
        for file in os.listdir(f"{root}/{dir}"):
            filepath = f"{root}/{dir}/{file}"
            if filepath.endswith("_test.go"):
                continue

            total += get_lines_amount(filepath)
    
    print(total)

    

if __name__ == "__main__":
    main()