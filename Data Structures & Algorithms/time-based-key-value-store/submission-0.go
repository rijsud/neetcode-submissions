type TimeMap struct {
	hashMap map[string][]pair
}

type pair struct {
	timestamp int
	value string
}

func Constructor() TimeMap {
	return TimeMap{hashMap : make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.hashMap[key] = append(this.hashMap[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	res := ""
	pairs, ok := this.hashMap[key]
	if !ok {
		return res
	}

	l, r := 0, len(pairs) - 1
	for l <= r {
		mid := l + ((r - l) / 2)
		if pairs[mid].timestamp <= timestamp {
			res = pairs[mid].value
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
