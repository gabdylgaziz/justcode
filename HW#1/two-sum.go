func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
    for i, v := range nums {
        j, ok := mp[-v]
        mp[v - target] = i
        if ok {
            return []int{j, i}
        }
    }
    return []int{}
}