import tokenize
import io
import sys

def CountTokens(content):
    tokens = tokenize.generate_tokens(io.StringIO(content).readline)
    return len(list(tokens))

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python CountTokens.py <file>")
        sys.exit(1)
    content = sys.argv[1]
    print(CountTokens(content))
