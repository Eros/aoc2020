package _1

import (
	"bufio"
	"os"
	"strconv"
)

func PartOne(numbers []int, setNumbers []int) int {
	for n := range numbers {
		for sn := range setNumbers {
			if n == setNumbers[sn] {
				return n * (2020 - n)
			}
		}
	}
	return 0
}

func PartTwo(numbers []int, setNumbers []int) int {
	for n := range numbers {
		for n2 := range numbers {
			for sn := range setNumbers {
				if (2020 - n - n2) == setNumbers[sn] {
					return n * (2020 - n - n2) * n2
				}
			}
		}
	}
	return 0
}

func main(fileName string) error {
	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}

		defer file.Close()

		scanner := bufio.NewReader(file)

		var line string

		for {
			line, err = scanner.ReadString('\n')

			if err != nil {
				return err
			}
			conv, _ := strconv.Atoi(line)
			//todo fix
			PartOne(conv, conv)
			PartTwo(conv, conv)
		}
	}
	return nil
}
