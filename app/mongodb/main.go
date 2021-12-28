package mongo_pool

import (
	"fmt"
	"github.com/jtzjtz/kit/conn/mongo_pool"
)

func main() {
	c, err := mongo_pool.NewClientPool("")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
