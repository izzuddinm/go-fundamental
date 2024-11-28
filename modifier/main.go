package modifier

func CountTotalMessage(page int, pageSize int) int {
	return page + pageSize
}

func substractTotalMessage(page int, pageSize int) int {
	return pageSize - page
}

var version string = "v.1.0"
var Application string = "Jarvis-Backend-Application"
