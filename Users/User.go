package users

import (
	"fmt"
	"golang/modelos"
	"time"
)

func AltUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "David", time.Now(), true)
	fmt.Println(u)
}
