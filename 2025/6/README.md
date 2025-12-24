https://adventofcode.com/2025/day/6

This one's a bit different, as the main "challenge" here seems to be parsing the input, while the actual problem (part 1, at least) looks easy

### Part 1
Add the solutions of all individual math problems and return the result

### Part 2


### Parsing the input
Splitting each line by space gets us the stringified numbers, except the last line where we get the operators. We can create a struct which stores both these numbers and the operators as separate fields. We can then create a method on the struct to solve a particular problem. We can then use another struct which is basically a collection of these individual problems and implement a `solve` method which can aggregate these results together and give us the final output.

Parsing the input for the second problem is a lot tougher. For one, we gotta read it RTL now. That's not the biggest headache though. The smaller number might start from either the rightmost or the leftmost position!

(The next day)
I slept over it and I think I have a solution now. Since a number is created from top to bottom, we have to process the input vertically. We also know that any line that vertically only consists of whitespaces is a boundary between 2 problems. Here is how we can go about it:
- We already have a way to process the operators (just call `strings.Fields()` on the last line), so we'll skip that line
- We will initialize a new `mathProblem` with an empty `numbers` field and a populated `operator`.
- This new problem will be our _current_ problem
- Process each line vertically (from top to bottom) and store everything we see as a string.
- Remove whitespaces from the string and convert it to an integer
- Add this integer to the `numbers` array of the current problem
- When we see a blank line, we will add the _current_ problem to the set of problems and reset _current_ problem to a new `mathProblem`

#### The `mathProblem` struct
Fields:
- `numbers` (`[]int`): The list of numbers the problem consists of
- `operator` (`string`): The operator to be applied to these problems

Methods:
- `solve() int`: Solves the problem according to the numbers and operator provided. Returns the solution

#### The `mathProblems` struct
Fields:
- `problems` (`[]mathProblem`): The collection of all math problems

Methods:
- `solve() int`: Returns the final solution by adding the solutions to all individual problems

### Approach

#### Part 1
- Parse the input and create a `problems (mathProblems)` variable
- Call `problems.solve()` and print the output

#### Part 2