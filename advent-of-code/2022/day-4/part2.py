import fileinput

overlaps = 0

for line in fileinput.input():
    [range1, range2] = line.strip().split(",")

    [a1, a2] = range1.split("-")
    a1 = int(a1)
    a2 = int(a2)
    [b1, b2] = range2.split("-")
    b1 = int(b1)
    b2 = int(b2)

    if not (b1 > a2 or b2 < a1):
        overlaps+= 1

print(overlaps)