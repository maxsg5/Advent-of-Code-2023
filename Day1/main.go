package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func IsNum(char byte) (int, bool) {
    num, err := strconv.Atoi(string(char))
    if err != nil {
        return 0, false
    }
    return num, true
}

func main() {

    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

	var lines []int;
    
    for scanner.Scan() {
		str := scanner.Text();
		var first int; 
		var last int;
		isFirst := true;

        for i := 0; i < len(str); i++ {
			num, ok := IsNum(str[i]);
			if ok && isFirst {
				first = num;
				last = num;
				isFirst = false;
			}else if ok {
				last = num;
			}

		}

		num := first * 10 + last;
		lines = append(lines, num);

    }
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
    }

	var total int;
	for _, num := range lines {
		total += num;
	}
	fmt.Printf("Total: %d\n", total)
}
