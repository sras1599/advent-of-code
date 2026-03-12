https://adventofcode.com/2025/day/9

## Part 1
Find the area of the largest rectange that can be created using two red tiles on opposite corners

## Part 2
Find the largest area of any rectange that can be created using only red and green tiles

### Parsing the input
The input contains coordinates for all the red tiles. Each line contains 2 comma separated ints that contain the tile's `x` and `y` axis. We can create a struct for convenience (something like `redTile`) and return an array of those

### Approach

#### Part 1

Since we have the coordinates of 2 corners of the rectangle, we can calculate it's area:
```go
area := ((c2.x + 1) - c1.x) * ((c2.y + 1) - c2.y)
```

We can do something like this:
- Split the grid in the middle vertically
- Store a variable to track the biggest area we've seen so far (`largestArea`)
- For each tile that exists on the left side of the grid:
    - Pick a tile on the right side
    - Use these 2 tiles as corners and calculate their diagonal distance (`area`)
    - if `area` > `largestArea`, update `largestArea`, if not continue

#### Part 2

This one asks us to do the same thing, but we're working with a constraint. We first need to figure out the co-ordinates this loop covers

To fill up the co-ordinates that the loop covers, we can divide the process into 2 parts:
- Fill the spaces between only adjacent (vertically or horizontally) red tiles first
- Fill in the gaps between these spaces to get the total "coverage". We only need to get this coverage along one axis (either `x` or `y`)

We can store the coverage as a map where the key will be the index of the line and the value will be a range