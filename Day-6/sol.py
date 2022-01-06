

def getFMap():
    f = open('input.txt', 'r') 
    arr = [int(p) for p in f.read().split(",")]
    fMap = {p : 0 for p in range(0,9)}
    for l in arr:
        fMap[l] += 1
    return fMap

    

def part1():

    p = getFMap()
    for _ in range(0,256):
        newMap = {p : 0 for p in range(0,9)}
        for ps in p.items():
            if(ps[0]==0):
                newMap[8] += ps[1]
                newMap[6] += ps[1]
            else:
                newMap[ps[0]-1] +=ps [1]
        p = newMap
    print(p)
    s = sum(p.values())
    print(s)
   


part1()