package main
// 导入gin包
import (
	"route"
	"database"
	"fmt"
	"reflect"
)

func init(){
	users1 := []int64{1, 2, 4, 8, 16, 32, 64, 128} // 切片
	users2 := [...]int64{1, 2, 4, 8, 16, 32, 64, 128} // 数组
	users3 := [8]int64{1, 2, 4, 8, 16, 32, 64, 128} // 数组
	//users5 := make([3][int64]int64) // 数组
	//users4 := make([][int64]int64) // 切片
	users6 := make(map[int64]int64) // map
	users7 := append(users1,256)
	fmt.Println("type of len(users) : %T",reflect.TypeOf(len(users1)))
	fmt.Println("type of len(users) : %T",reflect.TypeOf(len(users2)))
	fmt.Println("type of len(users) : %T",reflect.TypeOf(len(users3)))
	//fmt.Println("type of len(users) : %T",reflect.TypeOf(users4))
	//fmt.Println("type of len(users) : %T",reflect.TypeOf(users5))
	fmt.Println("type of len(users) : %T",reflect.TypeOf(users6))
	fmt.Println("type of len(users) : %T",reflect.TypeOf(users7))
	// 数组 固定大小 类型也包括大小，
}

// 入口函数
func main() {
	db := database.GetDB()
	defer db.Close()
	router := route.InitRouter()
	router.Run(":8000")
}

// 最终编译 go build -o app



