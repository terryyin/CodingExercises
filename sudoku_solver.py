from itertools import product
def checkio(grid):
    for i, j in product(range(9), range(9)):
        if grid[i][j] == 0: break
    if grid[i][j] != 0: return grid
    for choice in set(range(1, 10)) - set(grid[i]+[grid[x][j] for x in range(9)] + [grid[x+i//3*3][y+j//3*3] for x, y in product((0,1,2), (0,1,2))]):
        grid[i][j] = choice
        solution = checkio(grid)
        if solution: return solution
        grid[i][j] = 0
