package main

import (
       "fmt"
       "time"
)

const (
    EMPTY_SPACE = " "
    LIVE_CHAR = "x"
    DEAD_CHAR = "_"
    ALIVE = true
)

const (
    MAX_NEXT_GENERATION = 3
    MIN_NEXT_GENERATION = 2
    REPRODUCTION = 3
)

const (
    HEIGHT = 10
    WIDTH = 10
    LOWER_LIMIT = 0
)

func isOutBound(row int, col int) (bool) {
    return row < LOWER_LIMIT || col < LOWER_LIMIT || 
        row >= HEIGHT || col >= WIDTH
}

func numberOfAliveNeighbours(grid [][]bool, row int, col int) (int) {
    var retval = 0

    for i:=row-1; i <= row+1; i++ {
        for j:=col-1; j <= col+1; j++ {
            if !isOutBound(i, j) && 
                !(i==row && j==col) && 
                grid[i][j]{
                retval++
            }
        }
    } 

    return retval
}

func playLife(currentGrid [][]bool) ([][]bool) {
    newGrid:= make([][]bool, HEIGHT)
    for i := range newGrid {
        newGrid[i] = make([]bool, WIDTH)
    }

    for i:=0; i < HEIGHT; i++ {
        for j:=0; j < WIDTH; j++ {
            var neighbourAlive = numberOfAliveNeighbours(currentGrid, i, j)   

            if (neighbourAlive >= MIN_NEXT_GENERATION && 
                neighbourAlive <= MAX_NEXT_GENERATION && currentGrid[i][j]) {
                newGrid[i][j] = true
            } else if (neighbourAlive == REPRODUCTION && !currentGrid[i][j]) {
                newGrid[i][j] = true
            } 

        }
    }
    return newGrid
}

func displayGrid(grid [][]bool) {
    for i:=0; i < HEIGHT; i++ {
        for j:=0; j < WIDTH; j++ {
            printCell(grid[i][j])
        }
        fmt.Println("")
    }
    fmt.Println("====")
}

func printCell(cell bool) {
    if cell {
        fmt.Print(LIVE_CHAR)
    } else {
        fmt.Print(DEAD_CHAR)
    }
    fmt.Print(EMPTY_SPACE)
}

// Hardcoded test example
func populateGrid(grid [][]bool) {
    // TOAD Formation
    grid[3][3] = ALIVE
    grid[4][3] = ALIVE
    grid[5][3] = ALIVE
    grid[4][4] = ALIVE
    grid[5][4] = ALIVE
    grid[6][4] = ALIVE  
}

// Check if the grid still contain alive cells that
// can be used for reproduction
func isGameOver(grid [][]bool) (bool) {
    for i:=0; i < HEIGHT; i++ {
        for j:=0; j < WIDTH; j++ {
            if grid[i][j] {
                return false
            }
        }
    }

    return true
}

func initialiseGrid() ([][]bool) {
    grid:= make([][]bool, HEIGHT)
    for i := range grid {
        grid[i] = make([]bool, WIDTH)
    }

    return grid
}

func main() {
    grid := initialiseGrid()
    populateGrid(grid)

    for ;; {
        displayGrid(grid)
        copy(grid, playLife(grid))
        time.Sleep(500 * time.Millisecond)
        if isGameOver(grid) {
            displayGrid(grid)
            break
        }
    }
}
