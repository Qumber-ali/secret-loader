# GoLang Azure Keyvault Secret Loader with Excelize
![Qamber](assets/secret_loader.png)

![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)

This repository contains Golang module named as "secret-loader" having two Go packages(main and azure_keyvault). This Golang module takes excel workbook as input along with the keyvault name and sheet name containing "keys" and "values" that are desired to be loaded into Azure keyvault. Excel workbook path, sheet name and keyvault name should be passed as part of command line flags to the go executable built from the source code in this repository. A step by step guide is as follows:

Clone this repository by running:
```sh
git clone https://github.com/Qumber-ali/akv-secret-loader.git
```
After cloning the repository run the following command on the root of repository to resolve dependencies and build the executable:
```golang
go get -d ./... && go build seclo.go
```
After building the executable you are ready to load secrets. Make sure you are having AKV in place on Azure and the workbook you are referencing lies on your filesystem. Also make sure the sheet name that you are providing to load secrets from exists in the provided workbook, otherwise it will throw an error.

Once all made sure. Run the following command to load the secrets from excel workbook's sheet into Azure keyvault:
```golang
./seclo --file <path to excel workbook> --sheet <Sheet name containing keys and vaules> --akv <Azure keyvault name>
```

Note you can also move this binary to path that is enrolled in "PATH" environment variable or you can append the binary's path into PATH environment variable to call the executable from anywhere on your filesystem.  
