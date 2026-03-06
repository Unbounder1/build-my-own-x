from typing import List
from copy import deepcopy

def checkInfection(grid: List[List[int]]) -> bool:
    for fl in grid:
        for st in fl:
            if (st == 'X'):
                return False
            
    return True

# returns value and dead count
def simulateSpread(grid: List[List[int]], state: List[List[int]], T: int, D: int, K: int) -> int:
    
    n = len(grid)
    m = len(grid[0])
    prev = deepcopy(grid)
    deathCount = 0
    
    for i in range(n):
        for j in range(m):
            
            ## Skip immune or dead
            if grid[i][j] == 'I' or grid[i][j] == '#':
                continue 

            ## Check infectedness

            leftj = max(0, j - 1)
            bottomi = max(0, i - 1)
            rightj = min(m - 1, j + 1)
            topi = min(n - 1, i + 1)

            # i, j
            directions = [[i, leftj], [bottomi, j], [i, rightj], [topi, j],
                            [bottomi, leftj], [topi, leftj], [topi, rightj], [bottomi, rightj]]
            
            infectedCount = 0
            for direction in directions:
                if prev[direction[0]][direction[1]] == 'X':
                    infectedCount+=1

            if (infectedCount >= K):
                grid[i][j] = '#'
                deathCount+=1
                continue
            # check plant type
            match prev[i][j]:
                case '.':
                    
                    if (infectedCount >= T):

                        grid[i][j] = 'X'
                        state[i][j] = D

                case 'X':
                    state[i][j] -= 1
                    if (state[i][j] == -1):
                        grid[i][j] = 'I'
    
    return deathCount


def epidemicSpreadDays(grid: List[List[int]], T: int, D: int, K: int) -> int:
    if (not len(grid) or not len(grid[0])):
        return
    
    count = 0
    totalDeaths = 0
    state = deepcopy(grid)

    # fill out states (-1 ignore, 0 switch to immune)
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == 'X':
                state[i][j] = D
            else:
                if grid[i][j] == '#':
                    totalDeaths += 1
                state[i][j] = -1
    
    while (True):
        # Add state tracking for infection time
        totalDeaths += simulateSpread(grid, state, T, D, K)
        if checkInfection(grid):
            return count, totalDeaths
        count+=1

def main():
    # 2d grid input
    grid = [['.', 'X', 'X','I'],
            ['.', 'X', 'X','I'],
            ['#', 'X', '.','X'],
            ['.', '.', '.','.']]
    T = 2
    D = 1
    K = 4

    days, deaths = epidemicSpreadDays(grid, T, D, K)

    print("There were " + str(days) + " days until the epidemic ended, and " + str(deaths) + " deaths" )

    return 1
    
if __name__ == "__main__":
    main()