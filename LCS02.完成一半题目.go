

func halfQuestions(questions []int) int {
    cache := [1001]int{}
    for _, v := range questions{
        cache[v]++
    }
    sort.Ints(cache[:])
    count := len(questions) / 2
    res := 0 
    for i:=1000; count > 0 ;i-- {
        count -= cache[i]
        res++
    }
    return res

