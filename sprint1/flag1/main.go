package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

var ErrWrongNetAddressFormat = errors.New("wrong net address format")

func (n *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}

func (n *NetAddress) Set(val string) error {
	parts := strings.Split(val, ":")
	if len(parts) != 2 {
		return ErrWrongNetAddressFormat
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return ErrWrongNetAddressFormat
	}
	n.Host = parts[0]
	n.Port = port
	return nil
}

// допишите код реализации методов интерфейса
// ...

func main() {
	addr := new(NetAddress)
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)
	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
