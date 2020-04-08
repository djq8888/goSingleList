package goSingleList

func Build(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := ListNode{nums[0], nil}
	prev := &head
	for i := 1; i < len(nums); i++ {
		tmp := ListNode{nums[i], nil}
		prev.Next = &tmp
		prev = &tmp
	}
	return &head
}
