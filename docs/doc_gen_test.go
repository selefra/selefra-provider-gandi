package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/selefra/selefra-provider-gandi/provider"

	"github.com/selefra/selefra-provider-sdk/doc_gen"
)

func Test(t *testing.T) {

	fmt.Println("begin...")
	docOutputDirectory := os.Getenv("SELEFRA_DOC_OUTPUT_DIRECTORY")
	if docOutputDirectory == "" {
		docOutputDirectory = "./tables"
	}
	fmt.Println(docOutputDirectory)
	err := doc_gen.New(provider.GetProvider(), docOutputDirectory).Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("done...")

}
