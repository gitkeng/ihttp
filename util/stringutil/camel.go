package stringutil

//thank https://github.com/iancoleman/strcase

import "strings"

// s := "AnyKind of_string"
// ToSnake(s)								any_kind_of_string
// ToSnakeWithIgnore(s, '.')				any_kind.of_string
// ToScreamingSnake(s)						ANY_KIND_OF_STRING
// ToKebab(s)								any-kind-of-string
// ToScreamingKebab(s)						ANY-KIND-OF-STRING
// ToDelimited(s, '.')						any.kind.of.string
// ToScreamingDelimited(s, '.', ”, true)	ANY.KIND.OF.STRING
// ToScreamingDelimited(s, '.', ' ', true)	ANY.KIND OF.STRING
// ToCamel(s)								AnyKindOfString
// ToLowerCamel(s)							anyKindOfString

var uppercaseAcronym = map[string]string{
	"ID": "id",
}

// ConfigureAcronym allows you to add additional words which will be considered acronyms
func ConfigureAcronym(key, val string) {
	uppercaseAcronym[key] = val
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	if a, ok := uppercaseAcronym[s]; ok {
		s = a
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}

// ToCamel converts a string to CamelCase
func ToCamel(s string) string {
	return toCamelInitCase(s, true)
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	return toCamelInitCase(s, false)
}
