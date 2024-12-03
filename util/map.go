package util

func ToFrequencyMap(slice []int) map[int]int {
    m := make(map[int]int)

    for i := 0; i < len(slice); i++ {
        index := slice[i]
        _, ok := m[index]
        if !ok {
            m[index] = 1
        } else {
            m[index]++
        }
    }
    return m
}
