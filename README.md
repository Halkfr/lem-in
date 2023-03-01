# Lem-in

## Description
This project simulates the behavior of ants in an ant farm. The goal is to find the best path for a given number of ants to travel from the starting point to the destination, while avoiding traffic jams.

## Approach
To accomplish this, we use a combination of depth-first search (DFS) and a heuristic algorithm to find the best path. Here's how the algorithm works:

1. DFS is used to find all possible paths from the starting point to the destination.
2. To avoid redundant or intersecting paths, we compare all the paths found in step 1 and save only the shortest non-repetitive combinations.
3. The best combinations of paths for different numbers of ants are found by calculating the total length of each combination and selecting the ones with the shortest total length.
4. Next, ants are assigned to the tunnels based on the selected combination of paths, with no more than one ant per tunnel at a time.
5. The ants then run along their assigned tunnels to the destination, following the path laid out for them.
6. As a bonus, a visualization of the ant farm and the ants' movements is available by default.


  

## Usage

### How to run

To run on local machine:

- Download the repository

- Run with a command `go run . [examples/filename.txt]`

- Open [http://localhost:8080/](http://localhost:8080/) in browser to look at visualization

To test [01 Edu cases](https://github.com/01-edu/public/tree/master/subjects/lem-in/audit):

- Run `./test.sh`

- Open [http://localhost:8080/](http://localhost:8080/) in browser to look at test case visualization

- Stop server with `CONTROL-C`

- Press any key to go to the next case

### Input file structure

```
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

Example of valid input file:

```
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1
```


## Implementation details

- DFS

- Handle input errors specifically

- Visualize with [github.com/fzipp/canvas](https://github.com/fzipp/canvas)


## Authors
  
Orel Margarita @maggieeagle

Litvintsev Anton @Antosha7


## Issue solving

- If `./test.sh` returns `Permission denied`, add executable permission to script:

 ```
 chmod +x test.sh
 ```

