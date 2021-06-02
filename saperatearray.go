// Go Program to Put Positive and Negatives in a Separate Array
// Write a Go Program to Put Positive and Negative numbers in a Separate Array using for loop.
// First, we declared three arrays. The if condition (if posNegArr[i] >= 0) checks whether the array item is greater than or equal to zero within the loop.
// If True, we assigned that value to posArr (posArr[positiveCount] = posNegArr[i]) and incremented the positive count value.
// Otherwise, set the value to NegArr (NegArr[negativeCount] = posNegArr[i]) and increment the negative count value.
// Next, we created a function (printArray(posNegArr []int, size int)) that prints the positive and negative array items.