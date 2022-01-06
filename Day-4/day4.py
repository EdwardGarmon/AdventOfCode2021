with open('AOC_day4.txt', 'r') as f:
    num_data, *board_data = f.read().split('\n\n')

def create_boards(num_data, board_data):
    nums = map(int, num_data.split(','))
    boards_meta = []
    num_set = {}
    for n,board in enumerate(board_data):
        board = [row.split() for row in board.split('\n')]
        boards_meta.append({0:0,1:0,2:0,3:0,4:0, 5:0,6:0,7:0,8:0,9:0,'v':set({})})
        for x,row in enumerate(board):
            for y,num in enumerate(row):
                num = int(num)
                if(num in num_set):
                    num_set[num].append((n,x,y))
                else:
                    num_set[num] = [(n,x,y)]
                boards_meta[n]['v'].add(num)

    return nums, num_set, boards_meta

def get_winning_score(num, board):
    return (sum(sum(group) for group in board) - num) * num

def AOC_day4_pt1():
    nums, num_set, boards_meta = create_boards(num_data, board_data)
    for n in nums:
        if n not in num_set :
            continue
        for p in num_set[n]:
            print(n, p)
            boards_meta[p[0]][p[1]] = boards_meta[p[0]][p[1]] + 1
            boards_meta[p[0]][p[2] + 5] = boards_meta[p[0]][p[2] + 5 ] + 1
            
            boards_meta[p[0]]['v'].remove(n)

            if (boards_meta[p[0]][p[1]] > 4 or boards_meta[p[0]][p[2] + 5]>4):
                return sum(boards_meta[p[0]]['v']) * n


def AOC_day4_pt2():
    nums, num_set, boards_meta = create_boards(num_data, board_data)
    last_score = 0
    win_set = set({})
    for n in nums:
        if n not in num_set :
            continue
        for p in num_set[n]:

            if(n not in boards_meta[p[0]]['v'] or p[0] in win_set):
                continue

            boards_meta[p[0]][p[1]] = boards_meta[p[0]][p[1]] + 1
            boards_meta[p[0]][p[2] + 5] = boards_meta[p[0]][p[2] + 5 ] + 1
            
            boards_meta[p[0]]['v'].remove(n)

            if (boards_meta[p[0]][p[1]] > 4 or boards_meta[p[0]][p[2] + 5]>4):
                last_score = sum(boards_meta[p[0]]['v']) * n
                win_set.add(p[0])

    return last_score



print(AOC_day4_pt2())
            
    

