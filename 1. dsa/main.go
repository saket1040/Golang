package main

// pointer sbko bula skta hai
// value sirf value ko
import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

func sorting() {
	fmt.Println("Hello, World!")
	nums := []int{1, 2, 4, 3, 5}
	fmt.Println("Original slice:", nums)

	sort.Ints(nums)
	fmt.Println("Sorted slice:", nums)

	// sort.Strings([]string)
	// sort.Float64s([]float64)

	str := "ram"
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	sortedStr := string(bytes)
	fmt.Println("Sorted string:", sortedStr)

	strList := []string{"ram", "shyam", "mohan"}
	sort.Strings(strList)
	fmt.Println("Sorted string slice:", strList)

	nodes := make([]Node, 0)
	for _, num := range nums {
		nodes = append(nodes, Node{High: num * 10, Low: num})
	}
	fmt.Println("Nodes slice:", nodes)

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].High < nodes[j].High
	})
	fmt.Println("Sorted nodes slice:", nodes)
}

//map
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	anagramMap := make(map[string][]string)

	for _, word := range strs {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

// Helper function to sort a string
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// intervals
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]

		if last[1] >= intervals[i][0] {
			// Merge intervals
			last[1] = max(last[1], intervals[i][1])
			result[len(result)-1] = last 
		} else {
			// No overlap, add new interval
			result = append(result, intervals[i])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MyCalendarThree struct {
	log map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		log: make(map[int]int),
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	this.log[startTime]++
	this.log[endTime]--

	// Extract keys and sort them to simulate ordered map traversal (like C++ map)
	times := make([]int, 0, len(this.log))
	for time := range this.log {
		times = append(times, time)
	}
	sort.Ints(times)

	count := 0
	maxOverlaps := math.MinInt
	for _, time := range times {
		count += this.log[time]
		if count > maxOverlaps {
			maxOverlaps = count
		}
	}

	return maxOverlaps
}

// Count Intervals
type CountIntervals struct {
	mp      map[int]int
	lastRes int
}

func Constructor() CountIntervals {
	return CountIntervals{
		mp:      make(map[int]int),
		lastRes: -1,
	}
}

func (this *CountIntervals) Add(left int, right int) {
	this.mp[left]++
	this.mp[right+1]--
	this.lastRes = -1
}

func (this *CountIntervals) Count() int {
	if this.lastRes == -1 {
		this.lastRes = 0
		newEvents := make(map[int]int)
		currCount := 0
		start := 0

		// Get sorted keys to simulate ordered map
		keys := make([]int, 0, len(this.mp))
		for k := range this.mp {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			if currCount == 0 {
				start = k
			}
			currCount += this.mp[k]
			if currCount == 0 {
				this.lastRes += k - start
				newEvents[start]++
				newEvents[k]--
			}
		}
		this.mp = newEvents
	}
	return this.lastRes
}

//stack

func nse(arr []int) []int {
	n := len(arr)
	nse := make([]int, n)
	st := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && arr[st[len(st)-1]] >= arr[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			nse[i] = st[len(st)-1]
		} else {
			nse[i] = n
		}
		st = append(st, i)
	}
	return nse
}

func pse(arr []int) []int {
	n := len(arr)
	pse := make([]int, n)
	st := []int{}

	for i := 0; i < n; i++ {
		for len(st) > 0 && arr[st[len(st)-1]] > arr[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			pse[i] = st[len(st)-1]
		} else {
			pse[i] = -1
		}
		st = append(st, i)
	}
	return pse
}

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	mod := int(1e9 + 7)

	nseArr := nse(arr)
	pseArr := pse(arr)

	res := 0

	for i := 0; i < n; i++ {
		left := i - pseArr[i]
		right := nseArr[i] - i
		freq := (left * right) % mod
		res = (res + (arr[i]*freq)%mod) % mod
	}
	return res
}


// linked list
type ListNode struct {
    Val  int
    Next *ListNode
}

func getKthNode(head *ListNode, k int) *ListNode {
    for head != nil && k > 1 {
        head = head.Next
        k--
    }
    return head
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    temp := head
    var prevLast *ListNode = nil

    for temp != nil {
        kthNode := getKthNode(temp, k)
        if kthNode == nil {
            if prevLast != nil {
                prevLast.Next = temp
            }
            break
        }

        nextNode := kthNode.Next
        kthNode.Next = nil
        newHead := reverseList(temp)

        if temp == head {
            head = newHead
        } else {
            prevLast.Next = newHead
        }

        prevLast = temp
        temp = nextNode
    }

    return head
}


// tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func solve(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }

    leftMax := solve(root.Left, ans)
    if leftMax < 0 {
        leftMax = 0
    }

    rightMax := solve(root.Right, ans)
    if rightMax < 0 {
        rightMax = 0
    }

    *ans = max(*ans, root.Val + leftMax + rightMax)

    return root.Val + max(leftMax, rightMax)
}

func maxPathSum(root *TreeNode) int {
    ans := math.MinInt32
    solve(root, &ans)
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//level order traversal
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    result := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)
        level := []int{}

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:] // pop from front

            level = append(level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, level)
    }

    return result
}

// refernce
func getMinimumDifference(root *TreeNode) int {
    res := int(^uint(0) >> 1) // Max int
    var prev *TreeNode

    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }

        inorder(node.Left)

        if prev != nil {
            diff := node.Val - prev.Val
            if diff < res {
                res = diff
            }
        }
        prev = node

        inorder(node.Right)
    }

    inorder(root)
    return res
}

||||||| OR |||||||


func inorder(root *TreeNode, prev **TreeNode, res *int) {
    if root == nil {
        return
    }

    inorder(root.Left, prev, res)

    if *prev != nil {
        diff := root.Val - (*prev).Val
        if diff < *res {
            *res = diff
        }
    }
    *prev = root

    inorder(root.Right, prev, res)
}

func getMinimumDifference(root *TreeNode) int {
    res := int(^uint(0) >> 1) // Max int
    var prev *TreeNode
    inorder(root, &prev, &res)
    return res
}

//bit manipulation

func singleNumber(nums []int) int {
    res := 0

    for i := 0; i < 32; i++ {
        check := 1 << i
        cnt := 0

        for _, num := range nums {
            if num&check != 0 {
                cnt++
            }
        }

        if cnt%3 != 0 {
            res |= check
        }
    }

    // Handle negative numbers (since Go int is signed 32-bit)
    if res >= 1<<31 {
        res -= 1 << 32
    }

    return res
}

// trie

type Node struct {
    links [26]*Node
    flag  bool
}

func (n *Node) get(ch byte) *Node {
    return n.links[ch-'a']
}

func (n *Node) put(ch byte, node *Node) {
    n.links[ch-'a'] = node
}

func (n *Node) setEnd() {
    n.flag = true
}

func (n *Node) getEnd() bool {
    return n.flag
}

func (n *Node) containsKey(ch byte) bool {
    return n.links[ch-'a'] != nil
}

type Trie struct {
    root *Node
}

func Constructor() Trie {
    return Trie{root: &Node{}}
}

func (this *Trie) Insert(word string) {
    node := this.root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if !node.containsKey(ch) {
            node.put(ch, &Node{})
        }
        node = node.get(ch)
    }
    node.setEnd()
}

func (this *Trie) Search(word string) bool {
    node := this.root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if !node.containsKey(ch) {
            return false
        }
        node = node.get(ch)
    }
    return node.getEnd()
}

func (this *Trie) StartsWith(prefix string) bool {
    node := this.root
    for i := 0; i < len(prefix); i++ {
        ch := prefix[i]
        if !node.containsKey(ch) {
            return false
        }
        node = node.get(ch)
    }
    return true
}


// heap

type IntHeap struct {
    data []int
    isMin bool // true for min-heap, false for max-heap
}

func (h IntHeap) Len() int           { return len(h.data) }
func (h IntHeap) Less(i, j int) bool {
    if h.isMin {
        return h.data[i] < h.data[j]
    }
    return h.data[i] > h.data[j]
}

func (h IntHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *IntHeap) Push(x any) {
    h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() any {
    n := len(h.data)
    x := h.data[n-1]
    h.data = h.data[:n-1]
    return x
}