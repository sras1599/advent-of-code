We can create a struct to represent the dial. It can have the following properties:
- Start: The lowest value of the dial
- End: The highest point of the dial
- Range: The range of the dial, or the number of unique positions in the dial
- Current Position: The position of the dial after a rotation

### Rotation logic
- Since the dial has an upper and lower bound, we need to wrap the rotation amount in case it exceeds these bounds
- We can use the modulus operator for this. The dividend will be the rotation amount of the dial and the divisor will be the range of the dial. This will wrap the amount correctly and give us the updated position of the dial. An example:
    - 180 % 100 = 80

### Part 1
- We can maintain a counter that tracks how many times the dial position was 0 after a rotation.
- After each rotation, we can check the position of the dial and update the counter accordingly.

### Part 2
Now we also need to consider cases where the position of the dial was 0 *during* a rotation. Similar to part 1, we can track this number with a counter.

- We can handle all cases where the rotation amount is >=`{Range of Dial}` by just dividing the amount by `{Range of Dial}` using the integer division `/`. When dividing integers, go automatically strips any fractions from the results and rounds the value down to the nearest integer, which is exactly what we need. We can increment our counter by this _quotient_.
- We can now _normalize_ the amount using the modulus operator (`%`) and continue with the rotation. We can ignore any rotations where the normalized amount is `0`.
- For each _left_ rotation, we can increment the counter if `{Rotation Amount - Current Position of Dial}`  is greater than or equal to `0`. We know that we will either be at 0 or cross it during the rotation.
    - An edge case here is when the `{Current Position of Dial}` is `0`. In this case our formula will incorrectly end up incrementing the counter because the if check will always suffice. So we need to ignore this case. Our final if check can look something like this:
    ```go
    if (amount - curDialPos >= 0) && (curDialPos != 0) {
        // increment counter
    }
    ```
- For each _right_ rotation, we can increment the counter if `Rotation Amount + Current Position of Dial` is greater than 99