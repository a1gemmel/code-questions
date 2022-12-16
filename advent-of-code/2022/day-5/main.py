import fileinput

'''
            [L] [M]         [M]    
        [D] [R] [Z]         [C] [L]
        [C] [S] [T] [G]     [V] [M]
[R]     [L] [Q] [B] [B]     [D] [F]
[H] [B] [G] [D] [Q] [Z]     [T] [J]
[M] [J] [H] [M] [P] [S] [V] [L] [N]
[P] [C] [N] [T] [S] [F] [R] [G] [Q]
[Z] [P] [S] [F] [F] [T] [N] [P] [W]


'''

# stacks = [
# [],
# ['Z', 'N'],
# ['M', 'C', 'D'],
# ['P'],
# ]

stacks = [
    [], # empty to make stacks 1-indexed
    ['Z', 'P', 'M', 'H', 'R'],
    ['P', 'C', 'J', 'B'],
    ['S', 'N', 'H', 'G', 'L', 'C', 'D'],
    ['F', 'T', 'M', 'D', 'Q', 'S', 'R', 'L'],
    ['F', 'S', 'P', 'Q', 'B', 'T', 'Z', 'M'],
    ['T', 'F', 'S', 'Z', 'B', 'G'],
    ['N', 'R', 'V'],
    ['P', 'G', 'L', 'T', 'D', 'V', 'C', 'M'],
    ['W', 'Q', 'N', 'J', 'F', 'M', 'L'],
]

is9001Machine = True # false for part 1, true for part 2

for line in fileinput.input():
    [move, x, frum, y, to, z] = line.strip().split(" ")

    if not is9001Machine:
        for i in range(int(x)):
            stacks[int(z)].append(stacks[int(y)].pop())
    else:
        fromStack = stacks[int(y)]
        stacks[int(z)].extend(fromStack[len(fromStack) - int(x):])
        stacks[int(y)] = fromStack[0: len(fromStack) - int(x)]

top = ""
for stack in stacks:
    if len(stack) > 0:
        top += stack[len(stack) - 1]

print(top)