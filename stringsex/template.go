package stringsex

import (
	"regexp"
)

// ProcessStringAny обрабатывает строку, заменяя шаблоны ${key} на значения из mapping
func ProcessStringAny(input string, mapping map[string]any) (result string, ok bool) {
	// Регулярное выражение для поиска шаблонов ${key}
	re := regexp.MustCompile(`\${([^}]*)}`)

	var toString func(v any) string
	toString = func(v any) string {
		switch v := v.(type) {
		case []any:
			return "{" + JoinSeq(func(yield func(string) bool) {
				for _, i := range v {
					if !yield(toString(i)) {
						return
					}
				}
			}, ", ") + "}"
		case string:
			return v
		}

		return ""
	}

	// Функция замены
	ok = true
	result = re.ReplaceAllStringFunc(input, func(match string) string {
		// Извлекаем ключ из шаблона ${key}
		key := match[2 : len(match)-1] // Убираем ${ и }

		// Ищем значение в мапе
		if value, exists := mapping[key]; exists {
			return toString(value)
		}

		// Если ключ не найден, возвращаем оригинальный шаблон
		ok = false
		return match
	})

	return
}
