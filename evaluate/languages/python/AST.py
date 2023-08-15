import ast
import sys
from collections import Counter

def ParseToAST(content):
    tree = ast.parse(content)
    return ast.dump(tree)

def EvaluateAST(tree):
    ident_counter = Counter()

    for node in ast.walk(tree):
        if isinstance(node, ast.FunctionDef):
            ident_counter[node.name] += 1
        elif isinstance(node, ast.Name):
            ident_counter[node.id] += 1

    ident_length_sum = sum(len(ident) * count for ident, count in ident_counter.items())
    ident_count = sum(count for count in ident_counter.values())

    average_ident_length = ident_length_sum / ident_count if ident_count > 0 else 0

    return ident_counter, average_ident_length


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python ParseToAST.py <file>")
        sys.exit(1)
    if sys.argv[1] == "evaluate":
        count, ave = EvaluateAST(ast.parse(sys.argv[2]))
        print(len(count), ave)
        sys.exit(0)
    content = sys.argv[1]
    print(ParseToAST(content))
