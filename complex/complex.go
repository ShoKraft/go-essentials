package complex

import (
	"strconv"
)

func PrettyComplex(z complex128) string {
	re := real(z)
	im := imag(z)

	// Случай 1: Чистый ноль
	if re == 0 && im == 0 {
		return "0"
	}

	// Случай 2: Только мнимая часть
	if re == 0 {
		return formatFloat(im) + "i"
	}

	// Случай 3: Только реальная часть
	if im == 0 {
		return formatFloat(re)
	}

	// Случай 4: Обе части (формируем строку вида a + bi или a - bi)
	res := formatFloat(re)
	if im > 0 {
		res += " + " + formatFloat(im) + "i"
	} else {
		res += " - " + formatFloat(-im) + "i" // инвертируем минус для красивого вывода
	}

	return res
}

// Вспомогательная функция для удаления лишних нулей (1.500 -> 1.5)
func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}
