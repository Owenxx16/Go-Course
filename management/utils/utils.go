package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveInt(prompt string) int {
	for {
		inout := ReadInput(prompt)
		value, err := strconv.Atoi(inout)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Vui long nhap so nguyen duong.")
	}
}

func GetPositiveFloat(prompt string) float64 {
	for {
		inout := ReadInput(prompt)
		value, err := strconv.ParseFloat(inout, 64)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Vui long nhap so nguyen duong.")
	}
}

func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}
		fmt.Println("Vui long nhap mot chuoi khong rong.")
	}
}

func GetOptionalString(prompt string, defaultValue string) string {
	input := ReadInput(prompt)
	if input == "" {
		return defaultValue
	}
	return input
}

func GetOptionalPositiveFloat(prompt string, defaultValue float64) float64 {
	input := ReadInput(prompt)
	if input == "" {
		return defaultValue
	}
	value, err := strconv.ParseFloat(input, 64)
	if err == nil && value > 0 {
		return value
	}

	return value
}

func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("X", err.Error())
	}
}

type HasId interface {
	GetId() int
}

func CheckIdDuplicate[T HasId](id int, list []T) bool {
	for _, item := range list {
		if item.GetId() == id {
			return false
		}
	}
	return true
}
