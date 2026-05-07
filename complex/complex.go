package complex

import (
	"fmt"
	"strconv"
)

// PrettyComplex возвращает красивое строковое представление комплексного числа.
//
// Примеры:
//   5.0        → "5"
//   0 + 3i     → "3i"
//   0 - 1i     → "-i"
//   2 + 1i     → "2 + i"
//   4.5 - 2.3i → "4.5 - 2.3i"
func PrettyComplex(z complex128) string {
	re, im := real(z), imag(z) // извлекаем действительную и мнимую части

	switch {
	// 1. Чисто действительное число (мнимая часть равна нулю)
	case im == 0:
		return fmt.Sprintf("%g", re)

	// 2. Чисто мнимое число (действительная часть равна нулю)
	case re == 0:
		if im == 1 {
			return "i"
		}
		if im == -1 {
			return "-i"
		}
		return fmt.Sprintf("%gi", im)

	// 3. Положительная мнимая часть
	case im > 0:
		if im == 1 {
			return fmt.Sprintf("%g + i", re) // особый случай: не пишем "1i", а просто "i"
		}
		return fmt.Sprintf("%g + %gi", re, im)

	// 4. Отрицательная мнимая часть
	default:
		if im == -1 {
			return fmt.Sprintf("%g - i", re) // особый случай: "-i"
		}
		// Преобразуем отрицательное значение в положительное для удобства чтения
		return fmt.Sprintf("%g - %gi", re, -im)
	}
}

// formatFloat — вспомогательная функция для красивого форматирования float64.
//
// Убирает лишние нули после запятой:
//   1.500  → "1.5"
//   2.000  → "2"
//   3.1410 → "3.141"
//
// Пока не используется в PrettyComplex, но может быть полезна при доработке.
func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}