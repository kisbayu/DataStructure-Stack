package main

import "fmt"

func main() {
    //stack - FILO
    stack := []int{}

    stack = push(stack, 1)
    stack = push(stack, 2)
    stack = push(stack, 3) 
    printStack(stack)
    stack = pop(stack)
    stack = pop(stack)
    stack = push(stack, 4)
    stack = push(stack, 5)
    printStack(stack)
}

func pop(arr []int) []int{
    arr = arr[0:len(arr)-1]
    return arr
}

func push(arr []int, val int) []int{
    arr = append(arr, val)
    return arr
}

func printStack(arr []int){
    for _, val := range arr{
        fmt.Printf("%v ",val)
    }
    fmt.Printf("\n")
}