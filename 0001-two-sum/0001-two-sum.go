func twoSum(nums []int, target int) []int {
   indexMap :=  make(map[int]int)
   for i,num := range nums{
    sec := target-num
    if index,ok := indexMap[sec];ok{
        return []int{i,index}
    }
    indexMap[num]=i
   }
   return []int{}
}