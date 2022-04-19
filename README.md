# goleetcode

This repository contains some of my solutions to problems that were provided by leetcode.com. 

## easy lvl

- [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/104_maximum_depth_of_binary_tree.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/104_maximum_depth_of_binary_tree_test.go)) <details> <summary>Description</summary>
  Given the `root` of a binary tree, return its maximum depth.

  A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

- [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/217_contains_duplicate.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/217_contains_duplicate_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.


- [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/283_move_zeroes.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/283_move_zeroes_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums`, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

  Note that you must do this in-place without making a copy of the array.


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
