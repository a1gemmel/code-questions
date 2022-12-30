from __future__ import annotations
import fileinput
from typing import List

paths = []

class coord:
    def __init__(self, x : int, y : int):
        self.x = x
        self.y = y

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def down(self) -> coord:
        return coord(self.x, self.y + 1)

    def left(self)-> coord:
        return coord(self.x - 1, self.y + 1)

    def right(self)-> coord:
        return coord(self.x + 1, self.y + 1)
    
    def at_rest(self, map : set) -> bool:
        return self.down() in map and self.left() in map and self.right() in map

    def fall(self, map : Map) -> coord:
        if self.down() not in map:
            return self.down()
        if self.left() not in map:
            return self.left()
        if self.right() not in map:
            return self.right()
        return self

    def move_towards(self, other) -> coord:
        if other.x > self.x:
            return coord(self.x + 1, self.y)
        if other.x < self.x:
            return coord(self.x - 1, self.y)
        if other.y > self.y:
            return coord(self.x, self.y + 1)
        if other.y < self.y:
            return coord(self.x, self.y - 1)
        return self

    def __repr__(self):
        return "(" + repr(self.x) + "," + repr(self.y) + ")"

    def __hash__(self):
        return hash((self.x, self.y))

class Map:
    def __init__(self):
        self.solids = set()
        self.minX = 500
        self.maxX = 500
        self.maxY = 0

    def add(self, c : coord, wall=False):
        self.solids.add(c)
        if c.x < self.minX:
            self.miX = c.x
        elif c.x > self.maxX:
            self.maxX = c.x
        if c.y > self.maxY and wall:
            self.maxY = c.y

    def __contains__(self, obj):
        return obj in self.solids or obj.y == self.floor()

    def out_of_range(self, obj):
        return obj.x < self.minX or obj.y > self.maxX

    def floor(self):
        return self.maxY + 2

    def __repr__(self):
        return repr(self.solids)


m = Map()

for line in fileinput.input():
    path = line.strip().split("->")
    coords : List[coord] = []
    for c in path:
        [x,y] = c.strip().split(",")
        coords.append(coord(int(x),int(y)))

    start : coord = coords[0]
    rest = coords[1:]
    
    while len(rest) > 0:
        m.add(start)
        start = start.move_towards(rest[0])
        if start == rest[0]:
            rest = rest[1:]

    m.add(start, wall=True)


print(m)


sands = 0

while coord(500,0) not in m:
    sand = coord(500, 0)
    moved = sand.fall(m)

    while sand != moved:
        sand = moved
        moved = sand.fall(m)

    m.add(sand)
    sands += 1

print(sands)
