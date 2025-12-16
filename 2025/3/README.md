https://adventofcode.com/2025/day/3

### Part 1
Find the largest possible joltage by turning on exactly 2 batteries in each bank. Add the result from each bank

### Part 2
Find the largest possible joltage by turning on exactly 12 batteries in each bank. Add the result from each bank

### Parsing the input
Each line contains a list of numbers, representing a battery bank. We can load this into an array of integers. Since there are multiple lines, the input will be represented as `[][]int`

We will call this input `BatteryBanks`. It will be an array of `BatteryBank` where the `BatteryBank` struct contains methods that make it easier to work with a battery bank

#### The `BatteryBank` struct
Fields:
- `joltageRatings` (`[]int`): Contains the joltage rating for each battery in the bank

Methods:
- `findLargestPossibleJoltage()` - Look at Approach -> Part 1

### Approach

#### Part 1
For each `batteryBank` in `BatteryBanks`, we need to find 2 values:
- The largest value of the first battery. Let's call this `left`
- The largest value of the second battery. Let's call this `right`

- Get the index and value of the highest element in the array. Let's call this `highest`
- If elements to the right exist, find the maximum value in the subarray. Let's call this `target`
    - If yes, `left` becomes `highest` and `right` becomes `target`
- Find the maximum value in the subarray of elements left of `highest`. This maximum value will become `left` and `highest` becomes `right`

### Generic approach?
After seeing part 2, it looks like the same function can be used for solving both parts. We just need to find the largest number of `n` digits in the battery bank

23423423423427142337333337736382222222226812
222222222222226812

target digits: 5
maintain result string, we will convert this to an int at the end

- get first index of max number in the array -> `m`. target digits: 4
- if number of digits right of `m` are > target digits, only work with that subarray
    - if equal, that is the answer
- if less, that subarray will be in the final answer, so take it out and add it to the result str
- repeat?

9843344143323335344
