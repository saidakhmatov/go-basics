package main

import (
    "fmt"
)

func main() {
    fmt.Println(MySquareRoot(10, 12))
}

func MySquareRoot(num, precision uint) (result float64) {
    // DO NOT USE math.Sqrt!
    start := 0
    end := int(num)
    mid := 0
    
    var answer float64 = 0

    for start <= end {
        mid = (start + end) / 2;
        if (mid * mid == int(num)) {
            answer = float64(mid)
            break
        }
 
        if (mid * mid < int(num)) {
            start = mid + 1;
            answer = float64(mid);
        }else {
            end = mid - 1;
        }
    }

    var increment float64 = 0.1;

    for i := 0; uint(i) < precision; i++ {
        for answer * answer <= float64(num) {
            answer += increment;
        }
 
        answer = answer - increment;
        increment = increment / 10;
    }
    return answer;
}

// result: 3.162277660168

