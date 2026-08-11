type MaxHeap [][]int
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([]int))}
func (h *MaxHeap) Pop() interface{} {
old := *h
n := len(old)
pop := old[n-1]
*h = old[:n-1]
return pop
}

// Less for MaxHeap (Largest element climbs to the top)
func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }

type Twitter struct {
	count int
	tweetMap map[int][][]int
	followMap map[int]map[int]bool
}


func Constructor() Twitter {
    return Twitter{
		count : 0,
		tweetMap : map[int][][]int{}, // map[userId] []{[cnt, tweetId], [cnt, tweetId]}
		followMap : map[int]map[int]bool{},
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.count++
	if _, ok := this.tweetMap[userId]; !ok {
		this.tweetMap[userId] = [][]int{}
	}
	this.tweetMap[userId] = append(this.tweetMap[userId],[]int{this.count, tweetId})
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	if this.followMap[userId] == nil {
        this.followMap[userId] = map[int]bool{}
    }
	this.followMap[userId][userId] = true
	following := this.followMap[userId]
	for followId, _ := range following {
		tweets := this.tweetMap[followId]
		if len(tweets) > 0 {
			index := len(tweets) - 1
			count, tweetId := tweets[index][0], tweets[index][1]
			// adds only latest tweets of all followers
			heap.Push(maxHeap, []int{count, tweetId, followId, index - 1})
		}
	}

	res := []int{}
	for len(res) < 10 && maxHeap.Len() > 0 {
		curr := heap.Pop(maxHeap).([]int)
		count, tweetId, followId, index := curr[0], curr[1], curr[2], curr[3]
		res = append(res, tweetId)

		// adds the next latest tweet into heap of the same userId
		if index >= 0 {
			tweets := this.tweetMap[followId]
			count, tweetId = tweets[index][0], tweets[index][1]
			heap.Push(maxHeap, []int{count, tweetId, followId, index - 1})
		}
	}

	return res
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	if this.followMap[followerId] == nil {
		this.followMap[followerId] = map[int]bool{}
	}
    this.followMap[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.followMap[followerId] != nil {
		delete(this.followMap[followerId], followeeId)
	}
}
