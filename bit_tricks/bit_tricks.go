package bit_tricks

import "fmt"

// Где реально применяется n % (2^power) через & 
/*

	Хэш-таблицы (HashMap) - Чтобы получить индекс корзины: hash & (size-1)

	Кольцевой буфер (Ring Buffer / Circular Queue) - Быстрый переход по кругу в массиве

	Игры и графика - Выравнивание текстур, тайлов, спрайтов (Power of Two textures)

	Менеджеры памяти / Аллокаторы - Выравнивание адресов памяти по 16, 32, 64 байта

	Криптография и RNG - Работа с битами, масками

	Буферизация данных - Быстрое деление на 512, 1024, 4096 и т.д.

*/

// FastModPowerOfTwo возвращает n % (2^power) очень быстро через битовые операции

func FastModPowerOfTwo(n uint64, power uint) uint64 {

	// Проверка, чтобы power не был слишком большим (максимум 63 для uint64)

	if power >= 64 {
		fmt.Printf("Because power >= 64, result = 0")
		return 0
	}

	result := n & ((1 << power) - 1)

	fmt.Printf("%d %% (2^%d) = %d\n", n, power, result)

	return result
}
