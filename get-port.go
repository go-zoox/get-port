package getport

import "net"

// MustGetPort is like GetPort but panics if an error occurs.
func MustGetPort() int {
	port, err := GetFreePort()
	if err != nil {
		panic(err)
	}

	return port
}

// GetFreePort returns a free port that is ready to use.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// GetFreePorts returns n free ports that are ready to use.
func GetFreePorts(count int) ([]int, error) {
	var ports []int
	for i := 0; i < count; i++ {
		port, err := GetFreePort()
		if err != nil {
			return nil, err
		}

		ports = append(ports, port)
	}

	return ports, nil
}
