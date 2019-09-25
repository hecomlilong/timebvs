package libs
type TimeMap struct {
	KeyTime   map[string][]int
	TimeValue map[int]string
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	timeMap := TimeMap{}
	timeMap.KeyTime = make(map[string][]int, 0)
	timeMap.TimeValue = make(map[int]string, 0)
	return timeMap
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.KeyTime[key] == nil {
		this.KeyTime[key] = []int{}
	}
	this.KeyTime[key] = append(this.KeyTime[key], timestamp)
	this.TimeValue[timestamp] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timestamps := this.KeyTime[key]
	if timestamps == nil || timestamp < timestamps[0] {
		return ""
	}

	low, high, mid, res := 0, len(timestamps)-1, 0, -1
	for low <= high {
		mid = (low + high) / 2
		if timestamp < timestamps[mid] {
			high = mid - 1
		} else if timestamp > timestamps[mid] {
			low = mid + 1
		} else {
			res = timestamps[mid]
			break
		}
	}
	if res == -1 {
		if mid > 0 {
			if timestamp < timestamps[mid] && timestamp >= timestamps[mid - 1] {
				res = timestamps[mid - 1]
			} else {
				res = timestamps[mid]
			}
		} else {
			if timestamp < timestamps[mid] {
				return ""
			} else {
				res = timestamps[mid]
			}
		}
	}

	return this.TimeValue[res]
}

