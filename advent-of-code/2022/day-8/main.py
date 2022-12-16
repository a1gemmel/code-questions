import fileinput

trees = []

for line in fileinput.input():
    trees.append(line.strip())

visibilityMap = {}


def part1():

    for x in range(len(trees)):
        maxSoFar = -1
        for y in range(len(trees[0])):
            if int(trees[x][y]) > maxSoFar:
                visibilityMap[str(x) + "-" + str(y)] = True
                maxSoFar = int(trees[x][y])

    for y in range(len(trees[0])):
        maxSoFar = -1
        for x in range(len(trees)): 
            if int(trees[x][y]) > maxSoFar:
                visibilityMap[str(x) + "-" + str(y)] = True
                maxSoFar = int(trees[x][y])


    for x in range(len(trees)):
        maxSoFar = -1
        for y in reversed(range(len(trees[0]))):
            if int(trees[x][y]) > maxSoFar:
                visibilityMap[str(x) + "-" + str(y)] = True
                maxSoFar = int(trees[x][y])

    for y in range(len(trees[0])):
        maxSoFar = -1
        for x in reversed(range(len(trees))): 
            if int(trees[x][y]) > maxSoFar:
                visibilityMap[str(x) + "-" + str(y)] = True
                maxSoFar = int(trees[x][y])

    print(len(visibilityMap))


def scoreView(x : int, y : int):
    height = int(trees[x][y])
    
    above = 0
    i = x - 1
    while i >= 0:
        above += 1
        if int(trees[i][y]) >= height:
            break
        i -= 1

    below = 0
    i = x + 1
    while i < len(trees):
        below += 1
        if int(trees[i][y]) >= height:
            break
        i += 1

    left = 0
    i = y - 1
    while i >= 0:
        left += 1
        if int(trees[x][i]) >= height:
            break
        i -= 1

    right = 0
    i = y + 1
    while i < len(trees):
        right += 1
        if int(trees[x][i]) >= height:
            break
        i += 1

    return above * below * left * right
    

def part2():
    maxScore = 0

    for x in range(len(trees)):
        for y in range(len(trees[0])):
            score = scoreView(x,y)
            if score > maxScore:
                maxScore = score

    print(maxScore)


part2()