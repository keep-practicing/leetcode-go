# Go Solution for LeetCode algorithm problems

[![Build Status](https://travis-ci.org/zwfang/leetcode.svg?branch=master)](https://travis-ci.org/zwfang/leetcode)
[![codecov](https://codecov.io/gh/zwfang/leetcode/branch/master/graph/badge.svg)](https://codecov.io/gh/zwfang/leetcode)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/86cf2613fa544ab5b254e2a7e5d9deb8)](https://www.codacy.com/app/zwfang/leetcode?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=zwfang/leetcode&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/zwfang/leetcode)](https://goreportcard.com/report/github.com/zwfang/leetcode)

continually updating 😃.

### [view by sorted](./src/README.md)

### Array
* [1. Two Sum](./src/0001_two_sum/twosum.go)&nbsp;&nbsp;&nbsp;*`lookup table;`*&nbsp;&nbsp;*`hash table`*
* [4. Median of Two Sorted Arrays](src/0004_median_of_two_sorted_arrays/motsa.go)&nbsp;&nbsp;&nbsp;*`binary search;`*&nbsp;&nbsp;*`divide and conquer`*
* [11. Container With Most Water](./src/0011_container_with_most_water/container_with_most_water.go)&nbsp;&nbsp;&nbsp;*`double index;`*&nbsp;&nbsp;*`array`*
* [26. Remove Duplicates from Sorted Array](./src/0026_remove_duplicates_from_sorted_array/rdfsa.go)&nbsp;&nbsp;&nbsp;*`double index;`*&nbsp;&nbsp;*`array`*
* [27. Remove Element](src/0027_remove_element/remove_element.go)&nbsp;&nbsp;&nbsp;*`double index;`*&nbsp;&nbsp;*`array`*
* [48. Rotate Image](src/0048_rotate_image/rotate_image.go)
* [53. Maximum Subarray](src/0053_maximum_subarray/maximum_subarray.go)&nbsp;&nbsp;&nbsp;*`dynamic programming`*
* [75. Sort Colors](./src/0075_sort_colors/sort_colors.go)&nbsp;&nbsp;&nbsp;*`sort;`*&nbsp;&nbsp;*`array`*
* [80. Remove Duplicates from Sorted Array II](./src/0080_remove_duplicates_from_sorted_array2/rdfsa2.go)&nbsp;&nbsp;&nbsp;*`double index;`*&nbsp;&nbsp;*`array`*
* [88. Merge Sorted Array](./src/0088_merge_sorted_array/msa.go)&nbsp;&nbsp;&nbsp;*`sort;`*&nbsp;&nbsp;*`array`*
* [121. Best Time to Buy and Sell Stock](src/0121_best_time_to_buy_and_sell_stock/maxprofit.go)&nbsp;&nbsp;&nbsp;*`dynamic programming;`*&nbsp;&nbsp;*`array`*
* [122. Best Time to Buy and Sell Stock II](src/0122_best_time_to_buy_and_sell_stock_2/maxprofit.go)&nbsp;&nbsp;&nbsp;*`greedy;`*&nbsp;&nbsp;*`array`*
* [167. Two Sum II - Input array is sorted](./src/0167_two_sum2/two_sum2.go)&nbsp;&nbsp;&nbsp;*`double index;`*&nbsp;&nbsp;*`binary search`*
* [209. Minimum Size Subarray Sum](./src/0209_minimum_size_subarray_sum/minimum_size_subarray_sum.go)&nbsp;&nbsp;&nbsp;*`sliding window`*
* [215. Kth Largest Element in an Array](./src/0215_kth_largest_element_in_an_array/kthleiaa.go)&nbsp;&nbsp;&nbsp;*`sort`*
* [283. Move Zeroes(solution1)](./src/0283_move_zeroes/move_zeroes.go)&nbsp;&nbsp;&nbsp;*`sliding window`*
* [283. Move Zeroes(solution2)](./src/0283_move_zeroes/move_zeroes2.go)&nbsp;&nbsp;&nbsp;*`sliding window`*
* [303. Range Sum Query - Immutable](src/0303_range_sum_query/rsqim.go)
* [349. Intersection of Two Arrays](./src/0349_intersection_of_2_arrays/intersection_of_two_arrays.go)&nbsp;&nbsp;&nbsp;*`set`*
* [350. Intersection of Two Arrays II](./src/0350_intersection_of_two_arrays2/intersection_of_two_arrays2.go)&nbsp;&nbsp;&nbsp;*`hash table`*
* [447. Number of Boomerangs](./src/0447_number_of_boomerangs/number_of_boomerangs.go)&nbsp;&nbsp;&nbsp;*`hash table`*
* [454. 4Sum II](./src/0454_4sum2/4sum2.go)&nbsp;&nbsp;&nbsp;*`hash table`*
* [674. Longest Continuous Increasing Subsequence](src/0674_longest_continuous_increasing_subsequence/lcis.go)
* [713. Subarray Product Less Than K](src/0713_subarray_product_less_than_k/spltk.go)&nbsp;&nbsp;&nbsp;*`sliding window`*&nbsp;&nbsp;*`array`*
* [717. 1-bit and 2-bit Characters](src/0717_1_bit_and_2_bit_characters/1bitand2bitc.go)
* [747. Largest Number At Least Twice of Others](./src/0747_largest_number_at_least_twice_of_others/largest_number_at_least_twice_of_others.go)

### String
* [3. Longest Substring Without Repeating Characters](./src/0003_longest_substring_without_repeating_characters/longest_substring_without_repeating_characters.go)&nbsp;&nbsp;&nbsp;*`sliding window;`*&nbsp;&nbsp;*`hash table`*
* [14. Longest Common Prefix](src/0014_longest_common_prefix/lcp.go)
* [17. Letter Combinations of a Phone Number](./src/0017_letter_combination_of_a_phone_number/letter_combination_of_phone_number.go)&nbsp;&nbsp;&nbsp;*`tree`*
* [20. Valid Parentheses](./src/0020_valid_parentheses/valid_parentheses.go)&nbsp;&nbsp;&nbsp;*`stack`*
* [28. Implement strStr()](src/0028_implement_strstr/implement_strstr.go)&nbsp;&nbsp;&nbsp;*`double index`*
* [58. Length of Last Word](src/0058_length_of_last_word/len_of_last_word.go)
* [67. Add Binary](./src/0067_add_binary/add_binary.go)&nbsp;&nbsp;&nbsp;*`brute force`*
* [76. Minimum Window Substring](./src/0076_minimum_window_substring/minimum_window_substring.go) &nbsp;&nbsp;&nbsp;*`sliding window`*
* [125. Valid Palindrome](./src/0125_valid_palindrome/valid_palindrome.go)&nbsp;&nbsp;&nbsp;*`string;`*&nbsp;&nbsp;*`double index`*
* [344. Reverse String](./src/0344_reverse_string/reverse_string.go)&nbsp;&nbsp;&nbsp;*`string;`*&nbsp;&nbsp;*`double index`*
* [345. Reverse Vowels of a String](./src/0345_reverse_vowels_of_a_string/reverse_vowels.go)&nbsp;&nbsp;&nbsp;*`string;`*&nbsp;&nbsp;*`double index`*
* [438. Find All Anagrams in a String](./src/0438_all_anagrams_in_a_string/all_anagrams_in_a_string.go)&nbsp;&nbsp;&nbsp;*`sliding window`*

### Linked List
* [2. Add Two Numbers](./src/0002_add_two_numbers/add_two_numbers.go)&nbsp;&nbsp;&nbsp;*`recursion;`*&nbsp;&nbsp;*`math`*
* [19. Remove Nth Node From End of List](src/0019_remove_nth_node_from_end_of_list/remove_nth_node_from_end_of_list.go)&nbsp;&nbsp;&nbsp;*`two pointers`*
* [21. Merge Two Sorted Lists](./src/0021_merge_two_sorted_lists/mergeTwoLists.go)
* [23. Merge k Sorted Lists](src/0023_merge_k_sorted_lists/mksl.go)&nbsp;&nbsp;&nbsp;*`heap`*
* [24. Swap Nodes in Pairs](src/0024_swap_nodes_in_pairs/swap_nodes_in_pairs.go)
* [25. Reverse Nodes in k-Group](./src/0025_reverse_nodes_in_k_group/reverse_node_k_group.go)
* [61. Rotate List](./src/0061_rotate_list/rotate_list.go)
* [82. Remove Duplicates from Sorted List II](src/0082_remove_duplicates_from_sorted_list_2/rdfsl.go)
* [83. Remove Duplicates from Sorted List](src/0083_remove_duplicates_from_sorted_list/rdfsl.go)
* [86. Partition List](src/0086_partition_list/partition_list.go)&nbsp;&nbsp;&nbsp;*`two pointers`*
* [92. Reverse Linked List II](src/0092_reverse_linked_list_2/reverse_linked_list2.go)

### Dynamic Programming
* [62. Unique Paths](./src/0062_unique_paths/unique_paths.go)&nbsp;&nbsp;&nbsp;*`array`*
* [63. Unique Paths 2](./src/0063_unique_paths_2/unique_paths2.go)&nbsp;&nbsp;&nbsp;*`array`*
* [64. Minimum Path Sum](./src/0064_minimum_path_sum/minimum_path_sum.go)&nbsp;&nbsp;&nbsp;*`array`*
* [70. Climbing Stairs](./src/0070_climbing_stairs/climbing_stairs.go)
* [120. Triangle](./src/0120_triangle/triangle.go)&nbsp;&nbsp;&nbsp;*`array`*
* [198. House Robber](./src/0198_house_robber/house_robber.go)
* [300. Longest Increasing Subsequence](./src/0300_longest_increasing_subsequence/lis.go)
* [343. Integer Break](./src/0343_integer_break/integer_break.go)&nbsp;&nbsp;&nbsp;*`math`*
* [376. Wiggle Subsequence](./src/0376_wiggle_subsequence/wiggle_subsequence.go)&nbsp;&nbsp;&nbsp;*`greedy;`*&nbsp;&nbsp;*`dynamic programming`*
* [416. Partition Equal Subset Sum](./src/0416_partition_equal_subset_sum/partition_equal_subset_sum.go)&nbsp;&nbsp;*`0-1 knapsack problem`*
* [435. Non-overlapping Intervals](./src/0435_non_overlapping_intervals/dp_solution.go)&nbsp;&nbsp;&nbsp;*`greedy;`*&nbsp;&nbsp;*`dynamic programming`*&nbsp;&nbsp;*`0-1 knapsack problem`*

### Greedy
* [392. Is Subsequence](./src/0392_is_subsequence/is_subsequence.go)
* [435. Non-overlapping Intervals](./src/0435_non_overlapping_intervals/greedy_solution.go)&nbsp;&nbsp;&nbsp;*`greedy;`*&nbsp;&nbsp;*`dynamic programming`*
* [455. Assign Cookies](./src/0455_assign_cookies/assign_cookies.go)

### Backtracking
* [77. Combinations](src/0077_combinations/combinations.go)&nbsp;&nbsp;&nbsp;*`combine`*
* [79. Word Search](src/0079_word_search/word_search.go)&nbsp;&nbsp;&nbsp;*`array`*

### Tree
* [94. Binary Tree Inorder Traversal](./src/0094_binary_tree_inorder_traversal/binary_tree_inorder_traversal.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*&nbsp;&nbsp;&nbsp;*`stack`*
* [100. Same Tree](./src/0100_same_tree/same_tree.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*&nbsp;&nbsp;&nbsp;*`dfs`*
* [101. Symmetric Tree](./src/0101_symmetric_tree/symmetric_tree.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*&nbsp;&nbsp;&nbsp;*`dfs;`*&nbsp;&nbsp;*`bfs;`*
* [102. Binary Tree Level Order Traversal](src/0102_binary_tree_level_order_traversal/binary_tree_level_order_traversal.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*&nbsp;&nbsp;&nbsp;*`dfs`*
* [107. Binary Tree Level Order Traversal II](./src/0107_binary_tree_level_order_traversal_2/binary_tree_level_order_traversal2.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*&nbsp;&nbsp;&nbsp;*`bfs`*
* [111. Minimum Depth of Binary Tree](./src/0111_minimum_depth_of_binary_tree/minimum_depth_of_binary_tree.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*&nbsp;&nbsp;&nbsp;*`dfs`*
* [112. Path Sum](./src/0112_path_sum/path_sum.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*&nbsp;&nbsp;&nbsp;*`dfs`*
* [144. Binary Tree Preorder Traversal](src/0144_binary_tree_preorder_traversal/binary_tree_preorder_traversal.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*&nbsp;&nbsp;&nbsp;*`pre-order traversal`*
* [226. Invert Binary Tree](./src/0226_invert_binary_tree/invert_binary_tree.go)&nbsp;&nbsp;&nbsp;*`binary tree;`*

### Binary Search
* [33. Search in Rotated Sorted Array](./src/0033_search_in_rotated_sorted_array/search_in_rotated_sorted_array.go)&nbsp;&nbsp;&nbsp;*`array;`*&nbsp;&nbsp;*`binary search`*
* [34. Find First and Last Position of Element in Sorted Array](./src/0034_find_first_and_last_position_of_element_in_sorted_array/find_first_and_last_position_of_element_in_sorted_array.go)&nbsp;&nbsp;&nbsp;*`array;`*&nbsp;&nbsp;*`binary search`*
* [35. Search Insert Position](src/0035_search_insert_position/search_insert_position.go)&nbsp;&nbsp;&nbsp;*`array;`*&nbsp;&nbsp;*`binary search`*
* [69. Sqrt(x)](./src/0069_sqrtx/sqrtx.go)&nbsp;&nbsp;&nbsp;*`math;`*&nbsp;&nbsp;*`binary search`*
* [704. Binary Search](./src/0704_binary_search/binary_search.go)

### Math
* [7. Reverse Integer](src/0007_reverse_integer/reverse_integer.go)
* [9. Palindrome Number](src/0009_palindrome_number/palindrome_number.go)
* [13. Roman to Integer](src/0013_roman_to_integer/roman_to_integer.go)&nbsp;&nbsp;&nbsp;*`string`*
* [66. Plus One](src/0066_plus_one/plus_one.go)&nbsp;&nbsp;&nbsp;*`array`*

<details>
</details>
