package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/simcord-llc/bitbon-system-blockchain/common"
	"github.com/simcord-llc/bitbon-system-blockchain/core/types"
	"github.com/simcord-llc/bitbon-system-blockchain/crypto"
	"github.com/simcord-llc/bitbon-system-blockchain/ethclient"
	"github.com/simcord-llc/bitbon-system-blockchain/params"
)

// Edit Test Configuration Here
var (
	configDevCpos = Config{
		intensity: 1, // tx per sec per node (i.e. if intensity is 25 tx per sec and there are 4 nodes, the overall load would be 100 tx per sec)
		duration:  30 * time.Minute,
		nodes: []NodeConfig{
			//{"46:10555-00", "ws://172.16.28.46:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c000"},
			//{"46:10556-02", "ws://172.16.28.46:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c002"},
			//{"45:10555-03", "ws://172.16.28.45:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c003"},
			//{"45:10556-05", "ws://172.16.28.45:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c005"},
			//{"46:10555-10", "ws://172.16.28.46:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c010"},
			//{"46:10556-12", "ws://172.16.28.46:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c012"},
			//{"45:10555-13", "ws://172.16.28.45:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c013"},
			//{"45:10556-15", "ws://172.16.28.45:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c015"},
			//{"46:10555-20", "ws://172.16.28.46:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c020"},
			//{"46:10556-22", "ws://172.16.28.46:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c022"},
			//{"45:10555-23", "ws://172.16.28.45:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c023"},
			//{"45:10556-25", "ws://172.16.28.45:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c025"},
			{"46:10555-30", "ws://172.16.28.46:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c030"},
			//{"46:10556-32", "ws://172.16.28.46:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c032"},
			//{"45:10555-33", "ws://172.16.28.45:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c033"},
			//{"45:10556-35", "ws://172.16.28.45:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c035"},
			//{"46:10555-40", "ws://172.16.28.46:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c040"},
			//{"46:10556-42", "ws://172.16.28.46:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c042"},
			//{"45:10555-43", "ws://172.16.28.45:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c043"},
			//{"45:10556-45", "ws://172.16.28.45:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c045"},
			//{"46:10555-50", "ws://172.16.28.46:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c050"},
			//{"46:10556-52", "ws://172.16.28.46:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c052"},
			//{"45:10555-53", "ws://172.16.28.45:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c053"},
			//{"45:10556-55", "ws://172.16.28.45:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c055"},
			//{"46:10555-60", "ws://172.16.28.46:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c060"},
			//{"46:10556-62", "ws://172.16.28.46:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c062"},
			//{"45:10555-63", "ws://172.16.28.45:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c063"},
			//{"45:10556-65", "ws://172.16.28.45:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c065"},
		},
	}
	configVPS = Config{
		intensity: 1, // tx per sec per node (i.e. if intensity is 25 tx per sec and there are 4 nodes, the overall load would be 100 tx per sec)
		//intensity: 1,
		duration: 5 * time.Minute,
		nodes: []NodeConfig{
			{"35:10555-00", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c000"},
			//{"35:10556-02", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c002"},
			//{"67:10555-03", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c003"},
			//{"68:10556-05", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c005"},
			//{"35:10555-10", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c010"},
			//{"35:10556-12", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c012"},
			//{"67:10555-13", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c013"},
			//{"68:10556-15", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c015"},
			//{"35:10555-20", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c020"},
			//{"35:10556-22", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c022"},
			//{"67:10555-23", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c023"},
			//{"68:10556-25", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c025"},
			//{"35:10555-30", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c030"},
			//{"35:10556-32", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c032"},
			//{"67:10555-33", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c033"},
			//{"68:10556-35", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c035"},
			//{"35:10555-40", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c040"},
			//{"35:10556-42", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c042"},
			//{"67:10555-43", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c043"},
			//{"68:10556-45", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c045"},
			//{"35:10555-50", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c050"},
			//{"35:10556-52", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c052"},
			//{"67:10555-53", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c053"},
			//{"68:10556-55", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c055"},
			//{"35:10555-60", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c060"},
			//{"35:10556-62", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c062"},
			//{"67:10555-63", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c063"},
			//{"68:10556-65", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c065"},
		},
	}
	configLocal = Config{
		intensity: 1, // tx per sec per node (i.e. if intensity is 25 tx per sec and there are 4 nodes, the overall load would be 100 tx per sec)
		//intensity: 1,
		duration: 21 * time.Minute,
		nodes: []NodeConfig{
			{"35:10555-00", "http://localhost:8545", "9e2d0a29db78ca225291da40ee1f5cd90b55af6d894ee784214ea51ab8e9fd8b"},
			//{"35:10556-02", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c002"},
			//{"67:10555-03", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c003"},
			//{"68:10556-05", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c005"},
			//{"35:10555-10", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c010"},
			//{"35:10556-12", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c012"},
			//{"67:10555-13", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c013"},
			//{"68:10556-15", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c015"},
			//{"35:10555-20", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c020"},
			//{"35:10556-22", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c022"},
			//{"67:10555-23", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c023"},
			//{"68:10556-25", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c025"},
			//{"35:10555-30", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c030"},
			//{"35:10556-32", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c032"},
			//{"67:10555-33", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c033"},
			//{"68:10556-35", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c035"},
			//{"35:10555-40", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c040"},
			//{"35:10556-42", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c042"},
			//{"67:10555-43", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c043"},
			//{"68:10556-45", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c045"},
			//{"35:10555-50", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c050"},
			//{"35:10556-52", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c052"},
			//{"67:10555-53", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c053"},
			//{"68:10556-55", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c055"},
			//{"35:10555-60", "ws://172.16.37.35:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c060"},
			//{"35:10556-62", "ws://172.16.37.35:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c062"},
			//{"67:10555-63", "ws://172.16.37.67:10555", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c063"},
			//{"68:10556-65", "ws://172.16.37.68:10556", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c065"},
		},
	}
	configGoodDEV = Config{
		intensity: 5, // tx per sec per node (i.e. if intensity is 25 tx per sec and there are 4 nodes, the overall load would be 100 tx per sec)
		duration:  5 * time.Minute,
		nodes: []NodeConfig{
			{"21:10511-000", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c000"},
			{"21:10512-002", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c002"},
			{"21:10513-003", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c003"},
			{"21:10514-005", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c005"},
			{"21:10515-010", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c010"},
			{"21:10516-012", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c012"},
			{"21:10517-013", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c013"},
			{"21:10511-015", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c015"},
			{"21:10512-020", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c020"},
			{"21:10513-022", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c022"},
			{"21:10514-023", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c023"},
			{"21:10515-025", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c025"},
			{"21:10516-030", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c030"},
			{"21:10517-032", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c032"},
			{"21:10511-033", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c033"},
			{"21:10512-035", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c035"},
			{"21:10513-040", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c040"},
			{"21:10514-042", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c042"},
			{"21:10515-043", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c043"},
			{"21:10516-045", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c045"},
			{"21:10517-050", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c050"},
			{"21:10511-052", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c052"},
			{"21:10512-053", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c053"},
			{"21:10513-055", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c055"},
			{"21:10514-060", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c060"},
			{"21:10515-062", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c062"},
			{"21:10516-063", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c063"},
			{"21:10517-065", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c065"},
			{"21:10512-066", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c066"},
			{"21:10513-067", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c067"},
			{"21:10514-068", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c068"},
			{"21:10515-069", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c069"},
			{"21:10516-070", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c070"},
			{"21:10517-071", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c071"},
			{"21:10511-072", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c072"},
			{"21:10512-073", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c073"},
			{"21:10513-074", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c074"},
			{"21:10514-075", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c075"},
			{"21:10515-076", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c076"},
			{"21:10516-077", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c077"},
			{"21:10517-078", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c078"},
			{"21:10511-079", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c079"},
			{"21:10512-080", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c080"},
			{"21:10513-081", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c081"},
			{"21:10514-082", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c082"},
			{"21:10515-083", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c083"},
			{"21:10516-084", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c084"},
			{"21:10517-085", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c085"},
			{"21:10511-086", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c086"},
			{"21:10512-087", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c087"},
			{"21:10513-088", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c088"},
			{"21:10514-089", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c089"},
			{"21:10515-090", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c090"},
			{"21:10516-091", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c091"},
			{"21:10517-092", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c092"},
			{"21:10511-093", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c093"},
			{"21:10512-094", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c094"},
			{"21:10513-095", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c095"},
			{"21:10514-096", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c096"},
			{"21:10515-097", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c097"},
			{"21:10516-098", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c098"},
			{"21:10517-099", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c099"},
			{"21:10511-100", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c100"},
			{"21:10512-101", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c101"},
			{"21:10513-102", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c102"},
			{"21:10514-103", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c103"},
			{"21:10515-104", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c104"},
			{"21:10516-105", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c105"},
			{"21:10517-106", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c106"},
			{"21:10511-107", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c107"},
			{"21:10512-108", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c108"},
			{"21:10513-109", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c109"},
			{"21:10514-110", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c110"},
			{"21:10515-111", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c111"},
			{"21:10516-112", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c112"},
			{"21:10517-113", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c113"},
			{"21:10511-114", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c114"},
			{"21:10512-115", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c115"},
			{"21:10513-116", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c116"},
			{"21:10514-117", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c117"},
			{"21:10515-118", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c118"},
			{"21:10516-119", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c119"},
			{"21:10517-120", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c120"},
			{"21:10511-121", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c121"},
			{"21:10512-122", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c122"},
			{"21:10513-123", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c123"},
			{"21:10514-124", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c124"},
			{"21:10515-125", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c125"},
			{"21:10516-126", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c126"},
			{"21:10517-127", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c127"},
			{"21:10511-128", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c128"},
			{"21:10512-129", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c129"},
			{"21:10513-130", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c130"},
			{"21:10514-131", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c131"},
			{"21:10515-132", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c132"},
			{"21:10516-133", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c133"},
			{"21:10517-134", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c134"},
			{"21:10511-135", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c135"},
			{"21:10512-136", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c136"},
			{"21:10513-137", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c137"},
			{"21:10514-138", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c138"},
			{"21:10515-139", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c139"},
			{"21:10516-140", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c140"},
			{"21:10517-141", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c141"},
			{"21:10511-142", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c142"},
			{"21:10512-143", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c143"},
			{"21:10513-144", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c144"},
			{"21:10514-145", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c145"},
			{"21:10515-146", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c146"},
			{"21:10516-147", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c147"},
			{"21:10517-148", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c148"},
			{"21:10511-149", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c149"},
			{"21:10512-150", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c150"},
			{"21:10513-151", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c151"},
			{"21:10514-152", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c152"},
			{"21:10515-153", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c153"},
			{"21:10516-154", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c154"},
			{"21:10517-155", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c155"},
			{"21:10511-156", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c156"},
			{"21:10512-157", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c157"},
			{"21:10513-158", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c158"},
			{"21:10514-159", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c159"},
			{"21:10515-160", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c160"},
			{"21:10516-161", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c161"},
			{"21:10517-162", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c162"},
			{"21:10511-163", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c163"},
			{"21:10512-164", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c164"},
			{"21:10513-165", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c165"},
			{"21:10514-166", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c166"},
			{"21:10515-167", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c167"},
			{"21:10516-168", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c168"},
			{"21:10517-169", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c169"},
			{"21:10511-170", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c170"},
			{"21:10512-171", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c171"},
			{"21:10513-172", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c172"},
			{"21:10514-173", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c173"},
			{"21:10515-174", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c174"},
			{"21:10516-175", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c175"},
			{"21:10517-176", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c176"},
			{"21:10511-177", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c177"},
			{"21:10512-178", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c178"},
			{"21:10513-179", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c179"},
			{"21:10514-180", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c180"},
			{"21:10515-181", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c181"},
			{"21:10516-182", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c182"},
			{"21:10517-183", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c183"},
			{"21:10511-184", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c184"},
			{"21:10512-185", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c185"},
			{"21:10513-186", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c186"},
			{"21:10514-187", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c187"},
			{"21:10515-188", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c188"},
			{"21:10516-189", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c189"},
			{"21:10517-190", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c190"},
			{"21:10511-191", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c191"},
			{"21:10512-192", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c192"},
			{"21:10513-193", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c193"},
			{"21:10514-194", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c194"},
			{"21:10515-195", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c195"},
			{"21:10516-196", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c196"},
			{"21:10517-197", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c197"},
			{"21:10511-198", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c198"},
			{"21:10512-199", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c199"},
			{"21:10513-200", "ws://192.168.253.21:10513", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c200"},
			{"21:10514-201", "ws://192.168.253.21:10514", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c201"},
			{"21:10515-202", "ws://192.168.253.21:10515", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c202"},
			{"21:10516-203", "ws://192.168.253.21:10516", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c203"},
			{"21:10517-204", "ws://192.168.253.21:10517", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c204"},
			{"21:10511-205", "ws://192.168.253.21:10518", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c205"},
			{"21:10512-206", "ws://192.168.253.21:10519", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c206"},
			{"21:10513-207", "ws://192.168.253.21:10510", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c207"},
			{"21:10514-208", "ws://192.168.253.21:10511", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c208"},
			{"21:10515-209", "ws://192.168.253.21:10512", "fffbd6067cb4612df85d9c10085cc47f259ced4d4cd99816b14f35650f59c209"},
		},
	}

	config = configLocal
)

const (
	baseValue       = 1 * params.Wei
	defaultGasLimit = 21000
	defaultGasPrice = 5 * params.Wei
	chainId         = 14888
)

// Config and helper structs
type NodeConfig struct {
	name       string
	uri        string
	privateKey string
}

type Config struct {
	nodes     []NodeConfig
	intensity uint64
	duration  time.Duration
	debug     bool
}

type Node struct {
	config     NodeConfig
	client     *ethclient.Client
	nonce      uint64
	address    common.Address
	privateKey *ecdsa.PrivateKey
	startBlock uint64
	endBlock   uint64

	startBlockMux sync.RWMutex
	endBlockMux   sync.RWMutex
	nonceMux      sync.RWMutex
}

func (n *Node) EndBlock() uint64 {
	n.endBlockMux.RLock()
	defer n.endBlockMux.RLock()
	return n.endBlock
}

func (n *Node) SetEndBlock(endBlock uint64) {
	n.endBlockMux.Lock()
	defer n.endBlockMux.Unlock()
	n.endBlock = endBlock
}

func (n *Node) StartBlock() uint64 {
	n.startBlockMux.RLock()
	defer n.startBlockMux.RUnlock()
	return n.startBlock
}

func (n *Node) SetStartBlock(startBlock uint64) {
	n.startBlockMux.Lock()
	defer n.startBlockMux.Unlock()
	n.startBlock = startBlock
}

func (n *Node) setNonce(nonce uint64) {
	n.nonceMux.Lock()
	defer n.nonceMux.Unlock()
	n.nonce = nonce
}

func (n *Node) incrementNonce() {
	n.nonceMux.Lock()
	defer n.nonceMux.Unlock()
	n.nonce += 1
}

func (n *Node) getNonce() uint64 {
	n.nonceMux.RLock()
	defer n.nonceMux.RUnlock()
	return n.nonce
}

func (n *Node) sendTransaction() (txHash common.Hash, err error) {
	value := baseValue + rand.Int63n(900)
	tx := types.NewTransaction(n.getNonce(), n.address, big.NewInt(value), defaultGasLimit, big.NewInt(defaultGasPrice), nil)

	signer := types.NewEIP155Signer(big.NewInt(chainId))
	rawTx, err := types.SignTx(tx, signer, n.privateKey)
	if err != nil {
		return common.Hash{}, err
	}

	return rawTx.Hash(), n.client.SendTransaction(context.TODO(), rawTx)
}

func (n *Node) startLoad(delay time.Duration, stopCh <-chan struct{}, errCh chan error, successCh chan common.Hash, wg *sync.WaitGroup) {
	errToChannel := func(err error) {
		select {
		case errCh <- errors.Wrap(err, fmt.Sprintf("On node %s", n.config.name)):
		default:
			return
		}
	}

	sendTxTicker := time.NewTicker(delay)
	privateKey, err := crypto.HexToECDSA(n.config.privateKey)
	if err != nil {
		errToChannel(err)
		return
	}
	n.privateKey = privateKey

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	n.address = address

	nonce, err := n.client.NonceAt(context.TODO(), address, nil)
	if err != nil {
		errToChannel(err)
		return
	}
	n.setNonce(nonce)
	fmt.Printf("Nonce is %d on node %s\n", n.getNonce(), n.config.name)

	// Get latest block
	block, err := n.client.BlockByNumber(context.TODO(), nil)
	if err != nil {
		errToChannel(err)
		return
	}
	n.SetStartBlock(block.NumberU64())

	for {
		select {
		case <-sendTxTicker.C:
			//now := time.Now().UnixNano() / 1e6
			//fmt.Printf("sendTransaction called on node %s at %d\n", n.config.name, now)
			hash, err := n.sendTransaction()
			if err != nil {
				errToChannel(err)
				return
			}
			select {
			case successCh <- hash:
			default:
			}
			n.incrementNonce()
			if config.debug {
				fmt.Printf("Successfully sent transaction %s to node %s\n", hash.Hex(), n.config.name)
			}
		case <-stopCh:
			// Get latest block
			block, err := n.client.BlockByNumber(context.TODO(), nil)
			if err != nil {
				errToChannel(err)
				return
			}
			n.SetEndBlock(block.NumberU64())
			fmt.Printf("Node %s started on block %d and ended on block %d\n", n.config.name, n.StartBlock(), n.EndBlock())
			n.client.Close()
			wg.Done()
			return
		}
	}
}

func TestLoadStress(t *testing.T) {
	if config.intensity > 35 {
		t.Fatal("The test is unable to run at more than 35 txs per node - please increase the number of nodes instead")
	}

	nodes := make([]*Node, len(config.nodes))

	for index, nodeConfig := range config.nodes {
		client, err := ethclient.Dial(nodeConfig.uri)
		if err != nil {
			fmt.Printf("Unable to connect to %s: %s\n", nodeConfig.name, err.Error())
			t.Fail()
		}
		nodes[index] = &Node{client: client, config: nodeConfig}
	}

	successfulTxs := 0
	errCh := make(chan error, len(nodes)*int(config.intensity))
	successCh := make(chan common.Hash, len(nodes)*int(config.intensity))
	stopCh := make(chan struct{})

	delay := (float64(1.0) / float64(config.intensity)) * float64(1000)
	fmt.Printf("-----Starting test (⟠ ETH)-----\n\tDelay %.2fms\n\tOverall load is %d tx per sec\n\tGoing to send %d txs over %s\n", delay, config.intensity*uint64(len(config.nodes)), config.intensity*uint64(len(config.nodes))*uint64(config.duration.Seconds()), config.duration.String())
	time.AfterFunc(config.duration, func() {
		close(stopCh)
	})

	var wg sync.WaitGroup
	for _, n := range nodes {
		wg.Add(1)
		go n.startLoad(time.Duration(delay)*time.Millisecond, stopCh, errCh, successCh, &wg)
	}

stopped:
	for {
		select {
		case err := <-errCh:
			close(stopCh)
			fmt.Printf("Error while performing load: %s\n", err.Error())
			t.Fatal()
		case <-successCh:
			successfulTxs++
		case <-stopCh:
			break stopped
		}
	}

	wg.Wait()
	close(successCh)
	successfulTxs += len(successCh)

	fmt.Printf("Test finished with %d successful transactions\n", successfulTxs)
	fmt.Println("Starting block checker")

	randomNode := nodes[rand.Intn(len(nodes))]
	randomNodeUri := randomNode.config.uri
	startBlock, endBlock := randomNode.StartBlock(), randomNode.EndBlock()
	node, err := ethclient.Dial(randomNodeUri)
	if err != nil {
		t.Fatalf("Unable to connect to node %s for block time checking", randomNodeUri)
	}

	data := make(map[uint64]uint64)
	for currentBlock := startBlock + 2; currentBlock < endBlock; currentBlock++ {
		prevBlock, err := node.BlockByNumber(context.TODO(), big.NewInt(int64(currentBlock-1)))
		if err != nil {
			t.Fatalf("Unable to get prev block - please, check them by hand (s: %d, e: %d) error: %s", startBlock, endBlock, err.Error())
		}
		currBlock, err := node.BlockByNumber(context.TODO(), big.NewInt(int64(currentBlock)))
		if err != nil {
			t.Fatalf("Unable to get current block - please, check them by hand (s: %d, e: %d) error: %s", startBlock, endBlock, err.Error())
		}

		blockTime := currBlock.Time() - prevBlock.Time()

		data[currBlock.Time()] = blockTime
	}

	f, err := os.Create(fmt.Sprintf("./load-stress-eth-%d-%d.csv", len(config.nodes)*int(config.intensity), time.Now().Unix()))
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}

	w := csv.NewWriter(f)
	err = w.Write([]string{"timestamp", "block_time"})
	if err != nil {
		t.Fatal(err)
	}

	for timestamp, blockTime := range data {
		err := w.Write([]string{fmt.Sprintf("%v", timestamp), fmt.Sprintf("%v", blockTime)})
		if err != nil {
			t.Fatal(err)
		}
	}
	w.Flush()

}
