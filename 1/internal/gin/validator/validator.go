// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package validator

import (
	"reflect"
	"strings"
	"sync"

	"myapp/internal/gin/validator/translation"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

const (
	requiredIfStringerTag = "required_if_stringer"
)

var (
	DefaultTranslator ut.Translator
	Translators       map[string]ut.Translator

	enTranslator ut.Translator
	idTranslator ut.Translator
)

func init() {
	en := en.New()
	id := id.New()

	uni := ut.New(en, id)

	enTranslator, _ = uni.GetTranslator(en.Locale())
	idTranslator, _ = uni.GetTranslator(id.Locale())

	DefaultTranslator = enTranslator

	Translators = map[string]ut.Translator{}
	Translators[en.Locale()] = enTranslator
	Translators[id.Locale()] = idTranslator

	overrideGinBindingValidator()
}

func overrideGinBindingValidator() {
	binding.Validator = New()
}

type Validator interface {
	binding.StructValidator

	ValidateVar(interface{}, string) error
}

func New() Validator {
	return &defaultValidator{}
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

type StructValidationErrors struct {
	ve       validator.ValidationErrors
	structNs string
}

func (ve StructValidationErrors) Error() string {
	return ve.ve.Error()
}

func (ve StructValidationErrors) Translate(translator ut.Translator) map[string]string {
	translations := map[string]string{}
	for _, err := range ve.ve {
		translations[ve.extractDomain(err)] = err.Translate(translator)
	}

	return translations
}

func (ve StructValidationErrors) extractDomain(err validator.FieldError) string {
	var (
		sep                   = "."
		domainChunks          = strings.Split(err.StructNamespace(), sep)
		formattedDomainChunks = []string{}
	)

	for _, domainChunk := range domainChunks {
		if domainChunk == ve.structNs || domainChunk == "" {
			continue
		}

		formattedDomainChunks = append(formattedDomainChunks, strcase.ToSnake(domainChunk))
	}

	domain := strings.Join(formattedDomainChunks, sep)

	switch domain {
	case "pagination_request.limit":
		domain = "limit"
	case "pagination_request.page":
		domain = "page"
	}

	return domain
}

type VarValidationErrors struct {
	ve validator.ValidationErrors
}

func (ve VarValidationErrors) Error() string {
	return ve.ve.Error()
}

func (ve VarValidationErrors) Translate(translator ut.Translator) string {
	for _, err := range ve.ve {
		return err.Translate(translator)
	}

	return ""
}

var (
	_ Validator               = &defaultValidator{}
	_ binding.StructValidator = &defaultValidator{}
)

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		return v.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return v.validateStruct(obj)
	default:
		return nil
	}
}

// validateStruct receives struct type
func (v *defaultValidator) validateStruct(obj interface{}) error {
	v.lazyinit()

	err := v.validate.Struct(obj)
	if err != nil {
		ve, ok := err.(validator.ValidationErrors)
		if ok {
			return StructValidationErrors{
				ve:       ve,
				structNs: extractStructNamespace(obj),
			}
		}
	}

	return err
}

func (v *defaultValidator) ValidateVar(field interface{}, tag string) error {
	v.lazyinit()

	err := v.validate.Var(field, tag)
	if err != nil {
		ve, ok := err.(validator.ValidationErrors)
		if ok {
			err = VarValidationErrors{
				ve: ve,
			}
		}
	}

	return err
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://pkg.go.dev/github.com/go-playground/validator/v10
func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		translation.RegisterEnTranslations(v.validate, enTranslator)
		translation.RegisterIdTranslations(v.validate, idTranslator)
	})
}

func extractStructNamespace(obj interface{}) string {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	return val.Type().Name()
}
