// Default Code Interface for named interfaces
type Copier interface {
	Copy(string, string) int
}

// Bởi vì chúng ta đang định nghĩa Copy Method mà không đặt tên cho các argument , bởi vì thế sẽ không thể xác định được các argument này là gì.
// Do đó, chúng ta sẽ không thể sử dụng các argument này trong phần thân của hàm Copy.
// Để giải quyết vấn đề này, chúng ta sẽ đặt tên cho các argument trong hàm Copy.

type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

// Lúc này nó sẽ dề hiểu và rõ ràng hơn