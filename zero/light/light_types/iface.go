package light_types

import (
	"math/big"

	"github.com/sero-cash/go-sero/common/hexutil"

	"github.com/sero-cash/go-czero-import/keys"
	"github.com/sero-cash/go-sero/zero/txs/assets"
	"github.com/sero-cash/go-sero/zero/txs/stx"
)

type GIn struct {
	SKr     keys.PKr
	Out     Out
	Witness Witness
}

type GOut struct {
	PKr   keys.PKr
	Asset assets.Asset
	Memo  keys.Uint512
}

type GTx struct {
	Gas      hexutil.Uint64
	GasPrice hexutil.Big
	Tx       stx.T
}

type GenTxParam struct {
	Gas      uint64
	GasPrice big.Int
	From     Kr
	Ins      []GIn
	Outs     []GOut
}

type ISLI interface {
	CreateKr() Kr
	DecOuts(outs []Out, skr *keys.PKr) ([]DOut, error)
	GenTx(param *GenTxParam) (GTx, error)
}

type ISRI interface {
	GetBlocksInfo(start uint64, count uint64) ([]Block, error)
	GetAnchor(roots []keys.Uint256) ([]Witness, error)
	CommitTx(tx *GTx) error
}
