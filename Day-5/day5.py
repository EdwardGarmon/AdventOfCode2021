
global count
count = 0

def getLines():
    f = open('input.txt', 'r') 
    lines = [ [ [int(num) for num in ps.split(",")] for ps in p.split("->") ] for p in f.read().split("\n")]
    return lines



def isVert(line):
    return line[0][0] == line[1][0]

def isHorz(line):
    return line[0][1] == line[1][1]

def isHV(line):
    return isHorz(line) or isVert(line)

def incDict(key,d):
    global count
    if(key in d):
        d[key] = d[key] + 1
        if(d[key]==2):
            count += 1
    else:
        d[key] = 1

def part1():
    global count
    count = 0
    grid = {}
    lines = getLines()

    for line in lines:

        x1 = int(line[0][0])
        y1 = int(line[0][1])
        x2 = int(line[1][0])
        y2 = int(line[1][1])

        if(isHorz(line)):
            for x in range(min(x1,x2) , max(x2, x1) + 1):
                incDict((x,y1), grid)
        elif(isVert(line)):
            for y in range(min(y1,y2) , max(y2, y1) + 1):
                incDict((x1,y), grid)
        else:
            continue

    print(count)


def part2():
    global count
    count = 0
    grid = {}
    lines = getLines()

    for line in lines:

        x1 = int(line[0][0])
        y1 = int(line[0][1])
        x2 = int(line[1][0])
        y2 = int(line[1][1])

        if(isHorz(line)):
            for x in range(min(x1,x2) , max(x2, x1) + 1):
                incDict((x,y1), grid)
        elif(isVert(line)):
            for y in range(min(y1,y2) , max(y2, y1) + 1):
                incDict((x1,y), grid)
        else:
            xc = x2 - x1
            yd = y2 - y1

            xd  = xc / abs(xc)
            yd /= abs(yd)
            
            x = x1
            y = y1
            
            for i in range(0,abs(xc) + 1 ):
                incDict((x,y), grid)
                print((x,y))
                x += xd
                y += yd

    print(count)    


part2() 



         




        



        



        



        

