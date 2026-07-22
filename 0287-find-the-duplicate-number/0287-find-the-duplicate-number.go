func findDuplicate(nums []int) int {
    var slow,fast int = nums[0],nums[0]
    for{
        slow = nums[slow] // 1 one jump
        fast = nums[nums[fast]] // 2 jump
        if slow==fast{
            break
        }
    }
    slow = nums[0]
    for slow!=fast{
        slow = nums[slow]  // 1 one jump
        fast = nums[fast] // 1 one jump
    }
    return slow
}