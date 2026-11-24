# Set Matrix Zeroes

> Striver SDE Sheet 001 · Day 1 Arrays — zero out full rows and columns in place, using the first row and column as the marker store. Go, stdlib only, tested.

## Problem

Given an m×n integer matrix, whenever a cell is 0, set that cell's entire row and entire column to 0 — without allocating separate marker arrays for rows and columns.

## Approach

You cannot zero while scanning (fresh zeroes would cascade into rows that never had one), so condemned rows/columns must be remembered first. Instead of O(m+n) marker arrays, reuse the matrix itself: the first row records which columns are condemned, the first column records which rows are. Two booleans remember whether the first row/column themselves contained a zero. Mark in one pass, zero the interior from the marks, then settle the first row and column last.

**Complexity:** O(m·n) time, O(1) extra space.

## Visual walkthrough

Read the picture top to bottom — each numbered frame is one moment of the algorithm
running on this repo's test example.

![Step-by-step visual walkthrough](viz.svg)

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
