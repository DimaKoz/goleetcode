# goleetcode

This repository contains some of my solutions to problems that were provided by leetcode.com. 

## easy lvl

- [1. Two Sum](https://leetcode.com/problems/two-sum/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1_two_sum.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1_two_sum_test.go)) <details> <summary>Description</summary>
  Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

  You may assume that each input would have exactly one solution, and you may not use the same element twice.

  You can return the answer in any order.
  
- [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/2_add_two_numbers.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/2_add_two_numbers_test.go)) <details> <summary>Description</summary>
  You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the `sum` as a linked list.

  You may assume the two numbers do not contain any leading zero, except the number 0 itself.
  
  `Input: l1 = [2,4,3], l2 = [5,6,4]`
  
  `Output: [7,0,8]`
  
  `Explanation: 342 + 465 = 807`

 - [8. String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/8_string_to_integer_atoi.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/8_string_to_integer_atoi_test.go)) <details> <summary>Description</summary>
  Implement the `myAtoi(string s)` function, which converts a string to a 32-bit signed integer (similar to C/C++'s `atoi` function).
  
    The algorithm for `myAtoi(string s)` is as follows:

    - Read in and ignore any leading whitespace.
    - Check if the next character (if not already at the end of the string) is `'-'` or `'+'`. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
    - Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
    - Convert these digits into an integer (i.e. `"123" -> 123`, `"0032" -> 32`). If no digits were read, then the integer is `0`. Change the sign as necessary (from step 2).
    - If the integer is out of the 32-bit signed integer range `[-2^31, 2^31 - 1]`, then clamp the integer so that it remains in the range. Specifically, integers less than `-2^31` should be clamped to `-2^31`, and integers greater than `2^31 - 1` should be clamped to `2^31 - 1`.
    - Return the integer as the final result.
  
    **Note:**

    Only the space character `' '` is considered a whitespace character.
    Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.
  
- [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/21_merge_two_sorted_lists.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/21_merge_two_sorted_lists_test.go)) <details> <summary>Description</summary>
  You are given the heads of two sorted linked lists `list1` and `list2`.

  Merge the two lists in a one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

  Return the head of the merged linked list.
  
- [29. Divide Two Integers](https://leetcode.com/problems/divide-two-integers/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/29_divide_two_integers.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/29_divide_two_integers_test.go)) <details> <summary>Description</summary>
  Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

  The integer division should truncate toward zero, which means losing its fractional part. For example, `8.345` would be truncated to `8`, and `-2.7335` would be truncated to `-2`.

  Return the quotient after dividing dividend by divisor.

  Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: `[−2^31, 2^31 − 1]`. For this problem, if the quotient is strictly greater than `2^31 - 1`, then return `2^31 - 1`, and if the quotient is strictly less than `-2^31`, then return `-2^31`.

- [38. Count and Say](https://leetcode.com/problems/count-and-say/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/38_count_and_say.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/38_count_and_say_test.go)) <details> <summary>Description</summary>
  The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

  `countAndSay(1) = "1"`
  `countAndSay(n)` is the way you would "say" the digit string from `countAndSay(n-1)`, which is then converted into a different digit string.
  To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.

  For example, the saying and conversion for digit string `"3322251"`: !["3322251":](https://assets.leetcode.com/uploads/2020/10/23/countandsay.jpg)


  Given a positive integer `n`, return the `n-th` term of the count-and-say sequence.

- [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/104_maximum_depth_of_binary_tree.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/104_maximum_depth_of_binary_tree_test.go)) <details> <summary>Description</summary>
  Given the `root` of a binary tree, return its maximum depth.

  A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

- [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/217_contains_duplicate.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/217_contains_duplicate_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.


- [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/283_move_zeroes.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/283_move_zeroes_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums`, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

  Note that you must do this in-place without making a copy of the array.

- [344. Reverse String](https://leetcode.com/problems/reverse-string/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/344_reverse_string.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/344_reverse_string_test.go)) <details> <summary>Description</summary>
  Write a function that reverses a string. The input string is given as an array of characters `s`.

  You must do this by modifying the input array in-place with **O(1)** extra memory.

- [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/496_next_greater_element_1.go), [tests](https://github.com/DimaKoz/goleetcode/blob/0caf9689320aca77618e24cd1ded4a6f75c7eb18/easy/496_next_greater_element_1_test.go#L5)) <details> <summary>Description</summary>
The **next greater element** of some element `x` in an array is the **first greater** element that is **to the right** of `x` in the same array.

  You are given two **distinct 0-indexed** integer arrays `nums1` and `nums2`, where `nums1` is a subset of `nums2`.
For each `0 <= i < nums1.length`, find the index `j` such that `nums1[i] == nums2[j]` and determine the next greater element of `nums2[j]` in `nums2`. If there is no next greater element, then the answer for this query is `-1`.

  Return an array `ans` of length `nums1.length` such that `ans[i]` is the **next greater element** as described above.

- [566. Reshape the Matrix](https://leetcode.com/problems/reshape-the-matrix/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/566_reshape_the_matrix.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/566_reshape_the_matrix_test.go)) <details> <summary>Description</summary>
  In MATLAB, there is a handy function called **reshape** which can reshape an `m x n` matrix into a new one with a different size `r x c` keeping its original data.

  You are given an `m x n` matrix `mat` and two integers `r` and `c` representing the number of rows and the number of columns of the wanted reshaped matrix.

  The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.

  If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

- [589. N-ary Tree Preorder Traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/589_n-ary_tree_preorder_traversal.go), [tests](https://github.com/DimaKoz/goleetcode/blob/330cf42a69deaffbb444bbcf8d3fa150461d9005/easy/589_n-ary_tree_preorder_traversal_test.go#L5)) <details> <summary>Description</summary>
Given the root of an n-ary tree, return the preorder traversal of its nodes' values.&nbsp;
Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value. 
  


- [709. To Lower Case](https://leetcode.com/problems/to-lower-case/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/709_to_lower_case.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/709_to_lower_case_test.go)) <details> <summary>Description</summary>
Given a string `s`, return the string after replacing every uppercase letter with the same lowercase letter.
  
- [1232. Check If It Is a Straight Line](https://leetcode.com/problems/check-if-it-is-a-straight-line/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1232_check_if_it_is_a_straight_line_test.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1232_check_if_it_is_a_straight_line_test.go)) <details> <summary>Description</summary>
You are given an array `coordinates`, `coordinates[i] = [x, y]`, where `[x, y]` represents the coordinate of a point. Check if these points make a straight line in the XY plane.

- [1290. Convert Binary Number in a Linked List to Integer](https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1290_convert_binary_number_in_a_linked_list_to_integer.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1290_convert_binary_number_in_a_linked_list_to_integer_test.go)) <details> <summary>Description</summary>
Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.

  Return the decimal value of the number in the linked list.  

 - [1356. Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1356_sort_integers_by_the_number_of_1_bits.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1356_sort_integers_by_the_number_of_1_bits_test.go)) <details> <summary>Description</summary>
You are given an integer array `arr`. Sort the integers in the array in ascending order by the number of 1's in their binary representation and in case of two or more integers have the same number of 1's you have to sort them in ascending order.
Return the array after sorting it.

- [1572. Matrix Diagonal Sum](https://leetcode.com/problems/matrix-diagonal-sum/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1572_matrix_diagonal_sum.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1572_matrix_diagonal_sum_test.go)) <details> <summary>Description</summary>
  Given a square matrix `mat`, return the sum of the matrix diagonals.

  Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.
  
  
- [1588. Sum of All Odd Length Subarrays](https://leetcode.com/problems/sum-of-all-odd-length-subarrays/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1588_sum_of_all_odd_length_subarrays.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1588_sum_of_all_odd_length_subarrays_test.go)) <details> <summary>Description</summary>
Given an array of positive integers `arr`, calculate the sum of all possible odd-length subarrays.

  A subarray is a contiguous subsequence of the array.

  Return the sum of all odd-length subarrays of `arr`.
  
- [1790. Check if One String Swap Can Make Strings Equal](https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1790_check_if_one_string_swap_can_make_strings_equal.go), [tests](https://github.com/DimaKoz/goleetcode/blob/330cf42a69deaffbb444bbcf8d3fa150461d9005/easy/1790_check_if_one_string_swap_can_make_strings_equal_test.go#L5)) <details> <summary>Description</summary>
You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.&nbsp;
&nbsp;
Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.&nbsp;
