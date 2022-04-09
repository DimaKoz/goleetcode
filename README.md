# goleetcode

This repository contains some of my solutions to problems that were provided by leetcode.com. 

## easy lvl

- [589. N-ary Tree Preorder Traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/589_n-ary_tree_preorder_traversal.go), [tests](https://github.com/DimaKoz/goleetcode/blob/330cf42a69deaffbb444bbcf8d3fa150461d9005/easy/589_n-ary_tree_preorder_traversal_test.go#L5)) <details> <summary>Description</summary>
Given the root of an n-ary tree, return the preorder traversal of its nodes' values.&nbsp;
Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value. 

- [1779. Find Nearest Point That Has the Same X or Y Coordinate](https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1790_check_if_one_string_swap_can_make_strings_equal.go), [tests](https://github.com/DimaKoz/goleetcode/blob/330cf42a69deaffbb444bbcf8d3fa150461d9005/easy/1790_check_if_one_string_swap_can_make_strings_equal_test.go#L5)) <details> <summary>Description</summary>
You are given two integers, x and y, which represent your current location on a Cartesian grid: (x, y). You are also given an array points where each points[i] = [ai, bi] represents that a point exists at (ai, bi). A point is valid if it shares the same x-coordinate or the same y-coordinate as your location.&nbsp;
Return the index (0-indexed) of the valid point with the smallest Manhattan distance from your current location. If there are multiple, return the valid point with the smallest index. If there are no valid points, return -1.&nbsp;
The Manhattan distance between two points (x1, y1) and (x2, y2) is abs(x1 - x2) + abs(y1 - y2).
