import fileinput


pointsMap = {
    "X": 1,
    "Y": 2,
    "Z": 3,
}

winMap = {
    "X" : "C",
    "Y" : "A",
    "Z": "B",
}

tieMap = {
    "X": "A",
    "Y": "B",
    "Z": "C",
}


points = 0
for line in fileinput.input():
    [them, us] = line.strip().split(" ")
    points += pointsMap[us]
    if them == winMap[us]:
        points += 6
    elif them == tieMap[us]:
        points += 3

print(points)

