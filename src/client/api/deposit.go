package api

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skycoin/skycoin-exchange/src/client/account"
	"github.com/skycoin/skycoin-exchange/src/pp"
	"github.com/skycoin/skycoin-exchange/src/sknet"
)

// GetDepositAddress get deposit address from exchange server.
func GetDepositAddress(se Servicer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rlt := &pp.EmptyRes{}
		for {
			a, err := account.GetActive()
			if err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrRes(err)
				break
			}

			cp := r.FormValue("coin_type")
			if cp == "" {
				err := errors.New("coin type empty")
				logger.Error(err.Error())
				rlt = pp.MakeErrRes(err)
				break
			}

			req := pp.GetDepositAddrReq{
				Pubkey:   pp.PtrString(a.Pubkey),
				CoinType: pp.PtrString(cp),
			}

			var res pp.GetDepositAddrRes

			if err := sknet.EncryGet(se.GetServAddr(), "/create/deposit_address", req, &res); err != nil {
				logger.Error(err.Error())
				rlt = pp.MakeErrResWithCode(pp.ErrCode_ServerError)
				break
			}

			sendJSON(w, res)
			return
		}
		sendJSON(w, rlt)
	}
}
