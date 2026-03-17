# Day 10

https://adventofcode.com/2025/day/10

## Part 1

What are the fewest number of button presses that can correctly configure the indicator lights on all the machines?

## Part 2

TODO: Add notes for part 2

### Parsing the input

Each line of the input contains information about a single machine. For us it represents a single problem as well, so we can use [a struct](#the-factorymachine-struct) to represent this information

#### The `factoryMachine` struct
Fields:
- `indicatorLightDiagram` (`string`): The correct configuration of indicator lights we need to get to
- `wiringSchematics` (`[][]int`): A sequence of button wiring schematics, where a single schematic tells us which lights it turns on
- `joltageRequirements` (`[]int`): Joltage requirements. Not mentioned in part 1

Methods:
- `getFewestButtonPressesToCorrectConfig`: Gets the fewest button presses to configure the indicator lights in the desired configuration. It's working is explained in the [approach](#part-1-1) section

### Approch

#### Part 1
Brute force solution:

We just need to try all possible combinations. To achieve this, we can do the following...

For each machine:
- Keep a track of the number of presses it takes us to get to the desired config (`presses`)
- Look at each button in the wiring schematic
    - If the first press gets us there, return `1`
    - Evaluate the button's press with every other button in the schematic. To evaluate means to check which lights are now on after all the buttons in the current evaluation have been pressed
    - If the evaluation returns the same result as the desired configuration, return `presses`. Otherwise, store this evaluation in a list so that we can evaluate with another schematic later
    - Repeat steps 2 & 3 until we get an evaluation that exactly matches the indicator light diagram

We can then simply sum up the result from each machine in the final output