https://adventofcode.com/2025/day/5

### Part 1
Count the number of fresh ingredient IDs

### Part 2
Count the number of IDs that are considered to be fresh

### Parsing the input
The input consists of 2 parts. Until we see a blank line, we see ID ranges. These are inclusive number ranges where each number in the range represents an ingredient ID that is _fresh_. I can store these ranges in some struct that has a min and max value.

After the blank line, we have a list of ingredient IDs (one per line). I can store this in an array

#### The `ingredientIDRange` struct
Fields:
- `min` (`int`): The minimum ID value in an ID range
- `max` (`int`): The maximum ID value in an ID range

Methods:
- `isIDFresh(id int) bool`: Checks if the given ID is fresh

### The `ingredientIDRanges` struct
Fields:
- `ranges` (`[]ingredientIDRange`): An array of id ranges

Methods:
- `IsIDFresh(id int) bool`: Checks if the given ID is fresh. Calls `isIDFresh` for each id range and returns `true` if any function call returned `true`, `false` otherwise
- `getDistinct() ingredientIDRanges`: Merges overlapping ID ranges and returns a new `ingredientIDRanges` with only distinct ID ranges
- `ProcessFreshIDRanges() int`: Processes all distinct ID ranges and returns the number of IDs considered to be fresh

### Approach

#### Part 1
- Create a `count` variable to store the number of fresh ids
- For each ingredient ID `id`, call `ingredientIDRanges`.`IsIDFresh(id)`
    - If the return value is `true`, increment `count`
- Return `count`

#### Part 2
- Get the initial set of fresh ID Ranges `ranges` from the input
- Call `ranges`.`ProcessFreshIDRanges` and return the count