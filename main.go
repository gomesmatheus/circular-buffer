package main

import "fmt"

func main() {
    bufferCapacity := 7
    writeIndex := 0
    readIndex := 0
    buffer := make([]int, bufferCapacity, bufferCapacity)

    print(&buffer)
    write(1, &writeIndex, &readIndex, &buffer)
    print(&buffer)
    write(2, &writeIndex, &readIndex, &buffer)
    write(3, &writeIndex, &readIndex, &buffer)
    print(&buffer)
    write(3, &writeIndex, &readIndex, &buffer)
    write(4, &writeIndex, &readIndex, &buffer)
    write(5, &writeIndex, &readIndex, &buffer)
    write(6, &writeIndex, &readIndex, &buffer)
    write(7, &writeIndex, &readIndex, &buffer)
    write(8, &writeIndex, &readIndex, &buffer)
    write(9, &writeIndex, &readIndex, &buffer)
    print(&buffer)
    write(10, &writeIndex, &readIndex, &buffer)
    write(11, &writeIndex, &readIndex, &buffer)
    print(&buffer)
    read(&readIndex, &buffer)
    read(&readIndex, &buffer)
    print(&buffer)
    read(&readIndex, &buffer)
    read(&readIndex, &buffer)
    print(&buffer)
}

func write(val int, writeIndex *int, readIndex *int, buffer *[]int) {
    if (*buffer)[*writeIndex] != 0 {
        *readIndex = ((*readIndex + 1) % cap(*buffer))
    }
    (*buffer)[*writeIndex] = val
    *writeIndex = ((*writeIndex + 1 ) % cap(*buffer))
}

func read(readIndex *int, buffer *[]int) int {
    val := (*buffer)[*readIndex]
    (*buffer)[*readIndex] = 0
    *readIndex = ((*readIndex + 1) % cap(*buffer))
    return val
}

func print(buffer *[]int) {
    for _, i := range *buffer {
        fmt.Print(i, " ")
    }
    fmt.Print("\n")
}

