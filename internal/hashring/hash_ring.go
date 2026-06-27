package hashring

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/ch-load-balancer/internal/config"
	"github.com/ch-load-balancer/internal/hashing"
	"github.com/rs/zerolog"
)

type Server struct {
	ID   string
	Host string
	Port int
}

type HashRing struct {
	Hashes     []uint64
	HashToNode map[uint64]*Server
}

func NewHashRing(cfg *config.Config, log *zerolog.Logger) *HashRing {
	ring := &HashRing{
		Hashes:     make([]uint64, 0),
		HashToNode: make(map[uint64]*Server),
	}

	VirtualNodes := cfg.VirtualNodes.Total

	for _, backend := range cfg.Backends {
		server := &Server{
			ID:   backend.Id,
			Host: backend.Host,
			Port: backend.Port,
		}
		for i := range VirtualNodes {
			hash_key := server.Host + ":" + strconv.Itoa(server.Port) + "#" + strconv.Itoa(i+1)
			hash := hashing.MurmurHash3Func(
				hash_key,
				0,
			)
			// fmt.Println(hash_key, hash)

			ring.Hashes = append(ring.Hashes, hash)
			ring.HashToNode[hash] = server
		}

	}

	slices.Sort(ring.Hashes)
	// for _, hash := range ring.Hashes {
	// 	svr, ok := ring.HashToNode[hash]
	// 	if !ok {
	// 		continue
	// 	}
	// 	fmt.Printf("svr: %v with Hash: %d\n", svr.Host+strconv.Itoa(svr.Port), hash)
	// }

	log.Info().Msgf(
		"Loaded %d servers into hash ring",
		len(ring.Hashes),
	)

	return ring

}

func (ring *HashRing) Lookup(userid string) Server {
	userIdHash := hashing.MurmurHash3Func(userid, 0)
	fmt.Printf("user-id hash: %d", userIdHash)
	l := 0
	r := len(ring.Hashes) - 1

	res := -1

	for l = 0; l <= r; {
		m := l + (r-l)/2
		if ring.Hashes[m] >= userIdHash {
			res = m
			r = m - 1
			continue
		} else {
			l = m + 1
		}
	}

	if res >= 0 {
		return *ring.HashToNode[ring.Hashes[res]]
	}

	return *ring.HashToNode[ring.Hashes[0]]
}
