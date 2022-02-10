Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:  
```
Вывод:  
2  
1

defer вызывается после того, как значение в return уже установлено, например:

func anotherTest() int {
	var x int -- создаем переменную
	defer func() {
		x++ -- увеличиваем переменную x, не значение возврата
	}()
	x = 1 -- она равна 1
	return x -- значение возврата равно 1
}



func test() (x int) { -- x является возвратным значением
	defer func() {
		x++ -- увеличиваем возвратное значение
	}()
	x = 1 -- равно 1
	return -- вернуть возвратное значение
}

```
