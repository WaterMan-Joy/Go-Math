// 1. 함수의 이름은
// 2. 매개변수로 정수 p를 받는다
// 3. int 타입의 slice sums 를 선언한다
// 4. i가 2부터 p보다 작거나 같을 때까지 반복한다
// 5. p를 i로 나눈 나머지가 0일 경우, p를 i로 나누고 sums에 i를 추가한다
// 6. 5번의 과정이 끝난 후, i를 1증가 시킨다
// 7. 4~6번의 과정이 끝난 후, sums를 반환한다
// 8. 함수 안에서 fmt 패키지를 import 한다
// 9. 함수 안에서 p를 0으로 나누는 경우가 없도록 조건문을 작성한다
// 10. 함수 안에서 소인수분해 결과를 출력하는 fmt.Println(sums) 함수를 호출한다
package main

import "fmt"

func primeFactorization(p int) []int {
	sums := []int{}

	for i := 2; i <= p; i++ {
		for p%i == 0 {
			p /= i
			sums = append(sums, i)
		}
	}
	fmt.Println(sums)
	return sums
}
