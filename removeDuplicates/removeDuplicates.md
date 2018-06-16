#  2. Remove Duplicates from Sorted Array

> Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

> Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.


给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成


### Example 1:

```go
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
```

### Example 2:

```go
Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
```

### Clarification:

> Confused why the returned value is an integer but your answer is an array?

> Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```go
// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

# solution

## 方法：双指针法

### 算法

数组完成排序后，我们可以放置两个指针 ii 和 jj，其中 ii 是慢指针，而 jj 是快指针。只要 nums[i] = nums[j]nums[i]=nums[j]，我们就增加 jj 以跳过重复项。

当我们遇到 nums[j] \neq nums[i]nums[j]≠nums[i] 时，跳过重复项的运行已经结束，因此我们必须把它（nums[j]nums[j]）的值复制到 nums[i + 1]nums[i+1]。然后递增 ii，接着我们将再次重复相同的过程，直到 jj 到达数组的末尾为止。


### 复杂度分析

- 时间复杂度：O(n)O(n)， 假设数组的长度是 nn，那么 ii 和 jj 分别最多遍历 nn 步。

- 空间复杂度：O(1)O(1)。