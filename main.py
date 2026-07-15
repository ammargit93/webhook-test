import sys
import math
def solve(s):
    l = 0
    maxl = 0
    for i in range(len(s)):
        if s[i]=="#":
            l +=1
        else:
            l =0
        maxl= max(maxl,l)
    return math.ceil(maxl/2) 

lines = sys.stdin.read().split()
t = int(lines[0])
idx = 1
for _ in range(t):
    n = int(lines[idx])
    idx += 1
    s = lines[idx]
    idx += 1
    print(solve(s))

