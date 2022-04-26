# Project description

## Word Search Jumble

## Summary

A popular puzzle format consists of random letters situated in a grid of rows and
columns called a jumble. Inside the jumble there are words that are hidden that you can
find.

## Objective

The objective of this word search task is to be given a target word and a jumble of
letters stored in a 2D string array and determine if the word is in the jumble. The 2D
array contains the grid of characters stored as type string (see example below).

## Guidelines

A word search should occur in 3 directions:
1. Horizontal (left to right) ➡
2. Vertical (top to bottom) ⬇
3. Diagonal (top-left to bottom-right) ↘

If the word is in the jumble in any direction provided above, return the position of the
word's starting letter as X, Y coordinates (starting from 0, 0).
If the word is not in the jumble, return -1, -1


