# maze-gen
Go Program that generates mazes

Cool resource for future go projects https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs

## BUGS
## TODO                                     
- [ ] Implement other maze generation algorithms that can be selected with a flag
- [ ] Possible add an enemy to the maze the attempts to catch the player, pac man style
- [ ] Could add timer for solving the maze proportional to its size
- [ ] Pac man like token to collect in the maze while being chased by enemies

## Usage
```sh
Usage of ./maze-gen:
  -s	Show maze being generated
  -x int
    	X dimension of maze to be generated (Minimum x dim is 5) (default 25)
  -y int
    	Y dimension of maze to be generated (Minimum y dim is 5) (default 25)

# Generate a default maze 
./maze-gen
# # # # # # # # # # # # # # # # # # # # # # # # # 
# >     #       #       #                   #   # 
#   #   #   #   #   #   # # # # # # # # #   #   # 
#   #   #   #   #   #                   #     X # 
#   # # #   #   #   # # # # # # #   # # # # #   # 
#       #   #   #   #   #       #               # 
#   #   # # #   #   #   #   #   # # # # # # # # # 
#   #   #       #       #   #   #               # 
#   #   #   # # # # #   #   #   #   # # # # # # # 
#   #           #       #   #   #       #       # 
#   # # # # #   #   # # #   #   # # #   #   #   # 
#   #       #   #   #       #           #   #   # 
#   #   #   #   #   #   # # # # # # # # #   #   # 
#       #   #       #   #       #   #       #   # 
#   # # #   # # # # #   # # #   #   #   # # #   # 
#       #   #           #       #   #   #       # 
# # #   #   #   #   # # #   # # #   #   #   #   # 
#   #   #       #           #           #   #   # 
#   #   # # # # # # #   # # #   # # # # #   #   # 
#   #   #   #       #   #       #       #   #   # 
#   #   #   #   #   #   #   # # #   #   #   # # # 
#   #       #   #   #               #   #       # 
#   # # # # #   # # # # # # # # # # # # # # #   # 
#                                               # 
# # # # # # # # # # # # # # # # # # # # # # # # #
 
# Generate a maze with custom dimensions
./maze-gen -x 51 -y 51

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # 
# >                     #                                                   #           #           X 
#   # # #   #   # # #   #   # # # # #   # # # # # # # # # # # # #   #   # # #   # # # # #   # # #   # 
#       #   #   #       #   #       #               #       #       #           #       #       #   # 
# # #   #   #   # # # # #   #   # # # # # # #   # # #   #   #   # # # # # # # # #   #   # # #   #   # 
#   #   #   #           #   #               #           #   #           #           #           #   # 
#   #   # # # # #   #   #   #   # # # # #   # # # # # # #   # # # # #   #   # # # # # # # # # # #   # 
#   #       #       #       #   #       #       #           #       #       #   #           #       # 
#   # # #   #   # # # # # # #   #   #   # # #   # # # # # # #   #   # # # # #   #   # # #   # # # # # 
#   #       #           #       #   #       #                   #               #       #           # 
#   #   # # #   #   #   #   # # #   # # # # # # # # # # #   # # #   # # # # # # # # #   # # # # #   # 
#   #       #   #   #           #       #       #           #       #               #   #           # 
#   # # #   #   # # # # # # #   # # #   #   #   #   #   #   #   # # #   # # # # #   #   # # # # # # # 
#       #   #       #       #       #       #       #   #   #                   #   #       #       # 
# # #   #   # # #   #   #   # # #   # # # # #   #   # # #   #   # # # # # # #   #   # # #   # # #   # 
#   #       #   #       #   #       #           #   #       #   #       #       #       #   #       # 
#   #   #   #   # # # # #   #   # # # # # # # # #   #   # # #   # # #   # # # # #   #   #   #   # # # 
#       #       #           #   #               #   #       #           #       #   #               # 
# # # # # # # # #   #   #   #   #   #   # # #   #   # # #   # # # # # # #   #   #   # # # # # # # # # 
#           #       #   #   #       #       #   #       #   #               #   #                   # 
#   # # #   #   #   # # #   # # # # #   #   #   # # #   #   #   # # # # # # #   # # # # # # #   #   # 
#   #       #   #       #   #           #   #       #   #       #           #           #   #   #   # 
#   #   # # #   #   #   # # #   # # # # # # # # #   #   # # # # #   # # #   # # #   #   #   #   #   # 
#   #       #   #   #       #               #       #           #   #   #           #       #   #   # 
#   # # # # #   #   # # #   # # # # # # #   # # # # # # # # # # #   #   # # # # #   # # # # #   #   # 
#               #       #               #           #               #               #       #   #   # 
#   # # # # # # # # #   # # # # # # #   # # # # #   #   # # # # #   # # # # # # # # #   #   #   #   # 
#   #       #           #           #   #           #   #       #               #   #   #       #   # 
#   #   # # #   # # # # #   # # #   #   #   # # # # #   #   # # # # # # # # #   #   #   # # # # # # # 
#   #   #       #               #   #   #           #   #                       #   #   #           # 
# # #   #   # # # # # # # # #   #   #   # # # # #   #   # # # # # # # # # # # # #   #   #   # # #   # 
#       #   #       #           #       #           #   #               #           #           #   # 
#   # # #   #   #   #   # # #   # # # # #   # # # # #   #   # # #   #   #   # # # # # # # # # # #   # 
#           #   #   #       #   #       #   #       #   #       #   #       #               #       # 
# # # # # # #   #   # # # # #   #   #   #   #   #   #   # # #   #   # # # # #   #   # # #   #   # # # 
#       #       #               #   #   #   #   #       #       #   #       #   #   #       #       # 
#   #   #   # # #   # # # # #   #   #   #   #   # # #   #   # # #   #   #   #   #   # # # # # # # # # 
#   #   #   #       #           #   #       #       #   #   #   #       #       #   #               # 
#   #   #   #   # # #   # # # # #   #   # # # # #   # # #   #   # # # # # # # # #   #   # # # # #   # 
#   #   #   #       #   #       #   #   #       #       #   #   #               #       #       #   # 
#   # # #   # # #   #   #   # # #   #   #   # # # # #   #   #   #   # # # # #   # # # # #   #   #   # 
#           #       #   #       #   #   #               #       #       #       #           #   #   # 
# # # # #   #   # # #   # # #   #   #   # # # # # # # # # # #   #   #   # # # # #   # # # # # # #   # 
#           #       #           #   #   #                   #   #   #       #       #               # 
#   # # # # # # #   # # # # # # #   #   #   # # # # # # #   #   #   # # #   #   # # #   # # # # # # # 
#               #           #       #   #       #       #       #   #       #           #           # 
# # # # # # # # # # #   #   # # #   # # #   #   #   #   # # # # #   #   # # # # # # # # #   #   #   # 
#       #           #   #   #       #       #   #   #       #       #               #       #   #   # 
#   #   # # #   #   # # #   #   # # #   # # #   #   #   # # #   # # # # # # #   # # #   # # #   #   # 
#   #           #           #               #       #                       #           #       #   # 
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
```

Using the arrow keys, you are able to traverse the maze and attempt to solve it
