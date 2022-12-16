import fileinput
import heapq

heap = []
current = 0
for line in fileinput.input():
    if line == "\n":
        if len(heap) == 3:
            heapq.heappush(heap, max(current, heapq.heappop(heap)))
        else:
            heapq.heappush(heap, current)
        current = 0
        continue
    current += int(line.strip())

print(heapq.heappop(heap) + heapq.heappop(heap) + heapq.heappop(heap))