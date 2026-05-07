package constatns

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota) // пропускаем 0
	MB                    // 1 048 576
	GB                    // 1 073 741 824
	TB                    // 1 099 511 627 776
	PB                    // 1 125 899 906 842 624
)

func FormatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	suffix := []string{"KB", "MB", "GB", "TB", "PB", "EB"}

	value := float64(b) / float64(div)

	return fmt.Sprintf("%.1f %s", value, suffix[exp])
}

/// Как это работает ?
/*
	1. Константа 1024: Мы используем двоичные префиксы
		(GiB/GB в контексте операционных систем обычно считаются по 1024).

	2. Цикл поиска: Мы делим число на 1024 до тех пор, пока оно не станет меньше тысячи,
		 попутно запоминая «степень» (сколько раз мы разделили).

	3. Форматирование: Функция fmt.Sprintf("%.1f %s", ...) округляет число до десятых
		 и добавляет нужный суффикс из массива.
*/
