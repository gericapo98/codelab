# 1. TwoSum:
#    Given an array of integers nums and an integer target, return indices of the two numbers 
#    such that they add up to the target. You may assume that each input would have exactly one solution, 
#    and you may not use the same element twice. Return the answer in any order.

# 2. Valid Parentheses:
#    Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', 
#    determine if the input string is valid. An input string is valid if:
#    - Open brackets must be closed by the same type of brackets.
#    - Open brackets must be closed in the correct order.

# 3. Merge Intervals:
#    Giv en an array of intervals where intervals[i] = [start_i, end_i], merge all overlapping intervals, 
#    and return an array of the non-overlapping intervals that cover all the intervals in the input.

# 4. Top-K Frequent Elements:
#    Given an integer array nums and an integer k, return the k most frequent elements. 
#    You may return the answer in any order.

# 5. First Non-Repeating Character in a String:
#    Given a string s, find the first non-repeating character in it and return its index. 
#    If it does not exist, return -1.

from typing import List

def two_sum(nums: List[int], target: int) -> List[int]:
    """
    Given an array of integers nums and an integer target, 
    return indices of the two numbers 
    such that they add up to the target.
    """
    
    num_to_index = {}
    for index, num in enumerate(nums):
        complement = target - num
        if complement in num_to_index:
            return [num_to_index[complement], index]
        num_to_index[num] = index
    return []    
