import fileinput
from typing import List

class Valve:
    def __init__(self, name :str, rate : int, neighbours: List[str]):
        self.name = name
        self.rate = rate
        self.neighbours = neighbours

    def __repr__(self):
        return "(" + self.name + " " + repr(self.rate) + " " + repr(self.neighbours) + ")"


valves = {}

for line in fileinput.input():
    tokenized = line.strip().split(" ")
    name = tokenized[1]
    rate = int(tokenized[4].rstrip(";").split("=")[1])  
    neighbours = list(map(lambda s : s.rstrip(","), tokenized[9:]))

    valves[name] = Valve(name, rate, neighbours)

#print(valves)

class Position:
    def __init__(self, nodes : List[Valve], rate : int, opened : set, score : int):
        self.nodes = nodes
        self.rate = rate
        self.opened = opened
        self.score = score

positions = [Position(valves["AA"], 0, set(), 0)]
best_at_node = {}

for time in range(1, 31):
    new_positions = []

    print(len(positions), " positions at time ", time)

    for pos in positions:
        
        for node in pos.nodes:

            if node.name not in pos.opened and pos.node.rate > 0:
                s = pos.opened.copy()
                s.add(pos.node.name)
                new_positions.append(Position(pos.node, pos.rate + pos.node.rate, s, pos.score + pos.rate))
            for n in pos.node.neighbours:
                new_positions.append(Position(valves[n], pos.rate, pos.opened, pos.score + pos.rate))

    options = []
    for pos in new_positions:
        [bestRate, bestScore] = best_at_node.get(pos.node.name, [0,0])
        if pos.rate >= bestRate or pos.score >= bestScore:
            options.append(pos)
            best_at_node[pos.node.name] = [pos.rate, pos.score]

    positions = options

max_score = 0

for pos in positions:
    max_score = max(max_score, pos.score)

print(max_score)




