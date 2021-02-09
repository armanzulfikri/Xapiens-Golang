package main

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//Mahasiswa sample
type Mahasiswa struct {
	Nama      string  `redis:"nama"`
	Nim       string  `redis:"nim"`
	IPK       float64 `redis:"ipk"`
	Semeseter int     `redis:"semester"`
}

//Main func
func main() {
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	}, 10,
	)
	pool.MaxActive = 10

	//mengambil satu koneksi
	conn := pool.Get()
	defer conn.Close()
	user := map[string]string{
		"nama": "arman",
		"kota": "balikpapan",
	}
	jd, _ := json.Marshal(user)
	fmt.Println(string(jd))

	_, _ = conn.Do("SET", "mahasiswa:1", string(jd))

	_, _ = conn.Do("EXPIRE", "mahasiswa:1", "60") //expire dalam detik

	reply, err := redis.Bytes(conn.Do("GET", "mahasiswa:1"))
	if err == nil {
		fmt.Println(reply)
		fmt.Println(string(reply))

	}
	_, _ = conn.Do("DEL", "mahasiswa:1")

}
