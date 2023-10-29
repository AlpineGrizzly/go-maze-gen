/**
 * maze-gen.go
 * 
 * Maze generation program
 * 
 * Created Oct 29th, 2023
 * Author R3V
 */

package main

import ( 
    "fmt"
    "math/rand"
    "time"
)

/* Coordinate structure */
type coord struct { 
    x int
    y int
}

/** Queue Code */
/** Thanks https://www.geeksforgeeks.org/queue-in-go-language/ */
func enqueue(queue []coord, element coord) []coord { 
    queue = append(queue, element) // Simply append to enqueue. 
    fmt.Println("Enqueued:", element) 
    return queue 
} 
  
func dequeue(queue []coord) (coord, []coord) { 
    element := queue[0] // The first element is the one to be dequeued. 
    if len(queue) == 1 { 
        var tmp = []coord{} 
        return element, tmp  
    }   
    return element, queue[1:] // Slice off the element once it is dequeued. 
}

/**
 * is_visited
 * 
 * Returns true if a cell is visited, false otherwise
 */
func is_visited(visited []coord, cell coord) bool { 
    for i := 0; i < len(visited); i++ { 
        if visited[i] == cell { 
            return true
        }
    }
    return false
}


/** Gets wall between two cells */
func get_wall(host coord, neighbor coord) coord { 
    wall := coord{x: -1, y: -1}
    
    /* Check left and right directions */
    if host.x > neighbor.x { 
        wall = coord{host.x-1,host.y}
    } else if host.x < neighbor.x { 
        wall = coord{host.x+1,host.y}
    }

    /* Check up and down directions */
    if host.y > neighbor.y { 
        wall = coord{host.y,host.y-1}
    } else if host.y < neighbor.y { 
        wall = coord{host.y,host.y+1}
    }
   
    fmt.Println("Wall between ", host, " and ", neighbor, " is ", wall)
    return wall /* Return the coordinates of the host if no walls */
} 
/**
 * get_unvisited_neighbor
 * 
 * @param visited Array of cells that have already been visited
 * @param cell Cell that we are retrieiving and checking neighbors of
 * @param dim_x X dimensions of maze
 * @param dim_y Y diminsion of maze
 *
 * @return an unvisited neighbor of the cell passed in and the wall between the two
 */
func get_unvisited_neighbor(visited []coord, cell coord, dim_x int, dim_y int) (coord, coord) {
    neighbors := []coord{}
    chosen := coord{x: -1, y: -1}
    wall := coord{x: -1, y: -1}
    wall_size := 2

    /* Check all 4 neighbors to see if they are valid neighbors */
    /** (1,0) (-1,0) (0,1) (0,-1) */
    // Manually checking cause im too lazy to think 
    check_cell := coord{cell.x+wall_size, cell.y} 
    if check_cell.x < dim_x && !is_visited(visited, check_cell) { 
        neighbors = append(neighbors, check_cell)
    }
    
    check_cell = coord{cell.x-wall_size, cell.y} 
    if check_cell.x >= 0 && !is_visited(visited, check_cell) {  
        neighbors = append(neighbors, check_cell )
    }

    check_cell = coord{cell.x, cell.y+wall_size} 
    if check_cell.y < dim_y && !is_visited(visited, check_cell) {  
        neighbors = append(neighbors, check_cell)
    }
    
    check_cell = coord{cell.x, cell.y-wall_size} 
    if check_cell.y >= 0 && !is_visited(visited, check_cell) {  
        neighbors = append(neighbors, check_cell)
    }

    fmt.Println("Neighbors: ", neighbors, len(neighbors))

    /* Randomly choose one of the unvisited neighbors */
    if len(neighbors) > 0 { 
        chosen = neighbors[rand.Intn(len(neighbors))] /* Neighbor */
        wall = get_wall(cell, chosen) /* Wall value */
        fmt.Println("Neighbor: ", chosen, "Wall: ", wall) 
    }
    return chosen, wall
}

/**
 * generate_maze
 * 
 * Given a pointer to an array, generate map and assign it to it
 *
 * @param maze Pointer to an array to generate maze
 */
func generate_maze(maze [][]string, dim_x int, dim_y int) { 
    /** Initialize pick at coordinates for beginning and end */ 
    rand.Seed(time.Now().UnixNano())
    // TODO add a check to make sure they aren't the same value
    start := coord{x: 0, y: rand.Intn(dim_y)}
    end := coord{x: dim_x - 1, y: rand.Intn(dim_y)}
    
    // DEBUG print of coordinates
    fmt.Print("Start: ", start) 
    fmt.Print("End: ", end)
    fmt.Println()

    /** Generate maze with start and finish points */ 
    for j := 0; j < dim_y; j++ { 
        for i := 0; i < dim_x; i++ { 
            if start.x == i && start.y == j { 
                maze[i][j] = "@"
                continue
            } else if end.x == i && end.y == j { 
                maze[i][j] = "&"
                continue
            }
            maze[i][j] = "#"
        }
    }
    
    /** Use depth first search to carve path into maze */
    /** https://www.algosome.com/articles/maze-generation-depth-first.html */
    visited_cells := []coord{} /* Array of cells visited */
    queue := []coord{}         /* Queue to hold cells to visit */
    
    /* Randomly select a node */
    rand_cell :=  start //coord{x: rand.Intn(dim_x), y: rand.Intn(dim_y )}
    neigh_cell := coord{}
    wall_cell := coord{}

    maze[rand_cell.x][rand_cell.y] = " " // Clear the cell

    for k := 0; k < 15; k++ { 
        /* Push the node onto the queue */
        queue = enqueue(queue, rand_cell)

        /* Mark the cell as visited */
        maze[rand_cell.x][rand_cell.y] = " "
        visited_cells = append(visited_cells, rand_cell)

        /* Randomly select an adjacent cell of the node that has not been visited */
        // NOTE may have to check for a null value in the future 
        neigh_cell, wall_cell = get_unvisited_neighbor(visited_cells, rand_cell, dim_x, dim_y)
        
        /** If all neighbors have been visited */
        if rand_cell.x == -1 && rand_cell.y == -1 { 
            break
        }
        // Continue to pop items off the queue until a node is encountered with at least one non visited neighbor"
        
        /* Break the wall between the node and the neighbor */
        maze[wall_cell.x][wall_cell.y] = " "

        /* Assign the random cell value to the neighbor */
        rand_cell = neigh_cell
    }

    fmt.Println("Visited cells: ", visited_cells)
    fmt.Println("Queue: ", queue)
    fmt.Println()

    
}

func main() {
    /* Initialize array that will represent the maze */
    /** Declare variables */
    var maze [][]string
    var dim_x, dim_y int = 20, 10

    /** Initialize memory for array */
    maze = make([][]string, dim_x)
    for i := range maze { 
        maze[i] = make([]string, dim_y)
    }
   
    /** Generate the maze */
    generate_maze(maze, dim_x, dim_y)

    /** Simple print of the maze */
    for j := 0; j < dim_y; j++ { 
        for i := 0; i < dim_x; i++ { 
            fmt.Print(maze[i][j], " ")
        }
        fmt.Println()
    }
    fmt.Println()

    /** Main loop that will allow the player to traverse the maze */
}