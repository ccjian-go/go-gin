package main
// 导入gin包
import (
	"my/database"
	"my/route"
	"fmt"
	"reflect"
)

// 入口函数
func main() {
	//newUsers := make(map[int64]int64)
	users := []int64{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println("type of len(users) : %T",reflect.TypeOf(len(users)))

	db := database.GetDB()
	defer db.Close()
	router := route.InitRouter()
	router.Run(":8000")
}

// 最终编译 go build -o app



