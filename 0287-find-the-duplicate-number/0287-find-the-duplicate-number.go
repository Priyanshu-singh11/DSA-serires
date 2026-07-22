func findDuplicate(nums []int) int {
    duplicateNum := 0
    seen := make(map[int]struct{})
    for _,num := range nums{
        _,ok:=seen[num];if ok{
            duplicateNum = num
            break
        }
        seen[num] = struct{}{}
    }
    return duplicateNum
}