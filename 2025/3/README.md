https://adventofcode.com/2025/day/3

### Part 1
Find the largest possible joltage each bank can produce

### Part 2

### Parsing the input
Each line contains a list of numbers, representing a battery bank. We can load this into an array of integers. Since there are multiple lines, the input will be represented as `[][]int`

We will call this input `BatteryBanks`. It will be an array of `BatteryBank` where the `BatteryBank` struct contains methods that make it easier to work with a battery bank

#### The `BatteryBank` struct
Fields:
- joltages (`[]int`): Contains the joltage rating for each battery in the bank

Methods:
- `createInverseIndexMap()`: Maps each unique joltage rating to an array of index positions it is found at in `joltages`

### Approach

#### Part 1
For each `batteryBank` in `BatteryBanks`, we need to find 2 values:
- The largest value of the first battery
- The largest value of the second battery
