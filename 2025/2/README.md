We can represent an ID range using a struct. The struct can have the following fields:
- start (`str`): The number (stringified) where the range starts
- end (`str`): The number (stringified) where the range ends

_Note:_ Storing them as strings makes sense because it's then easier for us to compare the 2 parts of the string (discussed in the Brute Force solution)

### Part 1
We need to find the sum of all the invalid IDs. This sum can be stored in a variable. Also, An invalid ID has to be _even_ digits long, so when evaluating a range, we can avoid checking any number which doesn't satisfy this criteria.

#### Brute Force
- Iterate over all the stringified numbers in a range. For each number:
    - If `{Length of number} % 2 != 0`, skip it
    - Otherwise, split the number in the middle and check if both parts are equal

### Part 2