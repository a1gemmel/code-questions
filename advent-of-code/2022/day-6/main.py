import fileinput

distinctCount = 14 # 4 for start of packet (part 1), 14 for start of message (part 2)


for line in fileinput.input():
    
    end = 0
    freq = {}
    while True:
        
        freq[line[end]] = freq.get(line[end], 0) + 1

        if end >=distinctCount:
            oldStart = line[end-distinctCount]
            toRemove = freq[oldStart]
            if toRemove > 1:
                freq[oldStart] -= 1
            else:
                del freq[oldStart]

        if len(freq) == distinctCount:
            print(end+1)
            break

        end += 1

