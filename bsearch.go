package binarysearch

func Search(values []int, target int) int {
    return search(values, target, 0, len(values)-1)
}

func search(values []int, target, start, end int) int {
    if start > end {
        return -1
    }

    middle := (start + end) / 2
    // fmt.Printf("target: %d, start: %d, middle: %d, end: %d\n", target, start, middle, end)
    value := values[middle]

    if value > target {
        return search(values, target, start, middle-1)
    }

    if value < target {
        return search(values, target, middle+1, end)
    }

    return middle
}
