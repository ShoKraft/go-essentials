package types

import (
	"fmt"
	"math/cmplx"
)

/// Built-in types in Go:
/*
bool
string

int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64, uintptr

float32, float64
complex64, complex128

byte  // alias for uint8
rune  // alias for int32 (Unicode code point)

any   // alias for interface{} (Go 1.18+)
error // built-in interface type

+ composite: array, slice, map, chan, func, struct, interface, pointer

*/

/// Золотые правила Go-разработчика 2026 года (пиши на стикере):
/*

int — по умолчанию для целых чисел

int64 — для денег и ID

uint64 — для битовых флагов и длин > 2³¹

string — везде, где текст

[]byte — везде, где бинарные данные

rune — только в range по строке

float64 — единственный float, который ты используешь

bool — не бойся, он занимает всего 1 байт

НИКОГДА не используй float32, int32, uint32 без очень веской причины (99.9% случаев — нет причины)

*/

func TypesExample() {
	fmt.Print("=== ВСТРОЕННЫЕ ТИПЫ в Go ===\n")

	// ==================== БАЗОВЫЕ ТИПЫ ====================
	var b bool = true
	var s string = "Привет, Go!"

	// ==================== ЦЕЛЫЕ ЧИСЛА ====================
	var i int = 42 // платформо-зависимый (32 или 64 бит)
	_ = i
	var i64 int64 = 9223372036854775807 // максимальное значение int64
	var i32 int32 = 2147483647
	_ = i32
	var i8 int8 = 127
	_ = i8

	var u uint = 42
	_ = u
	var u64 uint64 = 18446744073709551615 // максимальное значение uint64

	var u8 uint8 = 255 // часто используется как byte
	_ = u8

	// ==================== ПЛАВАЮЩАЯ ТОЧКА ====================
	var f32 float32 = 3.14159
	_ = f32
	var f64 float64 = 3.141592653589793
	_ = f64

	// ==================== КОМПЛЕКСНЫЕ ЧИСЛА ====================
	var c64 complex64 = 1 + 2i
	_ = c64
	var c128 complex128 = 3 + 4i
	_ = c128
	// ==================== СПЕЦИАЛЬНЫЕ АЛИАСЫ ====================
	var by byte = 'A'     // alias для uint8
	var r rune = '€'      // alias для int32 (Unicode code point)
	var anyValue any = 42 // alias для interface{} (Go 1.18+)
	_ = anyValue

	fmt.Println("bool:", b)
	fmt.Println("string:", s)
	fmt.Println("int64 (ID):", i64)
	fmt.Println("uint64 (hash):", u64)
	fmt.Println("byte:", by)
	fmt.Println("rune:", r, "→", string(r))

	// ==================== РЕАЛЬНЫЕ СЦЕНАРИИ ИСПОЛЬЗОВАНИЯ ====================

	// 1. int64 — ID пользователей, заказов и т.д.
	userID := int64(987654321098765432)
	fmt.Printf("\n1. int64 → User ID: %d\n", userID)

	// 2. uint64 — хэши, счётчики, битовые флаги
	var hash uint64 = 0xdeadbeefcafebabe
	fmt.Printf("2. uint64 → Hash: 0x%x\n", hash)

	// 3. []byte — работа с бинарными данными, файлами, сетью, криптографией
	data := []byte("Go is awesome!")
	fmt.Printf("3. []byte → %v (длина: %d)\n", data[:5], len(data))
	fmt.Println("   Как строка:", string(data))

	// 4. rune — работа с отдельными Unicode символами
	emoji := rune('🚀')
	cyrillic := rune('я')
	fmt.Printf("4. rune → %c (%U) и %c (%U)\n", emoji, emoji, cyrillic, cyrillic)

	// 5. complex128 — научные расчёты, сигналы, физика
	z := complex(3, 4) // 3 + 4i
	fmt.Printf("\n5. complex128 → z = %v\n", z)
	fmt.Printf("   |z| = %.2f (модуль)\n", cmplx.Abs(z))
	fmt.Printf("   sqrt(-1) = %v\n", cmplx.Sqrt(-1))

	// Дополнительно: деньги лучше хранить в int64 (в копейках/центах)
	moneyInCents := int64(19999) // 199.99 рублей
	fmt.Printf("\nДеньги (int64): %.2f руб.\n", float64(moneyInCents)/100)
	println()
}
