import fileinput


pointsMap = {
    "A": 1,
    "B": 2,
    "C": 3,
}

def modAdd(char, n):
    return chr((ord(char) - ord('A') + n) % 3 + ord('A'))

points = 0
for line in fileinput.input():
    [them, outcome] = line.strip().split(" ")
    if outcome == "X": # lose
        points += pointsMap[modAdd(them, 2)]
    elif outcome == "Y": # tie
        points += 3
        points += pointsMap[them]
    elif outcome == "Z": # win
        points += 6
        points += pointsMap[modAdd(them, 1)]

print(points)

