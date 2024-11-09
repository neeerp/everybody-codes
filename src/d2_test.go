package main

import "testing"

func TestRunicWordsExample1(t *testing.T) {
	words := "THE,OWE,MES,ROD,HER"
  runes := "AWAKEN THE POWER ADORNED WITH THE FLAMES BRIGHT IRE"

	expected := 4

	result := countRunicWords(words, runes)

	if expected != result {
		t.Errorf("countRunicWords(\"%s\", \"%s\") = %d; want %d", words, runes, result, expected)
	}
}

func TestRunicWordsExample2(t *testing.T) {
	words := "THE,OWE,MES,ROD,HER"
  runes := "THE FLAME SHIELDED THE HEART OF THE KINGS"

	expected := 3

	result := countRunicWords(words, runes)

	if expected != result {
		t.Errorf("countRunicWords(\"%s\", \"%s\") = %d; want %d", words, runes, result, expected)
	}
}

func TestRunicWordsExample3(t *testing.T) {
	words := "THE,OWE,MES,ROD,HER"
  runes := "POWE PO WER P OWE R"

	expected := 2

	result := countRunicWords(words, runes)

	if expected != result {
		t.Errorf("countRunicWords(\"%s\", \"%s\") = %d; want %d", words, runes, result, expected)
	}
}

func TestRunicWordsExample4(t *testing.T) {
	words := "THE,OWE,MES,ROD,HER"
  runes := "THERE IS THE END"

	expected := 3

	result := countRunicWords(words, runes)

	if expected != result {
		t.Errorf("countRunicWords(\"%s\", \"%s\") = %d; want %d", words, runes, result, expected)
	}
}

func TestRunicWordsActual(t *testing.T) {
  words := "LOR,LL,SI,OR,RE,NA,AB"
  runes := "LOREM IPSUM DOLOR SIT AMET, CONSECTETUR ADIPISCING ELIT, SED DO EIUSMOD TEMPOR INCIDIDUNT UT LABORE ET DOLORE MAGNA ALIQUA. UT ENIM AD MINIM VENIAM, QUIS NOSTRUD EXERCITATION ULLAMCO LABORIS NISI UT ALIQUIP EX EA COMMODO CONSEQUAT. DUIS AUTE IRURE DOLOR IN REPREHENDERIT IN VOLUPTATE VELIT ESSE CILLUM DOLORE EU FUGIAT NULLA PARIATUR. EXCEPTEUR SINT OCCAECAT CUPIDATAT NON PROIDENT, SUNT IN CULPA QUI OFFICIA DESERUNT MOLLIT ANIM ID EST LABORUM."

	expected := 32

	result := countRunicWords(words, runes)

	if expected != result {
		t.Errorf("countRunicWords(\"%s\", \"%s\") = %d; want %d", words, runes, result, expected)
	}
}
