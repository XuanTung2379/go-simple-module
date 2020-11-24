package go_simple_module

import (
	"fmt"
	"log"

	"github.com/sugarme/tokenizer/pretrained"
)

// Preprocess returns preprocessed text
func Preprocess(sentence string) []string {
	tk := pretrained.BertBaseUncased()
	en, err := tk.EncodeSingle(sentence)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tokens: %q\n", en.Tokens)
	return en.Tokens
}
