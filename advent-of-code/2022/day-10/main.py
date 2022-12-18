import fileinput


lines = []

for line in fileinput.input():
    lines.append(line.strip())

def part1():
    signalStrengthSum = 0
    curLine = 0
    addInProgress = False 
    X = 1

    for cycle in range(1, 221):
        if (cycle + 20) % 40 == 0:
            signalStrengthSum += cycle * X 
        line = lines[curLine]
        if line == "noop":
            curLine += 1
        elif addInProgress == False:
            addInProgress = True
        else:
            [addx, V] = line.split(" ")
            X += int(V)
            addInProgress = False
            curLine += 1

    print(signalStrengthSum)

def part2():
    curLine = 0
    addInProgress = False 
    X = 1

    for cycle in range(0, 241):
        if cycle % 40 == 0:
            print("")

        if abs((cycle) % 40 - X) <= 1:
            print("#", end="")
        else:
            print(" ", end="")

        line = lines[curLine]
        if line == "noop":
            curLine += 1
        elif addInProgress == False:
            addInProgress = True
        else:
            [addx, V] = line.split(" ")
            X += int(V)
            addInProgress = False
            curLine += 1

part2()