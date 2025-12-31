https://adventofcode.com/2025/day/6


### Part 1
Find the number of times a tachyon beam is split

### Part 2
Find the number of timelines, or the number of possible journeys the beam could've taken to reach the end of the manifold

### Parsing the input
Should be relatively simple for the way I'm trying to approach this. Just parse each line as a list of chars

There should also be no need of any structs here, just some state variables and we should be good to go

### Approach

#### Part 1
We need a few state variables:
- `splits (int)`: The number of times a tachyon beam has been split

Here's how we can go about it
- Process a line at any given time and have access to the next 2 lines
- For line 1, locate the index of "S" (`sIndex`) and update the next line's {sIndex}
- For every other line at index `lineIndex`:
    - if `lineIndex` is more than `numLines - 2`, break out of the loop
    - Note that the state for the line has already been updated, so the beams will already be in the right positions
    - For every char `c` at index `i`:
        - If we see an empty space or splitter, continue
        - If we see a beam at index:
            - if {nextLine[i]} is a splitter, place a beam at `lines[lineIndex+2][i-1]` and `lines[LineInex+2][i+1]`. Check for out of bounds here
            - Add 1 to `splits`

Return the value of `splits`

#### Part 2
After starting the beam, we need to process each timeline separately. We can count the number of timelines we go into using a counter variable. A splitter is basically a point where a new timeline is created. We can split the beam on either side of the splitter and then store the state of the manifold in separate variables. We can then process these states separately. After processing each state we can increment the counter and return the result.

_THE NEXT DAY_
I was able to figure out the answer, but that required processing the beam through all available timelines. While it was relatively quick to do this for the test input, the script ran for more than 20 minutes on the real input before I abandoned the approach. Need to figure out something else now

The first question I need to answer is: **Could this be done in one shot so that we iterate over the array just once?**

##### THE WEIGHTED BEAM!
We can change how we interpret the manifold. Each empty space can represent `0`, and a splitter can represent `-1`. The starting point can be `-2`. Here's the idea:
- When the beam starts from `-2`, the empty space below it becomes `1`. This is the _weight_ of the beam
- We split the beam normally, but if the index below it is receiving beams from both the left and the right, we increment the weight of the beam with the individual weights of the beams making it. These are all the possible timelines for that particular index in the array. At max, beams from 3 indexes (left, center, right) can contribute to the weight of the beam below it.
- To get the answer, we just add the weights of the beams we find in the last line