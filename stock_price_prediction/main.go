package main

import (
	"fmt"
	"os/exec"
)

func runPythonScript(scriptName string) error {
	cmd := exec.Command("python3", scriptName)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing Python script %s: %w", scriptName, err)
	}
	return nil
}

func main() {
	fmt.Println("Fetching stock data...")
	if err := runPythonScript("fetch_stock_data.py"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Predicting stock prices...")
	if err := runPythonScript("predict_stock.py"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Stock price prediction complete.")
}
