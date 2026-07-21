func findMissingAndRepeatedValues(grid [][]int) []int {
    //grid = [[9,1,7],[8,9,2],[3,4,6]]
    n := len(grid)
    var repeated,missing int = -1,-1
    var actualSum,expectedSum = 0,0
    seen := make(map[int]struct{})

    // find repeated numbers 
    for _,row := range grid {
        for _,elem := range row{
            actualSum += elem //49
            if _,ok:=seen[elem];ok{
                repeated = elem //9
            }
        seen[elem]=struct{}{}
        }
    }

    // logic of finding missing number
    expectedSum = (n*n)*(n*n + 1)/2 //45
    missing = expectedSum - actualSum + repeated //5

    return []int{repeated,missing}
}