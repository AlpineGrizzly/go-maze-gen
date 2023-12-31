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
 * Given an array of visited elements and a cell, return a boolean on whether it is present in the array or not
 *
 * @param visited Coordinate array of visited cells 
 * @param cell Cell to be checked for presence in visited array
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

/**
 * get_wall
 * 
 * Get the coordinate value of a wall between two coordinates
 *
 * @param host Cell that is used as the base
 * @param neighbor Neighbor cell of host
 *
 * Return the coordinate of the value between the host and neighbor
 */
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

/**
 * clear_screen
 *
 * Clears the screen of text
 */
func clear_screen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/**
 * is_edge_box
 *
 * Given a cell, return a boolean if the cell is an edge box or not 
 */
func is_edge_box(cell coord) bool {
	if cell.x == 0 || cell.x == g_dim_x-1 || cell.y == 0 || cell.y == g_dim_y-1 {
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
func get_unvisited_neighbor(visited []coord, cell coord) (coord, coord) {
	neighbors := []coord{}
	chosen := coord{x: -1, y: -1}
	wall := coord{x: -1, y: -1}
	wall_size := 2

	/* Check all 4 neighbors to see if they are valid neighbors */
	/** (1,0) (-1,0) (0,1) (0,-1) */
	// Manually checking cause im too lazy to think
	check_cell := coord{cell.x + wall_size, cell.y}
	if check_cell.x < g_dim_x && !is_visited(visited, check_cell) && !is_edge_box(check_cell) {
		neighbors = append(neighbors, check_cell)
	}

	check_cell = coord{cell.x - wall_size, cell.y}
	if check_cell.x >= 0 && !is_visited(visited, check_cell) && !is_edge_box(check_cell) {
		neighbors = append(neighbors, check_cell)
	}

	check_cell = coord{cell.x, cell.y + wall_size}
	if check_cell.y < g_dim_y && !is_visited(visited, check_cell) && !is_edge_box(check_cell) {
		neighbors = append(neighbors, check_cell)
	}

	check_cell = coord{cell.x, cell.y - wall_size}
	if check_cell.y >= 0 && !is_visited(visited, check_cell) && !is_edge_box(check_cell) {
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
func generate_maze(maze [][]string) {
	/** Initialize pick at coordinates for beginning and end */
	rand.Seed(time.Now().UnixNano())

	start := coord{x: 1, y: 1}
	var end coord

	/** Generate maze with start and finish points */
	for j := 0; j < g_dim_y; j++ {
		for i := 0; i < g_dim_x; i++ {
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
		neigh_cell, wall_cell = get_unvisited_neighbor(visited_cells, rand_cell)

		/** If all neighbors have been visited */
		if neigh_cell.x == -1 && neigh_cell.y == -1 {
			for len(queue) > 0 {
				// Continue to pop items off the queue until a node is encountered with at least one non visited neighbor
				rand_cell, queue = dequeue(queue)
				end = rand_cell

				// Check if rand_cell has neighbors
				temp_cell, _ = get_unvisited_neighbor(visited_cells, rand_cell)

				if temp_cell.x != -1 && temp_cell.y != -1 {
					break
				}
			}

			// Check if we have anything in our queue
			if len(queue) == 0 {
				break
			}
			continue
		}

		/* Break the wall between the node and the neighbor */
		maze[rand_cell.x][rand_cell.y] = " "
		maze[neigh_cell.x][neigh_cell.y] = " "
		maze[wall_cell.x][wall_cell.y] = " "

		/* Assign the random cell value to the neighbor */
		rand_cell = neigh_cell
		if *g_show_gen {		
			fmt.Println(end)
			print_maze(maze)	
			time.Sleep(5 * time.Millisecond)
			
		}
	}
	/* Set the start and end cells  */
	g_player_loc = start
	g_goal_loc = end	
	fmt.Println("Start is ", start, " End is ", end)
	set_cell(maze, start, ">")
	set_cell(maze, end, "X")

}

/** Print the maze */
func print_maze(maze [][]string) {
	/** Simple print of the maze */
	for j := 0; j < g_dim_y; j++ {
		for i := 0; i < g_dim_x; i++ {
			fmt.Print(maze[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func is_wall(loc coord, maze [][]string) bool { 	
	maze_cell := maze[loc.x][loc.y]
	if maze_cell == "#" { 
		return true
	}
	return false

}

func validate_move(new_p_loc coord, maze [][]string) bool {
	/** Check if its an edge box */
	if is_edge_box(new_p_loc) || is_wall(new_p_loc, maze) { 
		return false
	}
	/** Check if its a wall */	
	return true 
}
// Initialize our global direction values
var g_k_up byte = 65
var g_k_down byte = 66
var g_k_left byte = 68
var g_k_right byte = 67
var character string 

func get_move(b []byte, maze [][]string) {	
	curr_p_loc := g_player_loc
	new_p_loc := g_player_loc

	os.Stdin.Read(b)
	if b[0] == 27 { 
		os.Stdin.Read(b)
		if b[0] == 91 {
			switch os.Stdin.Read(b); b[0] { 
				case g_k_up:
					new_p_loc = coord{x: curr_p_loc.x, y: curr_p_loc.y - 1}
                    character = "V"
				case g_k_down:
					new_p_loc = coord{x: curr_p_loc.x, y: curr_p_loc.y + 1}
                    character = "Λ"
				case g_k_left:
					new_p_loc = coord{x: curr_p_loc.x - 1, y: curr_p_loc.y}
                    character = ">"
				case g_k_right:
					new_p_loc = coord{x: curr_p_loc.x + 1, y: curr_p_loc.y}
                    character = "<"
			}
			if validate_move(new_p_loc, maze) {  
				// Update new player location
				maze[curr_p_loc.x][curr_p_loc.y] = " "
				maze[new_p_loc.x][new_p_loc.y] = character
				g_player_loc = new_p_loc		
			}
		}
	}

}

/* Global Variables */
var g_show_gen *bool			   /* Boolean flag for showing maze geneation if enabled */
var g_player_loc coord			   /* Starting location of the player in the maze */
var g_goal_loc coord			   /* Goal location of the player */
var g_dim_x, g_dim_y int = 25, 25  /* X and Y dimensions of Maze to be generated */
var MIN_SIZE int = 5 			   /* Minimum value of maze size */

func main() {
	var maze [][]string

	/* Get size of array from user args */
	flag.IntVar(&g_dim_x, "x", g_dim_x, "X dimension of maze to be generated (Minimum x dim is 5)")
	flag.IntVar(&g_dim_y, "y", g_dim_y, "Y dimension of maze to be generated (Minimum y dim is 5)")
	g_show_gen = flag.Bool("s", false, "Show maze being generated")
	flag.Parse()

	/** Check minimum requirement to generate a maze */
	if g_dim_x < MIN_SIZE || g_dim_y < MIN_SIZE {
		fmt.Println("Usage of ./maze-gen:")
		flag.PrintDefaults()
		os.Exit(0)	
	}	

	/** Initialize memory for array */
	maze = make([][]string, g_dim_x)
	for i := range maze {
		maze[i] = make([]string, g_dim_y)
	}

	/** Generate the maze */
	generate_maze(maze)

	/** Main loop that will allow the player to traverse the maze */
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)

	solved := false /** Bool for whether the maze has been solved by the player or not */
	for !solved { 
		print_maze(maze)
		get_move(b, maze) /* Get player move and update the maze */	
		if g_player_loc == g_goal_loc { 
			solved = true
			fmt.Println("You win!")
			break
		}
		clear_screen()
	}
	exec.Command("stty", "-F", "/dev/tty", "sane").Run() /** Reset stty */
}
