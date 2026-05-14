import os
import re

def remove_comments(file_path):
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()

    # Regex untuk mencocokkan string literal (agar tidak menghapus // di dalam string/URL)
    # dan mencocokkan komentar // serta /* */
    pattern = r'("(?:[^"\\\\]|\\\\.)*"|\'(?:[^\'\\\\]|\\\\.)*\')|//.*|/\*[\s\S]*?\*/'
    
    def replace(match):
        # Jika group(1) ada, berarti itu adalah string literal, kembalikan apa adanya
        if match.group(1):
            return match.group(1)
        # Jika tidak, itu adalah komentar, kembalikan string kosong
        return ""

    new_content = re.sub(pattern, replace, content)

    with open(file_path, 'w', encoding='utf-8') as f:
        f.write(new_content)

def main():
    # Direktori yang akan diproses
    target_dirs = ['cmd', 'internal', 'config']
    
    for target_dir in target_dirs:
        if not os.path.exists(target_dir):
            continue
            
        for root, dirs, files in os.walk(target_dir):
            for file in files:
                if file.endswith('.go'):
                    file_path = os.path.join(root, file)
                    print(f"Cleaning: {file_path}")
                    remove_comments(file_path)

if __name__ == "__main__":
    main()
