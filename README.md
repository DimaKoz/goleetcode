# goleetcode

This repository contains some of my solutions to problems that were provided by leetcode.com. 

## easy lvl

- [1. Two Sum](https://leetcode.com/problems/two-sum/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1_two_sum.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1_two_sum_test.go)) <details> <summary>Description</summary>
  Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

  You may assume that each input would have exactly one solution, and you may not use the same element twice.

  You can return the answer in any order.
  
- [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/21_merge_two_sorted_lists.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/21_merge_two_sorted_lists_test.go)) <details> <summary>Description</summary>
  You are given the heads of two sorted linked lists `list1` and `list2`.

  Merge the two lists in a one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

  Return the head of the merged linked list.

- [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/104_maximum_depth_of_binary_tree.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/104_maximum_depth_of_binary_tree_test.go)) <details> <summary>Description</summary>
  Given the `root` of a binary tree, return its maximum depth.

  A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

 - [125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/125_valid_palindrome.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/125_valid_palindrome_test.go)) <details> <summary>Description</summary>
  A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

    Given a string `s`, return `true` if it is a **palindrome**, or `false` otherwise.

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

- [771. Jewels and Stones](https://leetcode.com/problems/jewels-and-stones/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/771_jewels_and_stones.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/771_jewels_and_stones_test.go)) <details> <summary>Description</summary>
  You're given strings `jewels` representing the types of stones that are jewels, and `stones` representing the stones you have. Each character in `stones` is a type of stone you have. You want to know how many of the stones you have are also jewels.

  Letters are case sensitive, so `"a"` is considered a different type of stone from `"A"`.

- [1108. Defanging an IP Address](https://leetcode.com/problems/defanging-an-ip-address/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1108_defanging_an_ip_address.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1108_defanging_an_ip_address_test.go)) <details> <summary>Description</summary>
  Given a valid (IPv4) IP address, return a defanged version of that IP address.

  A defanged IP address replaces every period `"."` with `"[.]"`.

- [1221. Split a String in Balanced Strings](https://leetcode.com/problems/split-a-string-in-balanced-strings/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1221_split_a_string_in_balanced_strings.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1221_split_a_string_in_balanced_strings_test.go)) <details> <summary>Description</summary>
  Balanced strings are those that have an equal quantity of `'L'` and `'R'` characters.

  Given a balanced string `s`, split it in the maximum amount of balanced strings.

  Return the maximum amount of split balanced strings.
  
- [1232. Check If It Is a Straight Line](https://leetcode.com/problems/check-if-it-is-a-straight-line/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1232_check_if_it_is_a_straight_line_test.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1232_check_if_it_is_a_straight_line_test.go)) <details> <summary>Description</summary>
You are given an array `coordinates`, `coordinates[i] = [x, y]`, where `[x, y]` represents the coordinate of a point. Check if these points make a straight line in the XY plane.

- [1290. Convert Binary Number in a Linked List to Integer](https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1290_convert_binary_number_in_a_linked_list_to_integer.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1290_convert_binary_number_in_a_linked_list_to_integer_test.go)) <details> <summary>Description</summary>
Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.

  Return the decimal value of the number in the linked list.  

- [1313. Decompress Run-Length Encoded List](https://leetcode.com/problems/decompress-run-length-encoded-list/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1313_decompress_run_length_encoded_list.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1313_decompress_run_length_encoded_list_test.go)) <details> <summary>Description</summary>
  We are given a list nums of integers representing a list compressed with run-length encoding.

  Consider each adjacent pair of elements `[freq, val] = [nums[2*i], nums[2*i+1]]` (with `i >= 0`).  For each such pair, there are `freq` elements with value `val` concatenated in a sublist. Concatenate all the sublists from left to right to generate the decompressed list.

  Return the decompressed list.
  
- [1342. Number of Steps to Reduce a Number to Zero](https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1342_number_of_steps_to_reduce_a_number_to_zero.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1342_number_of_steps_to_reduce_a_number_to_zero_test.go)) <details> <summary>Description</summary>
  Given an integer `num`, return *the number of steps to reduce it to zero*.

  In one step, if the current number is even, you have to divide it by `2`, otherwise, you have to subtract `1` from it.

- [1356. Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1356_sort_integers_by_the_number_of_1_bits.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1356_sort_integers_by_the_number_of_1_bits_test.go)) <details> <summary>Description</summary>
You are given an integer array `arr`. Sort the integers in the array in ascending order by the number of 1's in their binary representation and in case of two or more integers have the same number of 1's you have to sort them in ascending order.
Return the array after sorting it.

- [1389. Create Target Array in the Given Order](https://leetcode.com/problems/create-target-array-in-the-given-order/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1389_create_target_array_in_the_given_order.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1389_create_target_array_in_the_given_order_test.go)) <details> <summary>Description</summary>
  Given two arrays of integers `nums` and `index`. Your task is to create target array under the following rules:

  -Initially target array is empty.
  
  -From left to right read `nums[i]` and `index[i]`, insert at index `index[i]` the value `nums[i]` in target array.
  
  -Repeat the previous step until there are no elements to read in `nums` and `index`.
    
  -Return the target array.

  It is guaranteed that the insertion operations will be valid.

- [1431. Kids With the Greatest Number of Candies](https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1431_kids_with_the_greatest_number_of_candies.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1431_kids_with_the_greatest_number_of_candies_test.go)) <details> <summary>Description</summary>
  There are `n` kids with candies. You are given an integer array `candies`, where each `candies[i]` represents the number of candies the `i-th` kid has, and an integer `extraCandies`, denoting the number of extra candies that you have.

  Return a boolean array `result` of length `n`, where `result[i]` is `true` if, after giving the `i-th` kid all the `extraCandies`, they will have the greatest number of candies among all the kids, or `false` otherwise.

  Note that **multiple** kids can have the **greatest** number of candies.   
  
- [1470. Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1470_shuffle_the_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1470_shuffle_the_array_test.go)) <details> <summary>Description</summary>
  Given the array `nums` consisting of `2n` elements in the form `[x1,x2,...,xn,y1,y2,...,yn]`.

  Return the array in the form `[x1,y1,x2,y2,...,xn,yn]`.
  
- [1486. XOR Operation in an Array](https://leetcode.com/problems/xor-operation-in-an-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1486_xor_operation_in_an_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1486_xor_operation_in_an_array_test.go)) <details> <summary>Description</summary>
  You are given an integer `n` and an integer `start`.

  Define an array nums where `nums[i] = start + 2 * i` (0-indexed) and `n == nums.length`.

  Return the bitwise XOR of all elements of `nums`.
  
- [1512. Number of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1512_number_of_good_pairs.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1512_number_of_good_pairs_test.go)) <details> <summary>Description</summary>
  Given an array of integers `nums`, return the number of **good pairs**.

  A pair `(i, j)` is called good if `nums[i] == nums[j]` and `i < j`.

- [1528. Shuffle String](https://leetcode.com/problems/shuffle-string/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1528_shuffle_string.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1528_shuffle_string_test.go)) <details> <summary>Description</summary>
  You are given a string `s` and an integer array `indices` of the same length. The string `s` will be shuffled such that the character at the `i-th` position moves to `indices[i]` in the shuffled string.

  Return the *shuffled string*.

- [1572. Matrix Diagonal Sum](https://leetcode.com/problems/matrix-diagonal-sum/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1572_matrix_diagonal_sum.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1572_matrix_diagonal_sum_test.go)) <details> <summary>Description</summary>
  Given a square matrix `mat`, return the sum of the matrix diagonals.

  Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.
   
- [1588. Sum of All Odd Length Subarrays](https://leetcode.com/problems/sum-of-all-odd-length-subarrays/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1588_sum_of_all_odd_length_subarrays.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1588_sum_of_all_odd_length_subarrays_test.go)) <details> <summary>Description</summary>
Given an array of positive integers `arr`, calculate the sum of all possible odd-length subarrays.

  A subarray is a contiguous subsequence of the array.

  Return the sum of all odd-length subarrays of `arr`.

- [1662. Check If Two String Arrays are Equivalent](https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1662_check_if_two_string_arrays_are_equivalent.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1662_check_if_two_string_arrays_are_equivalent_test.go)) <details> <summary>Description</summary>
  Given two string arrays `word1` and `word2`, return `true` if the two arrays represent the same string, and `false` otherwise.

  A string is represented by an array if the array elements concatenated in order forms the string.

- [1684. Count the Number of Consistent Strings](https://leetcode.com/problems/count-the-number-of-consistent-strings/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1684_count_the_number_of_consistent_strings.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1684_count_the_number_of_consistent_strings_test.go)) <details> <summary>Description</summary>
  You are given a string `allowed` consisting of distinct characters and an array of strings `words`. A string is consistent if all characters in the string appear in the string `allowed`.

  Return the number of consistent strings in the array`words`.
  
- [1688. Count of Matches in Tournament](https://leetcode.com/problems/count-of-matches-in-tournament/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1688_count_of_matches_in_tournament.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1688_count_of_matches_in_tournament_test.go)) <details> <summary>Description</summary>
  You are given an integer `n`, the number of teams in a tournament that has strange rules:

  If the current number of teams is even, each team gets paired with another team. A total of `n / 2` matches are played, and `n / 2` teams advance to the next round.
  
  If the current number of teams is odd, one team randomly advances in the tournament, and the rest gets paired. A total of `(n - 1) / 2` matches are played, and `(n - 1) / 2 + 1` teams advance to the next round.
  
  Return the number of matches played in the tournament until a winner is decided.
  
- [1720. Decode XORed Array](https://leetcode.com/problems/decode-xored-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1720_decode_xored_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1720_decode_xored_array_test.go)) <details> <summary>Description</summary>
  There is a hidden integer array `arr` that consists of `n` non-negative integers.

  It was encoded into another integer array `encoded` of length `n - 1`, such that `encoded[i] = arr[i] XOR arr[i + 1]`. For example, if `arr = [1,0,2,1]`, then `encoded = [1,2,3]`.

  You are given the `encoded` array. You are also given an integer `first`, that is the first element of `arr`, i.e. `arr[0]`.

  Return the original array `arr`. It can be proved that the answer exists and is unique.
  
- [1773. Count Items Matching a Rule](https://leetcode.com/problems/count-items-matching-a-rule/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1773_count_items_matching_a_rule.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1773_count_items_matching_a_rule_test.go)) <details> <summary>Description</summary>
  You are given an array `items`, where each `items[i] = [type-i, color-i, name-i]` describes the type, color, and name of the `i-th` item. You are also given a rule represented by two strings, `ruleKey` and `ruleValue`.

  The `i-th` item is said to match the rule if one of the following is true:

  `ruleKey == "type"` and `ruleValue == type-i`.
  
  `ruleKey == "color"` and `ruleValue == color-i`.
  
  `ruleKey == "name"` and `ruleValue == name-i`.
  
  Return the number of items that match the given rule.

- [1790. Check if One String Swap Can Make Strings Equal](https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1790_check_if_one_string_swap_can_make_strings_equal.go), [tests](https://github.com/DimaKoz/goleetcode/blob/330cf42a69deaffbb444bbcf8d3fa150461d9005/easy/1790_check_if_one_string_swap_can_make_strings_equal_test.go#L5)) <details> <summary>Description</summary>
You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.&nbsp;
&nbsp;
Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.&nbsp;

- [1859. Sorting the Sentence](https://leetcode.com/problems/sorting-the-sentence/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1859_sorting_the_sentence.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1859_sorting_the_sentence_test.go)) <details> <summary>Description</summary>
  A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each word consists of lowercase and uppercase English letters.

  A sentence can be shuffled by appending the 1-indexed word position to each word then rearranging the words in the sentence.

  For example, the sentence `"This is a sentence"` can be shuffled as `"sentence4 a3 is2 This1"` or `"is2 sentence4 This1 a3"`.
  Given a shuffled sentence `s` containing no more than `9` words, reconstruct and return the original sentence.

- [1920. Build Array from Permutation](https://leetcode.com/problems/build-array-from-permutation/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1920_build_array_from_permutation.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1920_build_array_from_permutation_test.go)) <details> <summary>Description</summary>
  Given a zero-based permutation nums (0-indexed), build an array `ans` of the same length where `ans[i] = nums[nums[i]]` for each `0 <= i < nums.length` and return it.

  A zero-based permutation nums is an array of distinct integers from `0` to `nums.length - 1` (inclusive).

- [1929. Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1929_concatenation_of_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1929_concatenation_of_array_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums` of length `n`, you want to create an array `ans` of length `2n` where `ans[i] == nums[i]` and `ans[i + n] == nums[i]` for `0 <= i < n` (0-indexed).

  Specifically, `ans` is the concatenation of two `nums` arrays.

  Return the array `ans`.
  
- [2006. Count Number of Pairs With Absolute Difference K](https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2006_count_number_of_pairs_with_absolute_difference_k.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2006_count_number_of_pairs_with_absolute_difference_k_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums` and an integer `k`, return the number of pairs `(i, j)` where `i < j` such that `|nums[i] - nums[j]| == k`.

  The value of |x| is defined as:

  `x` if `x >= 0`.
  `-x` if `x < 0`.

- [2011. Final Value of Variable After Performing Operations](https://leetcode.com/problems/final-value-of-variable-after-performing-operations/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2011_final_value_of_variable_after_performing_operations.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2011_final_value_of_variable_after_performing_operations_test.go)) <details> <summary>Description</summary>
  There is a programming language with only four operations and one variable `X`:

    - `++X` and `X++` increments the value of the variable `X` by `1`.
  
    - `--X` and `X--` decrements the value of the variable `X` by `1`.
  
  Initially, the value of `X` is `0`.

  Given an array of strings operations containing a list of operations, return the final value of `X` after performing all the operations.

- [2037. Minimum Number of Moves to Seat Everyone](https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2037_minimum_number_of_moves_to_seat_everyone.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2037_minimum_number_of_moves_to_seat_everyone_test.go)) <details> <summary>Description</summary>
  There are `n` seats and `n` students in a room. You are given an array `seats` of length `n`, where `seats[i]` is the position of the `i-th` seat. You are also given the array `students` of length `n`, where `students[j]` is the position of the `j-th` student.

  You may perform the following move any number of times:

  Increase or decrease the position of the `i-th` student by `1` (i.e., moving the `i-th` student from position `x` to `x + 1` or `x - 1`)
  Return the minimum number of moves required to move each student to a seat such that no two students are in the same seat.

  Note that there may be multiple seats or students in the same position at the beginning.
  
- [2114. Maximum Number of Words Found in Sentences](https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2114_maximum_number_of_words_found_in_sentences.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2114_maximum_number_of_words_found_in_sentences_test.go)) <details> <summary>Description</summary>
  A sentence is a list of words that are separated by a single space with no leading or trailing spaces.

  You are given an array of strings `sentences`, where each `sentences[i]` represents a single sentence.

  Return the maximum number of words that appear in a single sentence.
  
- [2194. Cells in a Range on an Excel Sheet](https://leetcode.com/problems/cells-in-a-range-on-an-excel-sheet/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2194_cells_in_a_range_on_an_excel_sheet.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2194_cells_in_a_range_on_an_excel_sheet_test.go)) <details> <summary>Description</summary>
  A cell `(r, c)` of an excel sheet is represented as a string `"<col><row>"` where:

  -`<col>` denotes the column number `c` of the cell. It is represented by alphabetical letters.
  
    For example, the `1st` column is denoted by `'A'`, the `2nd` by `'B'`, the `3rd` by `'C'`, and so on.
  
  -`<row>` is the row number `r` of the cell. The `r-th` row is represented by the integer `r`.
  
  You are given a string `s` in the format `"<col1><row1>:<col2><row2>"`, where `<col1>` represents the column `c1`, `<row1>` represents the row `r1`, `<col2>` represents the column c2, and <row2> represents the row `r2`, such that `r1 <= r2` and `c1 <= c2`.

  Return the list of cells `(x, y)` such that `r1 <= x <= r2` and `c1 <= y <= c2`. The cells should be represented as strings in the format mentioned above and be sorted in non-decreasing order first by columns and then by rows.
  
- [2235. Add Two Integers](https://leetcode.com/problems/add-two-integers/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2235_add_two_integers.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2235_add_two_integers_test.go)) <details> <summary>Description</summary>
  Given two integers `num1` and `num2`, return the **sum** of the two integers.
  
- [2236. Root Equals Sum of Children](https://leetcode.com/problems/root-equals-sum-of-children/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2236_root_equals_sum_of_children.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2236_root_equals_sum_of_children_test.go)) <details> <summary>Description</summary>
  You are given the `root` of a binary tree that consists of exactly 3 nodes: the root, its left child, and its right child.

  Return `true` if the value of the root is equal to the sum of the values of its two children, or `false` otherwise.

  
## medium lvl
  
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

- [1476. Subrectangle Queries](https://leetcode.com/problems/subrectangle-queries/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1476_subrectangle_queries.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1476_subrectangle_queries_test.go)) <details> <summary>Description</summary>
  Implement the class `SubrectangleQueries` which receives a `rows x cols` rectangle as a matrix of integers in the constructor and supports two methods:

    1. `updateSubrectangle(int row1, int col1, int row2, int col2, int newValue)`
       
    Updates all values with `newValue` in the subrectangle whose upper left coordinate is `(row1,col1)` and bottom right coordinate is `(row2,col2)`.
    
    2. `getValue(int row, int col)`
       
    Returns the current value of the coordinate `(row,col)` from the rectangle.
  
- [1769. Minimum Number of Operations to Move All Balls to Each Box](https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/1769_minimum_number_of_operations_to_move_all_balls_to_each_box.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/1769_minimum_number_of_operations_to_move_all_balls_to_each_box_test.go)) <details> <summary>Description</summary>
  You have n boxes. You are given a binary string `boxes` of length `n`, where `boxes[i]` is `'0'` if the `i-th` box is empty, and `'1'` if it contains one ball.

  In one operation, you can move one ball from a box to an adjacent box. Box `i` is adjacent to box `j` if `abs(i - j) == 1`. Note that after doing so, there may be more than one ball in some boxes.

  Return an array `answer` of size `n`, where `answer[i]` is the minimum number of operations needed to move all the balls to the `i-th` box.

  Each `answer[i]` is calculated considering the initial state of the boxes.
  
- [2149. Rearrange Array Elements by Sign](https://leetcode.com/problems/rearrange-array-elements-by-sign/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/2149_rearrange_array_elements_by_sign.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/2149_rearrange_array_elements_by_sign_test.go)) <details> <summary>Description</summary>
  You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.

  You should rearrange the elements of nums such that the modified array follows the given conditions:

  Every consecutive pair of integers have opposite signs.
  For all integers with the same sign, the order in which they were present in nums is preserved.
  The rearranged array begins with a positive integer.
  Return the modified array after rearranging the elements to satisfy the aforementioned conditions.

- [2161. Partition Array According to Given Pivot](https://leetcode.com/problems/partition-array-according-to-given-pivot/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/2161_partition_array_according_to_given_pivot.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/2161_partition_array_according_to_given_pivot_test.go)) <details> <summary>Description</summary>
  You are given a 0-indexed integer array nums and an integer pivot. Rearrange nums such that the following conditions are satisfied:

  - Every element less than `pivot` appears before every element greater than `pivot`.
  - Every element equal to `pivot` appears in between the elements less than and greater than `pivot`.
  The relative order of the elements less than `pivot` and the elements greater than `pivot` is maintained.
  More formally, consider every `pi`, `pj` where `pi` is the new position of the `i-th` element and `pj` is the new position of the `j-th` element. For elements less than pivot, if `i < j` and `nums[i] < pivot` and `nums[j] < pivot`, then `pi < pj`. Similarly for elements greater than pivot, if `i < j` and `nums[i] > pivot` and `nums[j] > pivot`, then `pi < pj`.
  Return `nums` after the rearrangement.
  
## hard lvl
  
- [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/hard/4_median_of_two_sorted_arrays.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/hard/4_median_of_two_sorted_arrays_test.go)) <details> <summary>Description</summary>
  Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

  The overall run time complexity should be `O(log (m+n))`.
  
  

