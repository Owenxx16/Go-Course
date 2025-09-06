package main

import (
	"fmt"

	"github.com/Owenxx16/library"
	"github.com/Owenxx16/utils"
)

func main() {
	for {
		fmt.Println("===========CHUONG TRINH QUAN LY THU VIEN===========")
		fmt.Println("1. Them Sach")
		fmt.Println("2. Xem Danh Sach Sach")
		fmt.Println("3. Them Nguoi Muon")
		fmt.Println("4. Xem Danh Sach Nguoi Muon")
		fmt.Println("5. Muon Sach")
		fmt.Println("6. Xem Lich Su Muon")
		fmt.Println("7. Tra Sach")
		fmt.Println("8. Tim Kiem Sach")
		fmt.Println("9. Thoat")

		choice := utils.GetPositiveInt("Nhap lua chon cua ban: ")
		utils.ClearScreen()
		switch choice {
		case 1:
			fmt.Println("Chuc nang them sach")
			if err := library.AddBook(); err != nil {
				fmt.Println("Loi khi them sach:", err)
			}
		case 2:
			fmt.Println("Chuc nang xem danh sach sach")
			if err := library.ViewBooks(); err != nil {
				fmt.Println("Loi khi xem danh sach sach:", err)
			}
		case 3:
			fmt.Println("Chuc nang them nguoi muon")
			if err := library.AddBorrower(); err != nil {
				fmt.Println("Loi khi them nguoi muon:", err)
			}
		case 4:
			fmt.Println("Chuc nang xem danh sach nguoi muon")
			if err := library.ViewBorrowers(); err != nil {
				fmt.Println("Loi khi xem danh sach nguoi muon:", err)
			}
		case 5:
			fmt.Println("Chuc nang muon sach")
			if err := library.BorrowBook(); err != nil {
				fmt.Println("Loi khi muon sach:", err)
			}
		case 6:
			fmt.Println("Chuc nang xem lich su muon")
			if err := library.ViewBorrowHistory(); err != nil {
				fmt.Println("Loi khi xem lich su muon:", err)
			}
		case 7:
			fmt.Println("Chuc nang tra sach")
			if err := library.ReturnBook(); err != nil {
				fmt.Println("Loi khi tra sach:", err)
			}
		case 8:
			fmt.Println("Chuc nang tim kiem sach")
			if err := library.SearchBooks(); err != nil {
				fmt.Println("Loi khi tim kiem sach:", err)
			}
		case 9:
			fmt.Println("Thoat chuong trinh")
			return
		default:
			fmt.Println("Lua chon khong hop le, vui long thu lai.")
		}
		utils.ReadInput("Nhan Enter de tiep tuc...")
	}
}
