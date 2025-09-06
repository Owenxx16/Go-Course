package main

import "fmt"

type rectangular struct {
	chieudai  float32 `desc: "Chiều dài"`
	chieurong float32 `desc: "Chiều rộng"`
}

func (r rectangular) chuvi() float32 {
	return 2 * (r.chieudai + r.chieurong)
}

func (r rectangular) dientich() float32 {
	return r.chieudai * r.chieurong
}

func main() {
	r := rectangular{}
	fmt.Println("Vui lòng nhập chiều dài")
	fmt.Scan(&r.chieudai)
	fmt.Println("Vui lòng nhập chiều rộng")
	fmt.Scan(&r.chieurong)

	fmt.Printf("Chu vi hình chữ nhật: %.2f ", r.chuvi())
	fmt.Println()
	fmt.Printf("Diện tích hình chữ nhật: %.2f", r.dientich())
}
