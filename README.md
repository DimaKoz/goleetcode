# goleetcode

This repository contains some of my solutions to problems that were provided by leetcode.com. 

## easy lvl

- [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/496_next_greater_element_1.go), [tests](https://github.com/DimaKoz/goleetcode/blob/0caf9689320aca77618e24cd1ded4a6f75c7eb18/easy/496_next_greater_element_1_test.go#L5)) <details> <summary>Description</summary>
The **next greater element** of some element `x` in an array is the **first greater** element that is **to the right** of `x` in the same array.

  You are given two **distinct 0-indexed** integer arrays `nums1` and `nums2`, where `nums1` is a subset of `nums2`.
For each `0 <= i < nums1.length`, find the index `j` such that `nums1[i] == nums2[j]` and determine the next greater element of `nums2[j]` in `nums2`. If there is no next greater element, then the answer for this query is `-1`.

  Return an array `ans` of length `nums1.length` such that `ans[i]` is the **next greater element** as described above.
  
- [589. N-ary Tree Preorder Traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/589_n-ary_tree_preorder_traversal.go), [tests](https://github.com/DimaKoz/goleetcode/blob/330cf42a69deaffbb444bbcf8d3fa150461d9005/easy/589_n-ary_tree_preorder_traversal_test.go#L5)) <details> <summary>Description</summary>
Given the root of an n-ary tree, return the preorder traversal of its nodes' values.&nbsp;
Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value. 

- [1790. Check if One String Swap Can Make Strings Equal](https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1790_check_if_one_string_swap_can_make_strings_equal.go), [tests](https://github.com/DimaKoz/goleetcode/blob/330cf42a69deaffbb444bbcf8d3fa150461d9005/easy/1790_check_if_one_string_swap_can_make_strings_equal_test.go#L5)) <details> <summary>Description</summary>
You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.&nbsp;
&nbsp;
Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.&nbsp;
