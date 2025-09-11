package stringsex

import (
	"regexp"
)

// ProcessStringAny обрабатывает строку, заменяя шаблоны ${key} на значения из mapping
func ProcessStringAny(input string, mapping map[string]any) string {
	// Регулярное выражение для поиска шаблонов ${key}
	re := regexp.MustCompile(`\${([^}]*)}`)

	// Функция замены
	result := re.ReplaceAllStringFunc(input, func(match string) string {
		// Извлекаем ключ из шаблона ${key}
		key := match[2 : len(match)-1] // Убираем ${ и }

		// Ищем значение в мапе
		if value, exists := mapping[key]; exists {
			switch v := value.(type) {
			case []any:
				return JoinSeq(func(yield func(string) bool) {
					for _, i := range v {
						if !yield(i.(string)) {
							return
						}
					}
				}, ", ")
			case string:
				return v
			}
		}

		// Если ключ не найден, возвращаем оригинальный шаблон
		return match
	})

	return result
}
