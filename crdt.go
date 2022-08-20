package consistency

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Rock-liyi/p2pdb/domain/log/entity"
)

type Consistency interface {
	Compare(mc *MerKleDag)
	Merge(mc *MerKleDag)
}

type MerKleDag struct {
	Nodes map[string]*entity.NodeEntity
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(nodes map[string]*entity.NodeEntity) (*MerKleDag) {
	return &MerKleDag{Nodes: nodes}
}

// Compare with other diff nodes
func (rs *MerKleDag) Compare(mc *MerKleDag) {

}

func (rs *MerKleDag) Merge(mc *MerKleDag) {

}




