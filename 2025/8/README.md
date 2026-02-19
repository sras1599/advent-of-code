https://adventofcode.com/2025/day/8

## Part 1
Connect the closest 1000 pairs of Junction boxes and return the product of the 3 largest circuits

## Part 2
Process all junction connections in order of distance until all boxes are merged into a single circuit. Return the product of the X coordinates of the last two junction boxes that merged together.

### Parsing the input
Each line of the input contains three comma-separated integers representing the 3D coordinates (`x`, `y`, `z`) of a junction box. Parse each line to extract these coordinates.

### Approach

#### The `junctionBox` struct
An abstraction representing a junction box in 3D space. Stores the `x`, `y`, and `z` coordinates.

#### The `junctionConnection` struct
Represents a pair of junction boxes and the Euclidean straight-line distance between them. Fields:
- `box1`: The first junction box
- `box2`: The second junction box
- `distance`: The Euclidean distance between the two boxes

#### Part 1

Here's the approach:
- Maintain a list of circuits. Each circuit is of type `[]junctionBox`
    - We need helper methods to manage these circuits, so we use a struct wrapper `junctiionCircuits`
- Parse the input to get all the boxes (`[]junctionBox`)
- Generate all possible junction connections between pairs of boxes
    - Sort this list by straight-line distance (closest first)
- Process the closest 1000 connections. For each connection, there are 4 possibilities:
    - **Both boxes are in no circuit**: Create a new circuit with both boxes
    - **Both boxes are in the same circuit**: Skip (already connected)
    - **Boxes are in different circuits**: Merge the two circuits
    - **One box is in a circuit, the other is not**: Add the box without a circuit to the circuit of the other
- After processing, sort the circuits by size in descending order
- Calculate the product of the sizes of the 3 largest circuits

#### Part 2

Here's the approach:
- Parse the input to get all the boxes
- Generate all possible junction connections and sort them by distance (same as Part 1)
- Process connections one by one, applying the same circuit management logic
- After each connection is processed, check if all boxes have been merged into a single circuit
    - This happens when the largest circuit contains all junction boxes
- When all boxes are finally merged, return the product of the X coordinates of the last two boxes that were connected