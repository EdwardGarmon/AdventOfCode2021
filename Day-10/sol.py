

import math


def score(p):
    if(p=="]"):
        return 57
    if(p==")"):
        return 3
    if(p =="}"):
        return 1197
    if(p==">"):
        return 25137


def scoreLine(l, closeMap):
    stack = []
    for c in l:
        if(c in closeMap.keys()):
            stack.append(c)
        else:
            o = stack.pop()
            if(closeMap[o] != c):
                return score(c) , []
    return (0, stack)

def point(c):
    if(c == ")"):
        return 1
    if(c == "]"):
        return 2
    if(c == "}"):
        return 3
    if(c == ">"):
        return 4
    print(c)

def scoreCompletion(c):
    score = 0
    for m in list(c): 
        score *= 5
        score += point(m)
    return score

def part1():
    closeMap = { "{":"}" , "(" : ")", "<" : ">" , "[" : "]"}
    totalScore = 0
    complScores = []
    for l in open("input.txt", 'r').read().split("\n"):
        score, stack = scoreLine(l,closeMap)
        totalScore += score
        stack.reverse()
        correction = "".join((list(map(lambda x : closeMap[x], stack))))
        if(len(correction) > 0):
            complScores.append(scoreCompletion(correction))

    complScores.sort()
    print(complScores[math.floor(len(complScores) / 2)])
    print(totalScore)
        




part1()