import fileinput



def toList(el):
    if type(el) == list:
        return el
    return [el]

def ordered(left, right):
    #print(left, right)

    if len(left) == 0:
        if len(right) == 0:
            return None
        return True
    if len(right) == 0:
        return False
    if type(left[0]) == list or type(right[0]) == list:
        inOrder = ordered(toList(left[0]), toList(right[0]))
        if inOrder != None:
            return inOrder
        return ordered(left[1:], right[1:])
    if left[0] > right[0]:
        return False
    if left[0] < right[0]:
        return True
    return ordered(left[1:], right[1:])

def part1():
    lines = []
    for line in fileinput.input():
        lines.append(line.strip())

    rightOrderIndexSum = 0

    for i in range(len(lines) // 3):
        left = eval(lines[0])
        right =eval(lines[1])
        lines = lines[3:]
        print(i + 1, left, right,  ordered(left, right))
        if ordered(left, right):
            rightOrderIndexSum += i + 1

    print(rightOrderIndexSum)

def part2():
    packets = []

    divider1 = [[2]]
    divider2 = [[6]]
    divider1Index = 1
    divider2Index = 2

    for line in fileinput.input():
        if line.strip() == "":
            continue
        packet = eval(line.strip())
        if ordered(packet, divider1):
            divider1Index += 1
        if ordered(packet, divider2):
            divider2Index += 1

    print(divider1Index * divider2Index)


if __name__ == '__main__':
    # part1()
    part2()
