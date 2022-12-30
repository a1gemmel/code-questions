import fileinput
from typing import List
import sys

class coord:
    def __init__(self, x : int, y: int):
        self.x = x
        self.y = y

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

class sensor:
    def __init__(self, sX: int, sY : int, bX: int, bY: int):
        self.location = coord(sX, sY)
        self.beacon = coord(bX, bY)

    def distance(self, other):
        return abs(self.location.x - other.x) + abs(self.location.y - other.y)

    def no_beacon(self, c : coord):
        return self.distance(c) <= self.distance(self.beacon) and c != self.beacon

    def minX(self):
        return self.location.x - self.distance(self.beacon)

    def maxX(self):
        return self.location.x + self.distance(self.beacon)


sensors : List[sensor] = []
rangeMin = sys.maxsize
rangeMax = - sys.maxsize



for line in fileinput.input():
    tokenized = line.strip().split(" ")
    sX = int(tokenized[2].split("=")[1].rstrip(","))
    sY = int(tokenized[3].split("=")[1].rstrip(":"))

    bX = int(tokenized[8].split("=")[1].rstrip(","))
    bY = int(tokenized[9].split("=")[1])

    s = sensor(sX, sY, bX, bY)
    rangeMin = min(s.minX(), rangeMin)
    rangeMax = max(s.maxX(), rangeMax)

    sensors.append(s)



def part1():
    impossible_locations = 0
    y_check = 2000000

    for i in range(rangeMin - 1, rangeMax + 1):
        print("checking coord " + repr(i))
        c = coord(i, y_check)
        for s in sensors:
            if s.no_beacon(c):
                impossible_locations += 1
                break

    print(impossible_locations)

def part2():

    search = 4000000
    y = 0
    while y <= search:
        x = 0
        min_of_maxes = sys.maxsize
        if y % 1000 == 0:
            print("checking Y " + repr(y))
        while x <= search:
            distance_differences = []
            for s in sensors:
                diff = s.distance(s.beacon) - s.distance(coord(x,y))
                distance_differences.append(diff)
            
            max_dist = max(distance_differences)
            min_of_maxes = min(min_of_maxes, max_dist)
            if max_dist < 0:
                print(x, y, x * 4000000 + y)
                return
            
            x += max(max_dist, 1)

        y += 1


part2()

