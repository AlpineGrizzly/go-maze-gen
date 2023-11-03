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
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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
		wall = coord{host.x - 1, host.y}
	} else if host.x < neighbor.x {
		wall = coord{host.x + 1, host.y}
	}

	/* Check up and down directions */
	if host.y > neighbor.y {
		wall = coord{host.x, host.y - 1}
	} else if host.y < neighbor.y {
		wall = coord{host.x, host.y + 1}
	}

	return wall /* Return the coordinates of the host if no walls */
}

func clear_screen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/** Check if the cell is on the boundary of the maze */
func is_edge_box(cell coord, dim_x int, dim_y int) bool {
	if cell.x == 0 || cell.x == dim_x-1 || cell.y == 0 || cell.y == dim_y-1 {
		return true
	}
	return false
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
	check_cell := coord{cell.x + wall_size, cell.y}
	if check_cell.x < dim_x && !is_visited(visited, check_cell) && !is_edge_box(check_cell, dim_x, dim_y) {
		neighbors = append(neighbors, check_cell)
	}

	check_cell = coord{cell.x - wall_size, cell.y}
	if check_cell.x >= 0 && !is_visited(visited, check_cell) && !is_edge_box(check_cell, dim_x, dim_y) {
		neighbors = append(neighbors, check_cell)
	}

	check_cell = coord{cell.x, cell.y + wall_size}
	if check_cell.y < dim_y && !is_visited(visited, check_cell) && !is_edge_box(check_cell, dim_x, dim_y) {
		neighbors = append(neighbors, check_cell)
	}

	check_cell = coord{cell.x, cell.y - wall_size}
	if check_cell.y >= 0 && !is_visited(visited, check_cell) && !is_edge_box(check_cell, dim_x, dim_y) {
		neighbors = append(neighbors, check_cell)
	}

	/* Randomly choose one of the unvisited neighbors */
	if len(neighbors) > 0 {
		chosen = neighbors[rand.Intn(len(neighbors))] /* Neighbor */
		wall = get_wall(cell, chosen)                 /* The wall between the neighbor and the host */
	}
	return chosen, wall
}

/**
 * set_cell
 *
 * Will set a coordinate in the maze to a string value
 *
 * @param cell Coordinate in the maze to be set
 * @param val Value to set coordinate to
 */
func set_cell(maze [][]string, cell coord, val string) {
	maze[cell.x][cell.y] = val
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

	start := coord{x: 1, y: 1}
	var end coord

	/** Generate maze with start and finish points */
	for j := 0; j < dim_y; j++ {
		for i := 0; i < dim_x; i++ {
			maze[i][j] = "#"
		}
	}

	/** Use depth first search to carve path into maze */
	/** https://www.algosome.com/articles/maze-generation-depth-first.html */
	visited_cells := []coord{} /* Array of cells visited */
	queue := []coord{}         /* Queue to hold cells to visit */

	/* Randomly select a node */
	rand_cell := start //coord{x: rand.Intn(dim_x), y: rand.Intn(dim_y)}
	var neigh_cell coord
	var wall_cell coord
	var temp_cell coord

	for {
		clear_screen()
		/* Push the node onto the queue */
		queue = enqueue(queue, rand_cell)

		/* Mark the cell as visited */
		visited_cells = append(visited_cells, rand_cell)

		/* Randomly select an adjacent cell of the node that has not been visited */
		neigh_cell, wall_cell = get_unvisited_neighbor(visited_cells, rand_cell, dim_x, dim_y)

		/** If all neighbors have been visited */
		if neigh_cell.x == -1 && neigh_cell.y == -1 {
			for len(queue) > 0 {
				// Continue to pop items off the queue until a node is encountered with at least one non visited neighbor
				rand_cell, queue = dequeue(queue)

				// Check if rand_cell has neighbors
				temp_cell, _ = get_unvisited_neighbor(visited_cells, rand_cell, dim_x, dim_y)

				if temp_cell.x != -1 && temp_cell.y != -1 {
					break
				}
			}

			// Check if we have anything in our queue
			if len(queue) == 0 {
				break
			}
			end = rand_cell
			continue
		}

		/* Break the wall between the node and the neighbor */
		maze[rand_cell.x][rand_cell.y] = " "
		maze[neigh_cell.x][neigh_cell.y] = " "
		maze[wall_cell.x][wall_cell.y] = " "

		/* Assign the random cell value to the neighbor */
		rand_cell = neigh_cell
		if *g_show_gen {
			print_maze(maze, dim_x, dim_y)
			time.Sleep(5 * time.Millisecond)
		}
	}
	/* Set the start and end cells  */
	fmt.Println("Start is ", start, " End is ", end)
	set_cell(maze, start, ">")
	set_cell(maze, end, "X")

}

/** Print the maze */
func print_maze(maze [][]string, dim_x int, dim_y int) {
	/** Simple print of the maze */
	for j := 0; j < dim_y; j++ {
		for i := 0; i < dim_x; i++ {
			fmt.Print(maze[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

/* Global Variables */
var g_show_gen *bool

func main() {
	var dim_x, dim_y int = 25, 25
	var maze [][]string

	/* Get size of array from user args */
	flag.IntVar(&dim_x, "x", dim_x, "X dimension of maze to be generated")
	flag.IntVar(&dim_y, "y", dim_y, "Y dimension of maze to be generated")
	g_show_gen = flag.Bool("s", false, "Show maze being generated")

	flag.Parse()

	/** Initialize memory for array */
	maze = make([][]string, dim_x)
	for i := range maze {
		maze[i] = make([]string, dim_y)
	}

	/** Generate the maze */
	generate_maze(maze, dim_x, dim_y)

	print_maze(maze, dim_x, dim_y)

	/** Main loop that will allow the player to traverse the maze */
}
