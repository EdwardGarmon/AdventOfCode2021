from collections import Counter

def readInput():
    f = open("input.txt",'r')
    text = [l.split("|") for l in f.read().split("\n")]
    return text


def part1():

    sizes = {2:1,4:4,3:7,7:8}


    

    print(sum(counts.values()))
    


