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

- [169. Majority Element](https://leetcode.com/problems/majority-element/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/169_majority_element.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/169_majority_element_test.go)) <details> <summary>Description</summary>
  Given an array `nums` of size `n`, return the majority element.

  The majority element is the element that appears more than `[n / 2]` times. You may assume that the majority element always exists in the array.

- [191. Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/191_number_of_1_bits.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/191_number_of_1_bits_test.go)) <details> <summary>Description</summary>
  Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the [Hamming weight](http://en.wikipedia.org/wiki/Hamming_weight)).

- [202. Happy Number](https://leetcode.com/problems/happy-number/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/202_happy_number.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/202_happy_number_test.go)) <details> <summary>Description</summary>
  Write an algorithm to determine if a number `n` is happy.

  A happy number is a number defined by the following process:

  Starting with any positive integer, replace the number by the sum of the squares of its digits.
  
  Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
  
  Those numbers for which this process ends in `1` are happy.
  
  Return `true` if `n` is a happy number, and `false` if not.

- [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/217_contains_duplicate.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/217_contains_duplicate_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.
  
- [232. Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/232_implement_queue_using_stacks.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/232_implement_queue_using_stacks_test.go)) <details> <summary>Description</summary>
  Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (`push`, `peek`, `pop`, and `empty`).

  Implement the MyQueue class:

  - `void push(int x)` Pushes element x to the back of the queue.
  
  - `int pop()` Removes the element from the front of the queue and returns it.

  - `int peek()` Returns the element at the front of the queue.

  - `boolean empty()` Returns true if the queue is empty, false otherwise.

  Notes:

  - You must use only standard operations of a stack, which means only `push to top`, `peek/pop from top`, `size`, and `is empty` operations are valid.

  - Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.

- [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/242_valid_anagram.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/242_valid_anagram_test.go)) <details> <summary>Description</summary>
  Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

  An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

- [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/283_move_zeroes.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/283_move_zeroes_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums`, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

  Note that you must do this in-place without making a copy of the array.

- [303. Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/303_range_sum_query_immutable.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/303_range_sum_query_immutable_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums`, handle multiple queries of the following type:

  1. Calculate the sum of the elements of `nums` between indices `left` and `right` inclusive where `left <= right`.

  Implement the NumArray class:

  - `NumArray(int[] nums)` Initializes the object with the integer array `nums`.

  - `int sumRange(int left, int right)` Returns the sum of the elements of `nums` between indices `left` and `right` inclusive (i.e. `nums[left] + nums[left + 1] + ... + nums[right]`).

- [344. Reverse String](https://leetcode.com/problems/reverse-string/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/344_reverse_string.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/344_reverse_string_test.go)) <details> <summary>Description</summary>
  Write a function that reverses a string. The input string is given as an array of characters `s`.

  You must do this by modifying the input array in-place with **O(1)** extra memory.
  
- [389. Find the Difference](https://leetcode.com/problems/find-the-difference/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/389_find_the_difference.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/389_find_the_difference_test.go)) <details> <summary>Description</summary>
  You are given two strings `s` and `t`.

  String `t` is generated by random shuffling string `s` and then add one more letter at a random position.

  Return the letter that was added to `t`.

- [404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/404_sum_of_left_leaves.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/404_sum_of_left_leaves_test.go)) <details> <summary>Description</summary>
  Given the `root` of a binary tree, return *the sum of all left leaves*.

  A **leaf** is a node with no children. A **left leaf** is a leaf that is the left child of another node.

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

- [728. Self Dividing Numbers](https://leetcode.com/problems/self-dividing-numbers/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/728_self_dividing_numbers.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/728_self_dividing_numbers_test.go)) <details> <summary>Description</summary>
  A self-dividing number is a number that is divisible by every digit it contains.

  - For example, `128` is a self-dividing number because `128 % 1 == 0`, `128 % 2 == 0`, and `128 % 8 == 0`.

  A self-dividing number is not allowed to contain the digit zero.

  Given two integers `left` and `right`, return a list of all the self-dividing numbers in the range `[left, right]`.

- [771. Jewels and Stones](https://leetcode.com/problems/jewels-and-stones/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/771_jewels_and_stones.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/771_jewels_and_stones_test.go)) <details> <summary>Description</summary>
  You're given strings `jewels` representing the types of stones that are jewels, and `stones` representing the stones you have. Each character in `stones` is a type of stone you have. You want to know how many of the stones you have are also jewels.

  Letters are case sensitive, so `"a"` is considered a different type of stone from `"A"`.

- [804. Unique Morse Code Words](https://leetcode.com/problems/unique-morse-code-words/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/804_unique_morse_code_words.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/804_unique_morse_code_words_test.go)) <details> <summary>Description</summary>
  International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows:

  - `'a'` maps to `".-"`,
  - `'b'` maps to `"-..."`,
  - `'c'` maps to `"-.-."`, and so on.
  
  For convenience, the full table for the `26` letters of the English alphabet is given below:

  `[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]`
  
  Given an array of strings `words` where each word can be written as a concatenation of the Morse code of each letter.

  - For example, `"cab"` can be written as `"-.-..--..."`, which is the concatenation of `"-.-."`, `".-"`, and `"-..."`. We will call such a concatenation the transformation of a word.

  Return the number of different transformations among all words we have.
  
- [832. Flipping an Image](https://leetcode.com/problems/flipping-an-image/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/832_flipping_an_image.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/832_flipping_an_image_test.go)) <details> <summary>Description</summary>
  Given an `n x n` binary matrix `image`, flip the image horizontally, then invert it, and return the resulting image.

  To flip an image horizontally means that each row of the image is reversed.

  - For example, flipping `[1,1,0]` horizontally results in `[0,1,1]`.

  To invert an image means that each `0` is replaced by `1`, and each `1` is replaced by `0`.

  - For example, inverting `[0,1,1]` results in `[1,0,0]`.

- [876. Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/876_middle_of_the_linked_list.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/876_middle_of_the_linked_list_test.go)) <details> <summary>Description</summary>
  Given the `head` of a singly linked list, return the middle node of the linked list.

  If there are two middle nodes, return the second middle node.

- [953. Verifying an Alien Dictionary](https://leetcode.com/problems/verifying-an-alien-dictionary/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/953_verifying_an_alien_dictionary.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/953_verifying_an_alien_dictionary_test.go)) <details> <summary>Description</summary>
  In an alien language, surprisingly, they also use English lowercase letters, but possibly in a different `order`. The `order` of the alphabet is some permutation of lowercase letters.

  Given a sequence of `words` written in the alien language, and the `order` of the alphabet, return `true` if and only if the given `words` are sorted lexicographically in this alien language.

- [1021. Remove Outermost Parentheses](https://leetcode.com/problems/remove-outermost-parentheses/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1021_remove_outermost_parentheses.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1021_remove_outermost_parentheses_test.go)) <details> <summary>Description</summary>
  A valid parentheses string is either empty `""`, `"(" + A + ")"`, or `A + B`, where `A` and `B` are valid parentheses strings, and `+` represents string concatenation.

  For example, `""`, `"()"`, `"(())()"`, and `"(()(()))"` are all valid parentheses strings.
  
  A valid parentheses string `s` is primitive if it is nonempty, and there does not exist a way to split it into `s = A + B`, with `A` and `B` nonempty valid parentheses strings.

  Given a valid parentheses string `s`, consider its primitive decomposition: `s = P1 + P2 + ... + Pk`, where `Pi` are primitive valid parentheses strings.

  Return `s` after removing the outermost parentheses of every primitive string in the primitive decomposition of `s`.

- [1108. Defanging an IP Address](https://leetcode.com/problems/defanging-an-ip-address/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1108_defanging_an_ip_address.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1108_defanging_an_ip_address_test.go)) <details> <summary>Description</summary>
  Given a valid (IPv4) IP address, return a defanged version of that IP address.

  A defanged IP address replaces every period `"."` with `"[.]"`.

- [1221. Split a String in Balanced Strings](https://leetcode.com/problems/split-a-string-in-balanced-strings/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1221_split_a_string_in_balanced_strings.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1221_split_a_string_in_balanced_strings_test.go)) <details> <summary>Description</summary>
  Balanced strings are those that have an equal quantity of `'L'` and `'R'` characters.

  Given a balanced string `s`, split it in the maximum amount of balanced strings.

  Return the maximum amount of split balanced strings.
  
- [1232. Check If It Is a Straight Line](https://leetcode.com/problems/check-if-it-is-a-straight-line/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1232_check_if_it_is_a_straight_line_test.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1232_check_if_it_is_a_straight_line_test.go)) <details> <summary>Description</summary>
You are given an array `coordinates`, `coordinates[i] = [x, y]`, where `[x, y]` represents the coordinate of a point. Check if these points make a straight line in the XY plane.

- [1281. Subtract the Product and Sum of Digits of an Integer](https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1281_subtract_digits.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1281_subtract_digits_test.go)) <details> <summary>Description</summary>
  Given an integer number `n`, return the difference between the product of its digits and the sum of its digits.
  
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

- [1365. How Many Numbers Are Smaller Than the Current Number](https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1365_how_many_numbers_are_smaller_than_the_current_number.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1365_how_many_numbers_are_smaller_than_the_current_number_test.go)) <details> <summary>Description</summary>
  Given the array `nums`, for each `nums[i]` find out how many numbers in the array are smaller than it. That is, for each `nums[i]` you have to count the number of valid j's such that `j != i` and `nums[j] < nums[i]`.

  Return the answer in an array.
  
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

- [1436. Destination City](https://leetcode.com/problems/destination-city/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1436_destination_city.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1436_destination_city_test.go)) <details> <summary>Description</summary>
  You are given the array `paths`, where `paths[i] = [cityAi, cityBi]` means there exists a direct path going from `cityAi` to `cityBi`. Return the destination city, that is, the city without any path outgoing to another city.

  It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be exactly one destination city.

- [1470. Shuffle the Array](https://leetcode.com/problems/shuffle-the-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1470_shuffle_the_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1470_shuffle_the_array_test.go)) <details> <summary>Description</summary>
  Given the array `nums` consisting of `2n` elements in the form `[x1,x2,...,xn,y1,y2,...,yn]`.

  Return the array in the form `[x1,y1,x2,y2,...,xn,yn]`.
  
- [1480. Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1480_running_sum_of_1d_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1480_running_sum_of_1d_array_test.go)) <details> <summary>Description</summary>
  Given an array `nums`. We define a running sum of an array as `runningSum[i] = sum(nums[0]???nums[i])`.

  Return the running sum of `nums`.

- [1486. XOR Operation in an Array](https://leetcode.com/problems/xor-operation-in-an-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1486_xor_operation_in_an_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1486_xor_operation_in_an_array_test.go)) <details> <summary>Description</summary>
  You are given an integer `n` and an integer `start`.

  Define an array nums where `nums[i] = start + 2 * i` (0-indexed) and `n == nums.length`.

  Return the bitwise XOR of all elements of `nums`.

- [1502. Can Make Arithmetic Progression From Sequence](https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1502_can_make_arithmetic_progression.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1502_can_make_arithmetic_progression_test.go)) <details> <summary>Description</summary>
  A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.

  Given an array of numbers `arr`, return `true` if the array can be rearranged to form an arithmetic progression. Otherwise, return `false`.

- [1512. Number of Good Pairs](https://leetcode.com/problems/number-of-good-pairs/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1512_number_of_good_pairs.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1512_number_of_good_pairs_test.go)) <details> <summary>Description</summary>
  Given an array of integers `nums`, return the number of **good pairs**.

  A pair `(i, j)` is called good if `nums[i] == nums[j]` and `i < j`.

- [1528. Shuffle String](https://leetcode.com/problems/shuffle-string/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1528_shuffle_string.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1528_shuffle_string_test.go)) <details> <summary>Description</summary>
  You are given a string `s` and an integer array `indices` of the same length. The string `s` will be shuffled such that the character at the `i-th` position moves to `indices[i]` in the shuffled string.

  Return the *shuffled string*.

- [1534. Count Good Triplets](https://leetcode.com/problems/count-good-triplets/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1534_count_good_triplets.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1534_count_good_triplets_test.go)) <details> <summary>Description</summary>
  Given an array of integers `arr`, and three integers `a`, `b` and `c`. You need to find the number of good triplets.

  A triplet `(arr[i], arr[j], arr[k])` is **good** if the following conditions are true:

  - `0 <= i < j < k < arr.length`

  - `|arr[i] - arr[j]| <= a`

  - `|arr[j] - arr[k]| <= b`

  - `|arr[i] - arr[k]| <= c`
  
  Where `|x|` denotes the absolute value of `x`.

  Return *the number of good triplets*.
  
- [1572. Matrix Diagonal Sum](https://leetcode.com/problems/matrix-diagonal-sum/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1572_matrix_diagonal_sum.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1572_matrix_diagonal_sum_test.go)) <details> <summary>Description</summary>
  Given a square matrix `mat`, return the sum of the matrix diagonals.

  Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.
   
- [1588. Sum of All Odd Length Subarrays](https://leetcode.com/problems/sum-of-all-odd-length-subarrays/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1588_sum_of_all_odd_length_subarrays.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1588_sum_of_all_odd_length_subarrays_test.go)) <details> <summary>Description</summary>
Given an array of positive integers `arr`, calculate the sum of all possible odd-length subarrays.

  A subarray is a contiguous subsequence of the array.

  Return the sum of all odd-length subarrays of `arr`.
  
- [1603. Design Parking System](https://leetcode.com/problems/design-parking-system/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1603_design_parking_system.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1603_design_parking_system_test.go)) <details> <summary>Description</summary>
  Design a parking system for a parking lot. The parking lot has three kinds of parking spaces: big, medium, and small, with a fixed number of slots for each size.

  Implement the `ParkingSystem` class:

  - `ParkingSystem(int big, int medium, int small)` Initializes object of the `ParkingSystem` class. The number of slots for each parking space are given as part of the constructor. 
  - `bool addCar(int carType)` Checks whether there is a parking space of `carType` for the car that wants to get into the parking lot. `carType` can be of three kinds: big, medium, or small, which are represented by `1`, `2`, and `3` respectively. A car can only park in a parking space of its `carType`. If there is no space available, return `false`, else park the car in that size space and return `true`.

- [1614. Maximum Nesting Depth of the Parentheses](https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1614_maximum_nesting_depth_of_the_parentheses.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1614_maximum_nesting_depth_of_the_parentheses_test.go)) <details> <summary>Description</summary>
  A string is a valid parentheses string (denoted VPS) if it meets one of the following:

  - It is an empty string `""`, or a single character not equal to `"("` or `")"`,
  
  - It can be written as `AB` (`A` concatenated with `B`), where `A` and `B` are VPS's, or

  - It can be written as `(A)`, where `A` is a VPS.

  We can similarly define the nesting depth `depth(S)` of any VPS `S` as follows:

  - `depth("") = 0`

  - `depth(C) = 0`, where `C` is a string with a single character not equal to `"("` or `")"`.

  - `depth(A + B) = max(depth(A), depth(B))`, where `A` and `B` are VPS's.

  - `depth("(" + A + ")") = 1 + depth(A)`, where `A` is a VPS.

  For example, `""`, `"()()"`, and `"()(()())"` are VPS's (with nesting depths 0, 1, and 2), and `")("` and `"(()"` are not VPS's.

  Given a VPS represented as string `s`, return the nesting depth of `s`.

- [1656. Design an Ordered Stream](https://leetcode.com/problems/design-an-ordered-stream/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1656_design_an_ordered_stream.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1656_design_an_ordered_stream_test.go)) <details> <summary>Description</summary>
  There is a stream of `n` `(idKey, value)` pairs arriving in an arbitrary order, where `idKey` is an integer between `1` and `n` and `value` is a string. No two pairs have the same id.

  Design a stream that returns the values in increasing order of their IDs by returning a chunk (list) of values after each insertion. The concatenation of all the chunks should result in a list of the sorted values.

  Implement the `OrderedStream` class:

  - `OrderedStream(int n)` Constructs the stream to take `n` values.
  - `String[] insert(int idKey, String value)` Inserts the pair `(idKey, value)` into the stream, then returns the largest possible chunk of currently inserted values that appear next in the order.

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

- [1704. Determine if String Halves Are Alike](https://leetcode.com/problems/determine-if-string-halves-are-alike/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1704_determine_if_string_halves_are_alike.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1704_determine_if_string_halves_are_alike_test.go)) <details> <summary>Description</summary>
  You are given a string `s` of even length. Split this string into two halves of equal lengths, and let `a` be the first half and `b` be the second half.

  Two strings are alike if they have the same number of vowels `('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U')`. Notice that `s` contains uppercase and lowercase letters.

  Return `true` if `a` and `b` are alike. Otherwise, return `false`.

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

- [1779. Find Nearest Point That Has the Same X or Y Coordinate](https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1779_nearest_valid_point.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1779_nearest_valid_point_test.go)) <details> <summary>Description</summary>
  You are given two integers, `x` and `y`, which represent your current location on a Cartesian grid: `(x, y)`. You are also given an array `points` where each `points[i] = [ai, bi]` represents that a point exists at `(ai, bi)`. A point is valid if it shares the same x-coordinate or the same y-coordinate as your location.

  Return the index (0-indexed) of the valid point with the smallest Manhattan distance from your current location. If there are multiple, return the valid point with the smallest index. If there are no valid points, return `-1`.

  The Manhattan distance between two points `(x1, y1)` and `(x2, y2)` is `abs(x1 - x2) + abs(y1 - y2)`.
  
- [1790. Check if One String Swap Can Make Strings Equal](https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1790_check_if_one_string_swap_can_make_strings_equal.go), [tests](https://github.com/DimaKoz/goleetcode/blob/330cf42a69deaffbb444bbcf8d3fa150461d9005/easy/1790_check_if_one_string_swap_can_make_strings_equal_test.go#L5)) <details> <summary>Description</summary>
You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.&nbsp;
&nbsp;
Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.&nbsp;

- [1791. Find Center of Star Graph](https://leetcode.com/problems/find-center-of-star-graph/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1791_find_center_of_star_graph.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1791_find_center_of_star_graph_test.go)) <details> <summary>Description</summary>
  There is an undirected star graph consisting of `n` nodes labeled from `1` to `n`. A star graph is a graph where there is one center node and exactly `n - 1` edges that connect the center node with every other node.

  You are given a 2D integer array `edges` where each `edges[i] = [ui, vi]` indicates that there is an edge between the nodes `ui` and `vi`. Return the center of the given star graph.

- [1816. Truncate Sentence](https://leetcode.com/problems/truncate-sentence/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1816_truncate_sentence.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1816_truncate_sentence_test.go)) <details> <summary>Description</summary>
  A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each of the words consists of only uppercase and lowercase English letters (no punctuation).

  - For example, "Hello World", "HELLO", and "hello world hello world" are all sentences.

  You are given a sentence `s` and an integer `k`. You want to truncate `s` such that it contains only the first `k` words. Return `s` after truncating it.

- [1832. Check if the Sentence Is Pangram](https://leetcode.com/problems/check-if-the-sentence-is-pangram/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1832_check_if_the_sentence_is_pangram.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1832_check_if_the_sentence_is_pangram_test.go)) <details> <summary>Description</summary>
  A **pangram** is a sentence where every letter of the English alphabet appears at least once.

  Given a string `sentence` containing only lowercase English letters, return `true` if `sentence` is a pangram, or `false` otherwise.

- [1844. Replace All Digits with Characters](https://leetcode.com/problems/replace-all-digits-with-characters/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1844_replace_all_digits_with_characters.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1844_replace_all_digits_with_characters_test.go)) <details> <summary>Description</summary>
  You are given a **0-indexed string** `s` that has lowercase English letters in its **even** indices and digits in its **odd** indices.

  There is a function `shift(c, x)`, where `c` is a character and `x` is a digit, that returns the `x-th` character after `c`.

  - For example, `shift('a', 5) = 'f'` and `shift('x', 0) = 'x'`.

  For every **odd** index `i`, you want to replace the digit `s[i]` with `shift(s[i-1], s[i])`.

  Return `s` after replacing all digits. It is **guaranteed** that `shift(s[i-1], s[i])` will never exceed `'z'`.

- [1859. Sorting the Sentence](https://leetcode.com/problems/sorting-the-sentence/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1859_sorting_the_sentence.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1859_sorting_the_sentence_test.go)) <details> <summary>Description</summary>
  A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each word consists of lowercase and uppercase English letters.

  A sentence can be shuffled by appending the 1-indexed word position to each word then rearranging the words in the sentence.

  For example, the sentence `"This is a sentence"` can be shuffled as `"sentence4 a3 is2 This1"` or `"is2 sentence4 This1 a3"`.
  Given a shuffled sentence `s` containing no more than `9` words, reconstruct and return the original sentence.

- [1913. Maximum Product Difference Between Two Pairs](https://leetcode.com/problems/maximum-product-difference-between-two-pairs/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1913_maximum_product_difference_between_two_pairs.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1913_maximum_product_difference_between_two_pairs_test.go)) <details> <summary>Description</summary>
  The product difference between two pairs `(a, b)` and `(c, d)` is defined as `(a * b) - (c * d)`.

  - For example, the product difference between `(5, 6)` and `(2, 7)` is `(5 * 6) - (2 * 7) = 16`.

  Given an integer array `nums`, choose four distinct indices `w`, `x`, `y`, and `z` such that the product difference between pairs `(nums[w], nums[x])` and `(nums[y], nums[z])` is maximized.

  Return the maximum such product difference.

- [1920. Build Array from Permutation](https://leetcode.com/problems/build-array-from-permutation/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1920_build_array_from_permutation.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1920_build_array_from_permutation_test.go)) <details> <summary>Description</summary>
  Given a zero-based permutation nums (0-indexed), build an array `ans` of the same length where `ans[i] = nums[nums[i]]` for each `0 <= i < nums.length` and return it.

  A zero-based permutation nums is an array of distinct integers from `0` to `nums.length - 1` (inclusive).

- [1929. Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1929_concatenation_of_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1929_concatenation_of_array_test.go)) <details> <summary>Description</summary>
  Given an integer array `nums` of length `n`, you want to create an array `ans` of length `2n` where `ans[i] == nums[i]` and `ans[i + n] == nums[i]` for `0 <= i < n` (0-indexed).

  Specifically, `ans` is the concatenation of two `nums` arrays.

  Return the array `ans`.

- [1967. Number of Strings That Appear as Substrings in Word](https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1967_number_of_strings_that_appear_as_substrings_in_word.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1967_number_of_strings_that_appear_as_substrings_in_word_test.go)) <details> <summary>Description</summary>
  Given an array of strings `patterns` and a string `word`, return the number of strings in `patterns` that exist as a substring in `word`.

  A substring is a contiguous sequence of characters within a string.

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

- [2103. Rings and Rods](https://leetcode.com/problems/rings-and-rods/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2103_rings_and_rods.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2103_rings_and_rods_test.go)) <details> <summary>Description</summary>
  There are `n` rings and each ring is either red, green, or blue. The rings are distributed across ten rods labeled from `0` to `9`.

  You are given a string `rings` of length `2n` that describes the `n` rings that are placed onto the rods. Every two characters in `rings` forms a color-position pair that is used to describe each ring where:

  - The first character of the `i-th` pair denotes the `i-th` ring's color (`'R'`, `'G'`, `'B'`).
  
  - The second character of the `i-th` pair denotes the rod that the `i-th` ring is placed on (`'0'` to `'9'`).
  
  For example, `"R3G2B1"` describes `n == 3` rings: a red ring placed onto the rod labeled 3, a green ring placed onto the rod labeled 2, and a blue ring placed onto the rod labeled 1.

  Return the number of rods that have all three colors of rings on them.  
  
- [2114. Maximum Number of Words Found in Sentences](https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2114_maximum_number_of_words_found_in_sentences.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2114_maximum_number_of_words_found_in_sentences_test.go)) <details> <summary>Description</summary>
  A sentence is a list of words that are separated by a single space with no leading or trailing spaces.

  You are given an array of strings `sentences`, where each `sentences[i]` represents a single sentence.

  Return the maximum number of words that appear in a single sentence.
  
- [2160. Minimum Sum of Four Digit Number After Splitting Digits](https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2160_minimum_sum_of_four_digit_number_after_splitting_digits.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2160_minimum_sum_of_four_digit_number_after_splitting_digits_test.go)) <details> <summary>Description</summary>
  You are given a positive integer `num` consisting of exactly four digits. Split `num` into two new integers `new1` and `new2` by using the digits found in `num`. Leading zeros are allowed in `new1` and `new2`, and all the digits found in `num` must be used.

  For example, given `num = 2932`, you have the following digits: two `2's`, one `9` and one `3`. Some of the possible pairs `[new1, new2]` are `[22, 93]`, `[23, 92]`, `[223, 9]` and `[2, 329]`.
  
  Return the minimum possible sum of `new1` and `new2`.

- [2176. Count Equal and Divisible Pairs in an Array](https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2176_count_equal_and_divisible_pairs_in_an_array.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2176_count_equal_and_divisible_pairs_in_an_array_test.go)) <details> <summary>Description</summary>
  Given a 0-indexed integer array `nums` of length `n` and an integer `k`, return the number of pairs `(i, j)` where `0 <= i < j < n`, such that `nums[i] == nums[j]` and `(i * j)` is divisible by `k`.  

- [2185. Counting Words With a Given Prefix](https://leetcode.com/problems/counting-words-with-a-given-prefix/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/2185_counting_words_with_a_given_prefix.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/2185_counting_words_with_a_given_prefix_test.go)) <details> <summary>Description</summary>
  You are given an array of strings `words` and a string `pref`.

  Return the number of strings in `words` that contain `pref` as a prefix.

  A prefix of a string `s` is any leading contiguous substring of `s`.

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

  Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: `[???2^31, 2^31 ??? 1]`. For this problem, if the quotient is strictly greater than `2^31 - 1`, then return `2^31 - 1`, and if the quotient is strictly less than `-2^31`, then return `-2^31`.

- [38. Count and Say](https://leetcode.com/problems/count-and-say/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/38_count_and_say.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/38_count_and_say_test.go)) <details> <summary>Description</summary>
  The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

  `countAndSay(1) = "1"`
  `countAndSay(n)` is the way you would "say" the digit string from `countAndSay(n-1)`, which is then converted into a different digit string.
  To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.

  For example, the saying and conversion for digit string `"3322251"`: !["3322251":](https://assets.leetcode.com/uploads/2020/10/23/countandsay.jpg)

  Given a positive integer `n`, return the `n-th` term of the count-and-say sequence.

- [535. Encode and Decode TinyURL](https://leetcode.com/problems/encode-and-decode-tinyurl/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/535_encode_and_decode_tinyurl.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/535_encode_and_decode_tinyurl_test.go)) <details> <summary>Description</summary>
  TinyURL is a URL shortening service where you enter a URL such as `https://leetcode.com/problems/design-tinyurl` and it returns a short URL such as `http://tinyurl.com/4e9iAk`. Design a class to encode a URL and decode a tiny URL.

  There is no restriction on how your encode/decode algorithm should work. You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.

  Implement the `Solution` class:

  - `Solution()` Initializes the object of the system.

  - `String encode(String longUrl)` Returns a tiny URL for the given `longUrl`.

  - `String decode(String shortUrl)` Returns the original long URL for the given `shortUrl`. It is guaranteed that the given `shortUrl` was encoded by the same object.

- [1282. Group the People Given the Group Size They Belong To](https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1282_group_the_people_given_the_group_size_they_belong_to.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1282_group_the_people_given_the_group_size_they_belong_to_test.go)) <details> <summary>Description</summary>
  There are `n` people that are split into some unknown number of groups. Each person is labeled with a **unique ID** from `0` to `n - 1`.

  You are given an integer array `groupSizes`, where `groupSizes[i]` is the size of the group that person `i` is in. For example, if `groupSizes[1] = 3`, then person `1` must be in a group of size `3`.

  Return a list of groups such that each person `i` is in a group of size `groupSizes[i]`.

  Each person should appear in exactly one group, and every person must be in a group. If there are multiple answers, return any of them. It is guaranteed that there will be at least one valid solution for the given input.

- [1329. Sort the Matrix Diagonally](https://leetcode.com/problems/sort-the-matrix-diagonally/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1329_sort_the_matrix_diagonally.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1329_sort_the_matrix_diagonally_test.go)) <details> <summary>Description</summary>
  A matrix diagonal is a diagonal line of cells starting from some cell in either the topmost row or leftmost column and going in the bottom-right direction until reaching the matrix's end. For example, the matrix diagonal starting from `mat[2][0]`, where `mat` is a `6 x 3` matrix, includes cells `mat[2][0]`, `mat[3][1]`, and `mat[4][2]`.

  Given an `m x n` matrix `mat` of integers, sort each matrix diagonal in ascending order and return the resulting matrix.

- [1409. Queries on a Permutation With Key](https://leetcode.com/problems/queries-on-a-permutation-with-key/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/easy/1409_queries_on_a_permutation_with_key.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/easy/1409_queries_on_a_permutation_with_key_test.go)) <details> <summary>Description</summary>
  Given the array `queries` of positive integers between `1` and `m`, you have to process all `queries[i]` (from `i=0` to `i=queries.length-1`) according to the following rules:

  - In the beginning, you have the permutation `P=[1,2,3,...,m]`.

  - For the current `i`, find the position of `queries[i]` in the permutation `P` (indexing from 0) and then move this at the beginning of the permutation `P`. Notice that the position of `queries[i]` in `P` is the result for `queries[i]`.

  Return an array containing the result for the given `queries`.

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

- [1828. Queries on Number of Points Inside a Circle](https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/1828_queries_on_number_of_points_inside_a_circle.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/1828_queries_on_number_of_points_inside_a_circle_test.go)) <details> <summary>Description</summary>
  You are given an array `points` where `points[i] = [xi, yi]` is the coordinates of the `i-th` point on a 2D plane. Multiple points can have the same coordinates.

  You are also given an array `queries` where `queries[j] = [xj, yj, rj]` describes a circle centered at `(xj, yj)` with a radius of `rj`.

  For each query `queries[j]`, compute the number of points inside the `j-th` circle. Points on the border of the circle are considered inside.

  Return an array `answer`, where `answer[j]` is the answer to the `j-th` query.
  
- [2079. Watering Plants](https://leetcode.com/problems/watering-plants/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/2079_watering_plants.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/2079_watering_plants_test.go)) <details> <summary>Description</summary>
  You want to water `n` plants in your garden with a watering can. The plants are arranged in a row and are labeled from `0` to `n - 1` from left to right where the `i-th` plant is located at `x = i`. There is a river at `x = -1` that you can refill your watering can at.

  Each plant needs a specific amount of water. You will water the plants in the following way:

  - Water the plants in order from left to right.

  - After watering the current plant, if you do not have enough water to completely water the next plant, return to the river to fully refill the watering can.
  
  - You cannot refill the watering can early.

  You are initially at the river (i.e., `x = -1`). It takes one step to move one unit on the x-axis.

  Given a 0-indexed integer array `plants` of `n` integers, where `plants[i]` is the amount of water the `i-th` plant needs, and an integer capacity representing the watering can capacity, return the number of steps needed to water all the plants.

- [2120. Execution of All Suffix Instructions Staying in a Grid](https://leetcode.com/problems/execution-of-all-suffix-instructions-staying-in-a-grid/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/2120_execution_of_all_suffix_instructions_staying_in_a_grid.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/2120_execution_of_all_suffix_instructions_staying_in_a_grid_test.go)) <details> <summary>Description</summary>
  There is an `n x n` grid, with the top-left cell at `(0, 0)` and the bottom-right cell at `(n - 1, n - 1)`. You are given the integer `n` and an integer array `startPos` where `startPos = [startrow, startcol]` indicates that a robot is initially at cell `(startrow, startcol)`.

  You are also given a 0-indexed string `s` of length `m` where `s[i]` is the `i-th` instruction for the robot: `'L'` (move left), `'R'` (move right), `'U'` (move up), and `'D'` (move down).

  The robot can begin executing from any `i-th` instruction in `s`. It executes the instructions one by one towards the end of `s` but it stops if either of these conditions is met:

  - The next instruction will move the robot off the grid.
  - There are no more instructions left to execute.
  
  Return an array `answer` of length `m` where `answer[i]` is the number of instructions the robot can execute if the robot begins executing from the `i-th` instruction in `s`.  

- [2125. Number of Laser Beams in a Bank](https://leetcode.com/problems/number-of-laser-beams-in-a-bank/) (solved [here](https://github.com/DimaKoz/goleetcode/blob/main/medium/2125_number_of_laser_beams_in_a_bank.go), [tests](https://github.com/DimaKoz/goleetcode/blob/main/medium/2125_number_of_laser_beams_in_a_bank_test.go)) <details> <summary>Description</summary>
  Anti-theft security devices are activated inside a bank. You are given a 0-indexed binary string array `bank` representing the floor plan of the bank, which is an `m x n` 2D matrix. `bank[i]` represents the `i-th` row, consisting of `'0'`s and `'1'`s. `'0'` means the cell is empty, while `'1'` means the cell has a security device.

  There is one laser beam between any two security devices if both conditions are met:

  - The two devices are located on two different rows: `r1` and `r2`, where `r1 < r2`.

  - For each row `i` where `r1 < i < r2`, there are no security devices in the `i-th` row.

  Laser beams are independent, i.e., one beam does not interfere nor join with another.

  Return the total number of laser beams in the bank.

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
  
  

