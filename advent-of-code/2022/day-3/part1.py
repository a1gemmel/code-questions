import fileinput

prioritySum = 0

for rucksack in fileinput.input():
    part1 = rucksack[0:len(rucksack) // 2]
    part2 = rucksack[len(rucksack) // 2:]
    for i in part1:
        if part2.find(i) != -1:
            if 'a' <= i and i <= 'z':
                prioritySum += ord(i) - ord('a') + 1
            else:
                prioritySum += ord(i) - ord('A') + 27
            break

print(prioritySum)
