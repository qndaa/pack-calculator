# Pack Calculator

## The Problem

You have packs of fixed sizes (like 23, 31, 53 items) and need to ship a certain number of items to a customer. You can't break packs open, so you might need to send slightly more than requested. The goal is to:

1. **Send as few extra items as possible** (minimize overage)
2. **Use as few packs as possible** (minimize total pack count)

## How It Works

### Step 1: Figure out how many items to actually send

Since you can't break packs, if someone orders 500,000 items and your smallest pack is 23, you need to find a reachable number between 500,000 and 500,023 (worst case you'd overshoot by one smallest pack).

### Step 2: Dynamic Programming magic

The code builds a table where `dp[i]` = "what's the minimum number of packs needed to get exactly `i` items?"

It also tracks `parent[i]` = "which pack size did I use to reach `i` items?"

For each number from 1 to max, it checks: "if I add each pack size, can I reach this number with fewer packs than before?"

### Step 3: Backtrack to get the answer

Once we know how many items to send, we walk backwards using the `parent` array: "I'm at 500,000... I got here using a 53-pack... so now I'm at 499,947... got here with another 53..." and count up how many of each pack we used.

### Step 4: Sort and return

Sort the results by pack size (biggest first) so the output is consistent, then return it.

---

That's it! Classic coin-change DP problem with a twist — we're allowed to overshoot slightly to hit a valid combination.
