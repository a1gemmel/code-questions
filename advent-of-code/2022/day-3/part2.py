import fileinput


sacks = []
prioritySum = 0
for rucksack in fileinput.input():
    sacks.append(rucksack.strip())

    if len(sacks) == 3:

        freq = {}
        for i in range(3):
            for item in sacks[i]:
                if  i == 0 or freq.get(item) == i - 1:
                    freq[item] = i
            
        for key in freq:
            val = freq[key]
            if val == 2:
                if 'a' <= key and key <= 'z':
                    prioritySum += ord(key) - ord('a') + 1
                else:
                    prioritySum += ord(key) - ord('A') + 27
                break

        sacks = []

print(prioritySum)

