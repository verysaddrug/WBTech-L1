package main

import "fmt"

/*
   Реализовать паттерн «адаптер» на любом примере.
*/

// Старая библиотека для перевода текста с английского на другие языки
type EnglishTranslator struct{}

func (e *EnglishTranslator) TranslateToRussian(text string) string {
	// Логика перевода с английского на русский
	return "Переведенный текст на русский: " + text
}

// Новая библиотека для перевода текста с английского на другие языки и наоборот
type NewTranslator struct{}

func (n *NewTranslator) TranslateToRussian(text string) string {
	// Логика перевода с английского на русский
	return "Переведенный текст на русский: " + text
}

func (n *NewTranslator) TranslateToEnglish(text string) string {
	// Логика перевода с русского на английский
	return "Translated text to English: " + text
}

// Адаптер для использования новой библиотеки вместо старой
type Adapter struct {
	NewTranslator *NewTranslator
}

func (a *Adapter) TranslateToRussian(text string) string {
	// Используем новую библиотеку для перевода с английского на русский
	return a.NewTranslator.TranslateToRussian(text)
}

func task21() {
	fmt.Printf("\n ============ Task 21: ============ \n\n")

	adapter := &Adapter{
		NewTranslator: &NewTranslator{},
	}

	// Используем адаптер для перевода текста с английского на русский
	translatedText := adapter.TranslateToRussian("Hello, world!")
	fmt.Println(translatedText) // Output: Переведенный текст на русский: Hello, world!

	fmt.Printf("\n ================================= \n\n")
}
