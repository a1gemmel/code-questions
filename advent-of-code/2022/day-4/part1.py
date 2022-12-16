import fileinput

fullyContained = 0

for line in fileinput.input():
    [range1, range2] = line.strip().split(",")

    [a1, a2] = range1.split("-")
    a1 = int(a1)
    a2 = int(a2)
    [b1, b2] = range2.split("-")
    b1 = int(b1)
    b2 = int(b2)

    if (a1 <= b1 and a2 >= b2) or (b1 <= a1 and b2 >= a2):
        fullyContained+= 1

print(fullyContained)