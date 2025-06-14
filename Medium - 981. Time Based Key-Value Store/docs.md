# 981. Time Based Key-Value Store

## Problem

Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the `TimeMap` class:

- `TimeMap()` Initializes the object of the data structure.
- `void set(String key, String value, int timestamp)` Stores the key `key` with the value `value` at the given time `timestamp`.
- `String get(String key, int timestamp)` Returns a value such that `set` was called previously, with `timestamp_prev <= timestamp`. If there are multiple such values, it returns the value associated with the largest `timestamp_prev`. If there are no values, it returns `""`.

**Example 1:**

```
Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
timeMap.get("foo", 1);         // return "bar"
timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
timeMap.get("foo", 4);         // return "bar2"
timeMap.get("foo", 5);         // return "bar2"

```

**Constraints:**

- `1 <= key.length, value.length <= 100`
- `key` and `value` consist of lowercase English letters and digits.
- `1 <= timestamp <= 107`
- All the timestamps `timestamp` of `set` are strictly increasing.
- At most `2 * 105` calls will be made to `set` and `get`.

## Approach
To solve this problem efficiently, we can use a binary search technique. This is because whenever we set a value, the problem states the the timestamps are strictly increasing. There are also no functions for removing or overwriting timestamped values. Therfore, the values for any given key will be sorted by timestamp. Using this, we can deduce the binary search will be the most efficient for retrieval of values.

### `Set` function
For setting, we can use two maps both of which take arrays as their values. One that stores the timestamps for a given key and the other that stores the values for the given key. Since there is no removing or overwriting and timestamps are strictly increasing when setting, the `Set` function is just as simple as appending the value and timestamp to the array in the respective map.

For example, if we are setting the key `"foo"`, with value `"bar"` at timestamp `1`, we simply append the value `"bar"` to the array in the `values` map at key `"foo"` and do the same for timestamp. This way they are linked via index. Our resulting maps would be as follows:
```go
type TimeMap struct {
    timestamps: {
        "foo": [1]
    },
    values: {
        "foo": ["bar"]
    }
}

```
After adding another value to `"foo"`, say `"bar2"` at timestamp `4`, we have:

```go
type TimeMap struct {
    timestamps: {
        "foo": [1, 4]
    },
    values: {
        "foo": ["bar", "bar2"]
    }
}

```

### `Get` function
For the `Get` function, we'll use the binary search approach. To start, we'll check if the key exists and then extract the `timestamps` and `values` arrays. Next, we'll do a binary search as usual, if the midpoint equals the target timestamp, we simply return the value at that index. Otherwise, if the midpoint is less than the target timestamp, we save that index as a potential result (in the case that we don't find the actual value), and move to the left side of the array. 

If the iteration completes and we haven't found the value, we check of our tracked index value is not `-1` (i.e. we had a smaller value than the target timestamp), we return the value at that index, otherwise we return an empty string as per the problem description.


## Complexity
### Time: `O(log n)`
The pertinent function in this problem is the `Get` function, which does a binary search so it is `O(log n)`. The `Set` function runs in constant time.

### Space: `O(n)`
We have maps for the timestamps and values both of which are `O(n)`.

## Solution

```go
type TimeMap struct {
	timestamps map[string][]int
	values     map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{
		timestamps: make(map[string][]int),
		values:     make(map[string][]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timestamps[key] = append(this.timestamps[key], timestamp)
	this.values[key] = append(this.values[key], value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timestamps, exists := this.timestamps[key]
	if !exists {
		return ""
	}
	values, exists := this.values[key]
	if !exists {
		return ""
	}


	i := -1
	l, r := 0, len(timestamps)-1
	for l <= r {
		m := (l + r) / 2
		if timestamps[m] == timestamp {
			return values[m]
		} else if timestamps[m] < timestamp {
			i = m
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if i == -1 {
		return  ""
	}

	return values[i]
}
```
