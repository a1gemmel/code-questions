from typing import List
import fileinput

map = []

start = []
ends = []

def hash(coords : List[int]) -> str:
    return str(coords[0]) + "x" + str(coords[1])

y = 0
for line in fileinput.input():
    gridLine = []
    x = 0
    for char in line.strip():
        if char == "S":
            gridLine.append(0)
            ends.append([x, y])
        elif char == "E":
            gridLine.append(25)
            start = [x,y]
        else:
            gridLine.append(ord(char) - ord("a"))
            if char == "a":
                ends.append([x, y])
        x += 1
    
    map.append(gridLine)
    y+=1

visited = {hash(start): 0}
toVisit = [start]

while len(toVisit) > 0:
    [x,y] = toVisit[0]
    toVisit = toVisit[1:]
    for [nX, nY] in [[x + 1, y], [x-1, y], [x, y-1], [x, y+1]]:
        if nX >= 0 and nY >= 0 \
        and nX < len(map[0]) and nY < len(map) \
        and not visited.get(hash([nX, nY])) \
        and map[y][x] -1 <= map[nY][nX]:

            visited[hash([nX, nY])] = visited[hash([x,y])] + 1
            toVisit.append([nX, nY])


minLen = None

for end in ends:
    pathLen = visited.get(hash(end))
    if minLen == None or (pathLen != None and  pathLen < minLen):
        minLen = visited[hash(end)]

print(minLen)