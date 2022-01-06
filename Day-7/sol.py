def getArr():
    f = open('input.txt', 'r') 
    arr = [int(p) for p in f.read().split(",")]
    return arr


def part1():
    crabs = getArr()

    mincost = -1

    for target in range(max(crabs)):
        cost = 0
        for move in crabs:
            steps = abs(move-target)
            fuel_burn = lambda x, p: abs(x-p)*(abs(x-p) + 1)//2
            cost += fuel_burn(move,target)
        
        if(mincost == -1):
            mincost = cost
        elif(cost < mincost):
            mincost = cost 
    
    print(mincost)

part1()