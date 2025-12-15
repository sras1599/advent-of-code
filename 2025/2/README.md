https://adventofcode.com/2025/day/2

### Part 1
We need to find the sum of all the invalid IDs. An invalid ID is one which is only made up of 2 similar sequences. This naturally means that the length of the ID needs to be an _even_ number

### Part 2
We need to find the sum of all the invalid IDs. An invalid ID is one which is only made up of 2 or more similar sequences

### Generic utilities
- `GetMultiples(n int)` - A function to return all multiples of the number (length of the ID), ignoring 1, and including the number itself
    - For Part 1, we can ignore all multiples other than 2
- `CanSplitIntoSimilarChunks(s string, n int)` - A function to check if a string can be split into similar chunks of length `n`
    - Returns `false` the moment it sees a chunk which is not the same as the first chunk of length `n`

### The IDRange struct
We can represent an ID range using a struct. The struct can have the following fields:
- `start` (`int`): The number where the range starts
- `end` (`int`): The number where the range ends
- `sumOfInvalidIDs` (`int`): The sum of all invalid IDs. 0 by default
- `multiplesMap` (`map[string][]int`): A mapping that contains multiples of numbers representing the stringified lengths of all numbers in the range

#### Methods
- `getUniqueLengths() []int` - Get the unique lengths of all numbers in the range
- `populateMultiplesMap()`: Method to populate the `multiplesMap` of the `IDRange`
    - Calls `getUniqueLengths()` and stores the result in a variable
    - For each number, calls `GetMultiples` and stores the result in the mappings
- `processNumber(num int, part int)` - Processes the number to check if it can be split into similar chunks of length 2 (or more, if part 2)
    - Converts the number to string
    - If `part == 1`
        - Calls `CanSplitIntoSimilarChunks(stringifiedNumber, 2)`
            - If the return value is `true`, appends the `num` to `sumOfInvalidIDs`
    - If `part == 2`
        - Gets the chunk lenghts it needs to process from `multiplesMap[len(strigifiedNum)]`
        - For each chunk length
            - Calls `CanSplitIntoSimilarChunks(stringifiedNumber, chunkLength)`
                - If the return value is `true`, appends the `num` to `sumOfInvalidIDs`. Also, break early since we don't need to process other chunk lengths
    - _Note:_ This can be a goroutine and called concurrently. We just need to make sure that we add to `sumOfInvalidIDs` concurrently
- `processRange(total int, part int)`: For each number `n` in the range:
    - Call `processNumber(n, part)`
    - After the for loop has finished, append the range's `sumOfInvalidIDs` to the `total`
    - _Note:_ This can be a goroutine and called concurrently.
- `GetsumOfInvalidIDs() int`: Returns the sum of invalid IDs


### Solutions

#### Part 1
- Maintain a counter `total` which contains the sum of all invalid IDs.
- Parse the input line to get a list of `IDRange`
- For each `IDRange`
    - Call `processRange(total, 1)`
- Print `total`

#### Part 2
- Maintain a counter `total` which contains the sum of all invalid IDs.
- Parse the input line to get a list of `IDRange`
- For each `IDRange`
    - Call `processRange(total, 2)`
- Print `total`