package main

import "fmt"

func main() {
	/*for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}*/
	/*count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}*/

	/*i := 0
	for next := true; next; next = i < 5 {
		i++
		fmt.Println(i)
	}*/

	/*r := rand.New(rand.NewSource(time.Now().UnixNano()))

	nums := make([]int, 10)
	for i := range nums {
		nums[i] = r.Intn(100)
	}

	fmt.Println("Исходный список:", nums)

	// Сортировка пузырьком
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	fmt.Println("Отсортированный список:", nums)
	*/
	//Вложенный цикл
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("i =%d; j=%d\n", i, j)
		}
	}

}
