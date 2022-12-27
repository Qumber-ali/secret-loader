package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	awssm "secret-loader/providers/aws"
	azkv "secret-loader/providers/azure"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	file_path := flag.String("file", "", "path to xlxs file.")
	sheet_name := flag.String("sheet", "", "name of the sheet in workbook.")
	provider := flag.String("provider", "", "name of the cloud provider.")
	secret_manager_name := flag.String("awssm", "", "name of the aws secrets manager instance.")
	aws_profile := flag.String("profile", "default", "name of the aws profile to load config and credentials from.")
	vault_name := flag.String("akv", "", "name of the akv.")

	flag.Parse()

	//var vault_name, aws_profile *string

	switch *provider {
	case "aws":
		if *secret_manager_name == "" {
			fmt.Fprintf(os.Stderr, "Error: you have provided aws provider but not provided awssm flag containing secret manager instance name.")
			os.Exit(1)
		}
	case "azure":
		if *vault_name == "" {
			fmt.Fprintf(os.Stderr, "Error: you have provided azure provider but not provided akv flag containing keyvault name.")
			os.Exit(1)
		}
	}

	flag.Parse()

	var f *excelize.File

	if path.IsAbs(*file_path) == true {
		f, err = excelize.OpenFile(*file_path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	} else {
		f, err = excelize.OpenFile(cwd + "/" + *file_path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

	index := f.GetSheetIndex(*sheet_name)

	if index == -1 {
		fmt.Println("Error: Given sheet does not exist in excel workbook provided.")
		os.Exit(1)
	} else {
		fmt.Println("The given sheet exists in the given excel workbook.")
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	cols, err := f.GetCols(*sheet_name)

	var key_flag bool = false

	keys, key_flag, err := CreateKeyValue(cols, key_flag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	values, key_flag, err := CreateKeyValue(cols, key_flag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch *provider {
	case "aws":
		if *aws_profile != "default" {
			awssm.LoadSecrets(*aws_profile, *secret_manager_name, keys, values)
		} else {
			awssm.LoadSecrets("default", *secret_manager_name, keys, values)
		}
	case "azure":
		azkv.LoadSecrets(*vault_name, keys, values)
	}

}

func CreateKeyValue(cols [][]string, key_flag bool) ([]string, bool, error) {

	for col_index, col := range cols {
		for row_index, row := range col {
			if strings.EqualFold(row, "key") == true && key_flag == false {
				keys := cols[col_index][row_index+1 : len(col)-1]
				key_flag = true
				return keys, key_flag, nil
			} else {
				if strings.EqualFold(row, "value") == true {
					values := cols[col_index][row_index+1 : len(col)-1]
					key_flag = true
					return values, key_flag, nil
				}
			}
		}
	}

	var slice []string = []string{}
	return slice, key_flag, errors.New("key or value as column title didn't exist")

}
