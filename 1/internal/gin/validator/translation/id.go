package translation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterIdTranslations registers a set of indonesia translations
// for all built in tag's/custom tag's in validator; you may add your own as desired.
func RegisterIdTranslations(v *validator.Validate, trans ut.Translator) (err error) {
	translations := []translation{
		{
			tag:         "required",
			translation: "{0} wajib diisi",
			override:    false,
		},
		{
			tag:         "required_if",
			translation: "{0} wajib diisi",
			override:    false,
		},
		{
			tag:         "required_with",
			translation: "{0} wajib diisi",
			override:    false,
		},
		{
			tag:             "required_without",
			translation:     "{0} wajib diisi, jika {1} tidak diisi",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "excluded_with",
			translation:     "{0} wajib kosong, jika {1} diisi",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:         "hostname_port",
			translation: "{0} format harus hostname:port",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("len-string", "panjang {0} harus {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("len-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("len-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} harus sama dengan {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} harus berisi {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("len-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("len-items-item", "{0} item", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					c, err = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("len-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-items", extractField(fe), c)

				default:
					t, err = ut.T("len-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("min-string", "panjang minimal {0} adalah {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("min-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("min-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} harus {1} atau lebih besar", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "panjang minimal {0} adalah {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("min-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("min-items-item", "{0} item", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-items", extractField(fe), c)

				default:
					t, err = ut.T("min-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("max-string", "panjang maksimal {0} adalah {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("max-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("max-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} harus {1} atau kurang", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} harus berisi maksimal {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("max-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("max-items-item", "{0} item", locales.PluralRuleOther, false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-items", extractField(fe), c)

				default:
					t, err = ut.T("max-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:             "eq",
			translation:     "{0} tidak sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "ne",
			translation:     "{0} tidak sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lt-string", "panjang {0} harus kurang dari {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lt-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lt-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} harus kurang dari {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} harus berisi kurang dari {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lt-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lt-items-item", "{0} item", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} harus kurang dari tanggal & waktu saat ini", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-items", extractField(fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lt-datetime", extractField(fe))

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("lte-string", "panjang maksimal {0} adalah {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lte-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lte-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} harus {1} atau kurang", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} harus berisi maksimal {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("lte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("lte-items-item", "{0} item", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} harus kurang dari atau sama dengan tanggal & waktu saat ini", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-items", extractField(fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lte-datetime", extractField(fe))

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gt-string", "panjang {0} harus lebih dari {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gt-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gt-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} harus lebih besar dari {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} harus berisi lebih dari {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gt-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gt-items-item", "{0} item", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} harus lebih besar dari tanggal & waktu saat ini", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-items", extractField(fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gt-datetime", extractField(fe))

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {
				if err = ut.Add("gte-string", "panjang minimal {0} adalah {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gte-string-character", "{0} karakter", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gte-string-character", "{0} karakter", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} harus {1} atau lebih besar", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} harus berisi setidaknya {1}", false); err != nil {
					return
				}

				// if err = ut.AddCardinal("gte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				// 	return
				// }

				if err = ut.AddCardinal("gte-items-item", "{0} item", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} harus lebih besar dari atau sama dengan tanggal & waktu saat ini", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {
					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-string", extractField(fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-items", extractField(fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gte-datetime", extractField(fe))

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-number", extractField(fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:             "eqfield",
			translation:     "{0} harus sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "eqcsfield",
			translation:     "{0} harus sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "necsfield",
			translation:     "{0} tidak sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "gtcsfield",
			translation:     "{0} harus lebih besar dari {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "gtecsfield",
			translation:     "{0} harus lebih besar dari atau sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "ltcsfield",
			translation:     "{0} harus kurang dari {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "ltecsfield",
			translation:     "{0} harus kurang dari atau sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "nefield",
			translation:     "{0} tidak sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "gtfield",
			translation:     "{0} harus lebih besar dari {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "gtefield",
			translation:     "{0} harus lebih besar dari atau sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "ltfield",
			translation:     "{0} harus kurang dari {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:             "ltefield",
			translation:     "{0} harus kurang dari atau sama dengan {1}",
			override:        false,
			customTransFunc: translateFuncFieldComparison,
		},
		{
			tag:         "alpha",
			translation: "{0} hanya dapat berisi karakter abjad",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} hanya dapat berisi karakter alfanumerik",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} harus berupa nilai numerik yang valid",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} harus berupa angka yang valid",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} harus berupa heksadesimal yang valid",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} harus berupa warna HEX yang valid",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} harus berupa warna RGB yang valid",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} harus berupa warna RGBA yang valid",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} harus berupa warna HSL yang valid",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} harus berupa warna HSLA yang valid",
			override:    false,
		},
		{
			tag:         "e164",
			translation: "{0} tidak memenuhi format nomor telepon +62 (XXX) XXXX-XXXX",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} harus berupa alamat email yang valid",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} harus berupa URL yang valid",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} harus berupa URI yang valid",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} harus berupa string Base64 yang valid",
			override:    false,
		},
		{
			tag:             "contains",
			translation:     "{0} harus berisi teks '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "containsany",
			translation:     "{0} harus berisi setidaknya salah satu karakter berikut '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "excludes",
			translation:     "{0} tidak boleh berisi teks '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "excludesall",
			translation:     "{0} tidak boleh berisi salah satu karakter berikut '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "excludesrune",
			translation:     "{0} tidak boleh berisi '{1}'",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:         "isbn",
			translation: "{0} harus berupa nomor ISBN yang valid",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} harus berupa nomor ISBN-10 yang valid",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} harus berupa nomor ISBN-13 yang valid",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} harus berupa UUID yang valid",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} harus berupa UUID versi 3 yang valid",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} harus berupa UUID versi 4 yang valid",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} harus berupa UUID versi 5 yang valid",
			override:    false,
		},
		{
			tag:         "ulid",
			translation: "{0} harus berupa ULID yang valid",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} hanya boleh berisi karakter ascii",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} hanya boleh berisi karakter ascii yang dapat dicetak",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} harus berisi karakter multibyte",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} harus berisi URI Data yang valid",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} harus berisi koordinat lintang yang valid",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} harus berisi koordinat bujur yang valid",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} harus berupa nomor SSN yang valid",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} harus berupa alamat IPv4 yang valid",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} harus berupa alamat IPv6 yang valid",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} harus berupa alamat IP yang valid",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} harus berisi notasi CIDR yang valid",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} harus berisi notasi CIDR yang valid untuk alamat IPv4",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} harus berisi notasi CIDR yang valid untuk alamat IPv6",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} harus berupa alamat TCP yang valid",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} harus berupa alamat TCP IPv4 yang valid",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} harus berupa alamat TCP IPv6 yang valid",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} harus berupa alamat UDP yang valid",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} harus berupa alamat IPv4 UDP yang valid",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} harus berupa alamat IPv6 UDP yang valid",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} harus berupa alamat IP yang dapat dipecahkan",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} harus berupa alamat IPv4 yang dapat diatasi",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} harus berupa alamat IPv6 yang dapat diatasi",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} harus berupa alamat UNIX yang dapat diatasi",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} harus berisi alamat MAC yang valid",
			override:    false,
		},
		{
			tag:         "unique",
			translation: "{0} harus unik",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} harus berupa warna yang valid",
			override:    false,
		},
		{
			tag:             "oneof",
			translation:     "{0} harus berupa salah satu dari [{1}]",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:         "json",
			translation: "{0} harus berupa json string yang valid",
			override:    false,
		},
		{
			tag:         "jwt",
			translation: "{0} harus berupa jwt string yang valid",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} harus berupa huruf kecil",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} harus berupa huruf besar",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} tidak memenuhi format {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, err := ut.T(fe.Tag(), extractField(fe), extractDatetimeParam(fe))
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:             "postcode_iso3166_alpha2",
			translation:     "{0} tidak memenuhi format kode pos dari negara {1}",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:             "postcode_iso3166_alpha2_field",
			translation:     "{0} tidak memenuhi format kode pos dari negara yang berada di kolom {1}",
			override:        false,
			customTransFunc: translateFuncValueComparison,
		},
		{
			tag:         "boolean",
			translation: "{0} harus berupa boolean yang valid",
			override:    false,
		},
	}

	for _, t := range translations {
		err = registerTranslation(v, trans, t)

		if err != nil {
			panic(err)
		}
	}

	return
}
