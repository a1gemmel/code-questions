import fileinput 
from typing import Callable

isPart1 = False

def makeOp(symbol : str, amount : str) -> Callable[[int], int]:
    if amount != "old":
        if symbol == "*":
            return lambda n : n * int(amount)
        else: # +
            return lambda n : n + int(amount)
    else:
        if symbol == "*":
            return lambda n : n * n
        else: # +
            return lambda n : n + n


class monkey:
    def __init__(self, index, items, op, opAmnt, test, trueThrow, falseThrow):
        self.index = index
        self.items = items
        self.op = op
        self.opAmnt = opAmnt
        self.test = test
        self.trueThrow = trueThrow
        self.falseThrow = falseThrow

    def __repr__(self):
        return f'''
Monkey {self.index}:
    Starting items: {self.items}
    Operation: new = old {self.op} {self.opAmnt}
    Test: divisible by {self.test}
        If true: throw to monkey {self.trueThrow}
        If false: throw to monkey {self.falseThrow}
    '''

    def inspect(self):
        if len(self.items) == 0:
            return None, None
        item = makeOp(self.op, self.opAmnt)(self.items[0])
        if isPart1:
            item = item // 3
        self.items = self.items[1:]

        testRes = item % self.test == 0
        if testRes:
            return item, self.trueThrow
        return item, self.falseThrow

    def giveItem(self, item):
        self.items.append(item)


lines = []
for line in fileinput.input():
    lines.append(line.strip())


monkeys = []


for i in range(len(lines) // 7):
    monkeyInput = lines[i * 7: (i+1) * 7]

    items = monkeyInput[1].split(" ")[2:]
    items = list(map(lambda s : int(s.replace(",", "")), items))
    [operation, amount] = monkeyInput[2].split(" ")[4:]
    test = int(monkeyInput[3].split(" ")[3])
    trueThrow = int(monkeyInput[4].split(" ")[5])
    falseThrow = int(monkeyInput[5].split(" ")[5])

    monkeys.append(monkey(i, items, operation, amount, test, trueThrow, falseThrow))

inspectionCount = []    

divisor = 1

for m in monkeys:
    inspectionCount.append(0)
    divisor *= m.test

rounds = 20 if isPart1 else 10000

for i in range(rounds):
    print(i)
    for m in monkeys:
        item, nextMonkey = m.inspect()
        while item != None:
            if not isPart1:
                item = item % divisor
            monkeys[nextMonkey].giveItem(item)
            inspectionCount[m.index] += 1
            item, nextMonkey = m.inspect()

inspectionCount.sort(reverse=True)
print(inspectionCount)
print(inspectionCount[0] * inspectionCount[1])