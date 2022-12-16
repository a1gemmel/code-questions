from __future__ import annotations
import fileinput

class file:
    def __init__(self, name,size):
        self.name :str = name
        self.size :int = size

class  directory:
    def __init__(self, parent : directory | None, name : str):
        self.name :str = name
        self.size :int = 0
        self.parent : directory = parent if parent != None else self
        self.subdirs : list[directory] = []
        self.files : list[file] = []

    def cd(self, dir : str) -> directory:
        if dir == "/":
            if self.parent == self:
                return self 
            return self.parent.cd(dir)
        elif dir == "..":
            return self.parent 
        else:
            for subdir in self.subdirs:
                if subdir.name == dir:
                    return subdir
        return self # this is never reached but required for typing


    def add_file(self, name : str, size : int):
        self.files.append(file(name, size))
        self.add_size(size)

    def add_size(self, size : int):
        self.size += size
        if self.parent != self:
            self.parent.add_size(size)

    def add_subdir(self, name :str):
        self.subdirs.append(directory(self, name))

    def string(self, indent = 0):
        whitespace = indent * " "
        print(whitespace + "/" + self.name + " (" + str(self.size) + ")")
        for file in self.files:
            print(whitespace + "--" + file.name + " " + str(file.size))
        for subdir in self.subdirs:
            subdir.string(indent+2)
        



current : directory = directory(None, "")
root = current
commands = []
command = []
for line in fileinput.input():
    if  line.startswith("$"):
        commands.append(command)
        command = []            
    command.append(line.strip())

commands = commands[1:]
commands.append(command)
i = 0

for c in commands:
    cmd_tokens = c[0].split(" ")

    if cmd_tokens[1] == "cd":
        current = current.cd(cmd_tokens[2])
    else: # ls
        output = c[1:]
        for obj in output:
            [sizeOrDir, name] = obj.split(" ")
            if sizeOrDir == "dir":
                current.add_subdir(name)
            else:
                current.add_file(name, int(sizeOrDir))



def part1():
    totalSize = 0
    nodes = [root]

    root.string()

    while len(nodes) > 0:
        n = nodes.pop()
        if n.size <= 100000:
            totalSize += n.size
        nodes.extend(n.subdirs)

    print(totalSize)

def part2():
    targetDeleteSize = root.size - (70000000 - 30000000)
    minDeletable = root.size
    nodes = [root]

    root.string()

    while len(nodes) > 0:
        n = nodes.pop()
        if n.size >= targetDeleteSize:
            if n.size < minDeletable:
                minDeletable = n.size
            nodes.extend(n.subdirs)

    print(minDeletable) 

part2()