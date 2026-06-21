package hashring

import (
	"slices"
	"strconv"

	"github.com/ch-load-balancer/internal/config"
	"github.com/ch-load-balancer/internal/hashing"
	"github.com/rs/zerolog"
)

type Server struct {
	ID   int
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

	for _, backend := range cfg.Backends {
		server := &Server{
			ID:   backend.Id,
			Host: backend.Host,
			Port: backend.Port,
		}

		hash := hashing.Fnv1aHashFunc(
			server.Host + ":" + strconv.Itoa(server.Port),
		)

		ring.Hashes = append(ring.Hashes, hash)
		ring.HashToNode[hash] = server
	}

	slices.Sort(ring.Hashes)

	log.Info().Msgf(
		"Loaded %d servers into hash ring",
		len(ring.Hashes),
	)

	return ring

}

func (ring *HashRing) Lookup(userid string) Server {
	userIdHash := hashing.Fnv1aHashFunc(userid)

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
