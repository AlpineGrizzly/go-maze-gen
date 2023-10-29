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

func get_random_neighbor(cell coord, dim_x int, dim_y int) coord {
    neighbors := []coord{}

    /* Check all 4 neighbors to see if they are valid neighbors */
    /** (1,0) (-1,0) (0,1) (0,-1) */
    // Manually checking cause im too lazy to think 
    if cell.x+1 < dim_x { 
        neighbors = append(neighbors, coord{cell.x+1, cell.y})
    }
    if cell.x-1 >= 0 {  
        neighbors = append(neighbors, coord{cell.x-1, cell.y})
    }
    if cell.y+1 < dim_y {  
        neighbors = append(neighbors, coord{cell.x, cell.y+1})
    }
    if cell.y-1 >= 0 {  
        neighbors = append(neighbors, coord{cell.x, cell.y-1})
    }
    fmt.Println("Neighbors: ", neighbors, len(neighbors))

    /* Randomly choose one of the neighbors */
    chosen := neighbors[rand.Intn(len(neighbors))]
    fmt.Println("Chosen: ", chosen) 

    return chosen

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
    rand_cell :=  coord{x: rand.Intn(dim_x), y: rand.Intn(dim_y )}


    /* Push the node onto the queue */
    queue = enqueue(queue, rand_cell)

    /* Mark the cell as visited */
    visited_cells = append(visited_cells, rand_cell)

    /* Randomly select an adjacent cell of the node that has not been visited */
    rand_cell = get_random_neighbor(rand_cell, dim_x, dim_y)

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