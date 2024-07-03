package producter

func product(toProduct []int) int {
    result := 1

    for _, v := range(toProduct) {
        result *= v
    }

    return result
}

func Producter(numbersToProduct ...[]int) []int {
    numResults := len(numbersToProduct)
    products := make([]int, numResults)

    for i, array := range numbersToProduct {
        products[i] = product(array)
    }

    return products
}
