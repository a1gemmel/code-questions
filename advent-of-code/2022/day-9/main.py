import fileinput


headX = 0
headY = 0
tailX = 0
tailY = 0

tailPosMap = {"0*0": True}
chain = []

chainLen = 10 # 2 for part 1

for i in range(chainLen):
    chain.append([0,0])

def computeMove(diff : int) -> int:
    if diff == 0:
        return 0
    return diff // abs(diff)
  
for line in fileinput.input():
    [direction, length] = line.strip().split(" ")
    length = int(length)

    for i in range(length):
        # start by moving head
        if direction == "U":
            chain[0][1] += 1
        elif direction == "D":
            chain[0][1] -=1 
        elif direction == "L":
            chain[0][0] -= 1
        elif direction == "R":
            chain[0][0] += 1
        
        for el in range(1, len(chain)):
            [nodeX, nodeY] = chain[el]
            [parentX, parentY] = chain[el - 1]
            verticalDiff = parentY - nodeY
            horizontalDiff = parentX - nodeX

            while abs(verticalDiff) > 1 or abs(horizontalDiff) > 1:
                nodeX += computeMove(horizontalDiff)
                nodeY += computeMove(verticalDiff)

                if el == len(chain) - 1:
                    tailPosMap[str(nodeX) + "*" + str(nodeY)] = True

                verticalDiff = parentY - nodeY
                horizontalDiff = parentX - nodeX

            chain[el] = [nodeX, nodeY]    
            
        

print(len(tailPosMap))
