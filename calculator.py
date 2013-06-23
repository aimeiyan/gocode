__author__ = 'feng'

OPS = '+-*/()'


def tokenize(expr):
    for op in OPS:
        expr = expr.replace(op, ' %s ' % op)

    return expr.split()


def high_precedence(op1, op2):
    def normanize(op):
        if op == '-': op = '+'
        if op == '/': op = '*'
        return op

    op1 = normanize(op1)
    op2 = normanize(op2)

    # '<'  => need pop, '>' delay computation

    table = {
        '+': {'+': '<', '*': '<', '(': '>', ')': '<'},
        '*': {'+': '>', '*': '<', '(': '>', ')': '<'},
        '(': {'+': '>', '*': '>', '(': '>', ')': '>'},
        ')': {'+': '<', '*': '<', '(': '<', ')': '<'},
    }
    return table[op1][op2] == '>'


def in2pos(expr):
    results = []
    op_stack = []
    for c in tokenize(expr):
        if c in OPS:
            while op_stack:
                top = op_stack[len(op_stack) - 1]
                # current operator has higher precedence, delay computation
                if high_precedence(c, top):
                    break
                else:
                    # current operator has lower precedence, we can compute previous one
                    if top not in "()":
                        results.append(top)
                    op_stack.pop()
                    # only pop one '(' if current is '('
                    if c == ')' and top == '(':
                        break
            op_stack.append(c)
        else:
            results.append(c)

    while len(op_stack) != 0:
        top = op_stack.pop()
        if top not in "()":
            results.append(top)

    return " ".join(results).strip()


def compute(expr):
    num_stack = []
    for token in tokenize(in2pos(expr)):
        if token in OPS:
            op1 = num_stack.pop()
            op2 = num_stack.pop()
            if token == '+':
                num_stack.append(op1 + op2)
            elif token == '*':
                num_stack.append(op1 * op2)
            elif token == '-':
                num_stack.append(op2 - op1)
            elif token == '/':
                num_stack.append(op2 / op1)
        else:
            num_stack.append(int(token))
    return num_stack.pop()


if __name__ == "__main__":
    tests = [
        ("11+21-3", "11 21 + 3 -"),
        ("(1+2)*(3+4)", "1 2 + 3 4 + *"),
        ("1 + 2 * 3 -4", "1 2 3 * + 4 -"),
        ("(1+2)*3", "1 2 + 3 *"),
        ("1+(2+3)", "1 2 3 + +"),
        ("(3 + 4) * 4", "3 4 + 4 *"),
        ("1 + 2 * 3", "1 2 3 * +"),
        ("1*((1+3)*4)", "1 1 3 + 4 * *")
    ]
    ok = True
    for expr, right in tests:
        if in2pos(expr) != right:
            ok = False
            print("ERROR", expr, '=>', right, "but get:", in2pos(expr))
        else:
            print("OK", expr)

    if ok:
        for expr, right in tests:
            if compute(expr) != eval(expr):
                print("ERROR, expect", eval(expr), "but get:", compute(expr))
            else:
                print("compute OK", expr, eval(expr))
