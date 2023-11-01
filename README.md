# maze-gen
Go Program that generates mazes

Cool resource for future go projects https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs

## BUGS
- [ ] Mazes are sometimes unsolvable, could be something funky with the algorithm.
## TODO                                     
- [ ] Allow for a player to be able to solve the mazes
- [ ] Implement other maze generation algorithms that can be selected with a flag
- [ ] Possible add an enemy to the maze the attempts to catch the player, pac man style

## Requirements

## Usage
```sh
# Generate a default maze of 100 by 50
./maze-gen

# Generate a maze with custom dimensions
./maze-gen -x 50 -y 50

          #       #       #   #           #           #           # # #           #               # 
  # # #   # # #   #   #   #   #   # # #   #   # # #   #   # # #   # # #   # # #   #   # # # # #   # 
  # # #           #   #   #   #   # # #   #   # # #       #       #       # # #   #       #   #   # 
  # # # # # # # # # # #   #   #   # # #   #   # # # # # # #   # # #   # # # # #   # # #   #   #   # 
  #   #                   #   #       #   #   #       #       # # #   #       #   # # #   #   #   # 
# #   # # # # # # # # # # # # # # #   #   # # #   #   #   # # # # #   # # #   #   # # #   #   #   # 
# #                           #       #           #       #           #       #   #       #       # 
# # # # # # # # # # # # # #   # # # # # # # # # # # # # # # # # # # # # # # # #   #   # # # # # # # 
      #                   #   #       #       #               #       #       #   #   #           # 
  # # # # # # # # # # #   #   # # #   #   #   # # # # # # #   #   #   # # # # #   #   # # # # #   # 
  #       #               #   #   #       #       #           #   #   #           #       # # #   # 
# #   #   # # # # # # # # #   #   # # # # # # #   #   # # # # # # # # #   # # # # # # #   # # #   # 
      #   #                   #   #   #       #   #       #           #       # # #       #       # 
  # # #   #   # # # # # # # # #   #   #   #   #   # # # # #   # # #   # # #   # # #   # # #   # # # 
  #   #       #   #   #           #   #   #   #   #       #   #   #   # # #       #   # # #       # 
  #   # # # # #   #   #   # # # # #   #   #   #   #   #   #   #   #   # # # # #   #   # # # # #   & 
      #       #       #   # # #       #   #       #   #       #   #           #       #           # 
# # # #   # # # # # # #   # # #   # # #   # # # # #   # # # # # # # # # # #   # # # # # # # # # # # 
          # # #       #           #       #       #   #                   #       #   #   #       # 
  # # # # # # #   # # # # # # # # #   # # #   #   #   #   # # # # # # # # # # #   #   #   # # #   # 
              #   #       #           #       #   #   #   # # #                   #   #           # 
# # # # # #   # # #   #   #   # # # # #   # # #   #   #   # # #   # # # # # # # # #   # # # # # # # 
          #           #   #   #       #       #   #   #           #           #   #       #       # 
# # # # # # # # # # # #   #   #   # # # # #   #   #   # # # # # # #   # # #   #   # # #   # # # # # 
      #       #           #   #   #       #   #   #   #       #       #       #       #           # 
  # # #   #   #   # # # # #   # # #   #   #   #   #   #   #   # # # # # # # # # # #   # # # # #   # 
  #       #       #       #   #       #       #   #   #   #               #       #       #       # 
  #   # # # # # # # # # # #   #   # # # # # # #   #   #   # # # # # # #   #   #   # # #   #   # # # 
  #   #           #       #       #       # # #   #       #       #       #   #       #   #   #   # 
# #   # # # # # # #   #   # # # # # # #   # # #   # # # # # # #   #   # # #   # # #   #   #   #   # 
      #           #   #               #   #       #   #           #       #   #       #   #       # 
  # # #   # # #   #   # # # # # # #   #   #   # # #   #   # # # # # # #   #   #   # # #   # # # # # 
          # # #       #       #   #       #   #       #   #       #       #   #   # # #   #   # # # 
# # # # # # # # # # # # # #   #   # # # # #   # # # # # # #   #   #   # # # # #   # # #   #   # # # 
              #           #       #       #       #       #   #   #   #       #       #   #       # 
  # # # # #   # # # # #   # # # # # # # # # # #   #   #   #   # # #   #   #   # # #   #   # # #   # 
  #       #   #           #           #       #       #   #   #       #   #           #       #   # 
  #   #   #   #   # # # # # # # # # # #   # # # # # # #   # # #   # # # # # # # # # # # # #   #   # 
  #   #       #           #               #   #           # # #   #       #       #       #   #   # 
  #   # # # # # # # # # # #   # # # # # # #   #   # # # # # # #   # # #   #   #   # # # # #   #   # 
  #           #               #   #       #   #       #           #       #   #       #       #   # 
  # # # # #   #   # # # # # # #   #   # # #   # # #   #   # # # # #   # # #   # # #   #   # # # # # 
      #   #   #       #           #   #   #   # # #       #           #       # # #   #           # 
# #   #   #   # # #   # # # # # # #   #   #   # # # # # # #   # # # # #   # # # # #   # # # # #   # 
  #       #   # # #           #       #   #       #           #   #       #       #           #   # 
  # # # # #   # # # # # # #   #   # # # # # # #   #   # # # # #   #   # # #   #   # # # # #   #   # 
          #               #   #   #       #   #   #   #           #   # # #   #   #       #   #   # 
# # # #   # # # # # # #   #   #   #   #   #   #   #   #   # # # # #   # # #   #   # # #   #   #   # 
      #               #       #       #       #       #   #       #           #           #       # 
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # 
```

