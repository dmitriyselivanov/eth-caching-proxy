package controller

import (
	"errors"
	"eth-caching-proxy/internal/util"
	"eth-caching-proxy/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

const (
	ErrBlockNotFound = "not found"
	ErrTxNotFound    = "not found"
)

type ethBlockController struct {
	services *service.Manager
}

func NewEthBlockController(services *service.Manager) *ethBlockController {
	return &ethBlockController{services: services}
}

func (ctrl *ethBlockController) RegisterRoutes(r *gin.Engine) {
	r.GET("/block/:blockNumber", ctrl.GetBlockByNumber)
	r.GET("/block/latest", ctrl.GetLatestBlock)

	r.GET("/block/:blockNumber/txs/:txNumOrHash", ctrl.GetTxFromBlock)
	r.GET("/block/latest/txs/:txNumOrHash", ctrl.GetTxFromLatestBlock)
}

func (ctrl *ethBlockController) GetLatestBlock(c *gin.Context) {
	blockDTO, err := ctrl.services.BlockService.LatestBlock()
	if err != nil && err.Error() == ErrBlockNotFound {
		// 404 Not found
		c.JSON(http.StatusNotFound, util.WrapError(err))
		return
	}

	if err != nil {
		// 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, util.WrapError(err))
		return
	}

	// 200 OK
	c.JSON(http.StatusOK, blockDTO)
}

func (ctrl *ethBlockController) GetBlockByNumber(c *gin.Context) {
	// Parse block number from request path
	blockNumber64, err := strconv.ParseInt(c.Param("blockNumber"), 10, 64)
	if err != nil {
		// 400 Bad Request
		c.JSON(http.StatusBadRequest, util.WrapError(errors.New("error parsing the provided block number")))
		return
	}
	blockNumber := big.NewInt(blockNumber64)

	// Get block from service
	blockDTO, err := ctrl.services.BlockService.BlockByNumber(blockNumber)

	if err != nil && err.Error() == ErrBlockNotFound {
		// 404 Not found
		c.JSON(http.StatusNotFound, util.WrapError(err))
		return
	}

	if err != nil {
		// 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, util.WrapError(err))
		return
	}

	// 200 OK
	c.JSON(http.StatusOK, blockDTO)
}

func (ctrl *ethBlockController) GetTxFromLatestBlock(c *gin.Context) {
	txNumOrHash := c.Param("txNumOrHash")

	// Try to parse transaction number from request path
	txNumber64, err := strconv.ParseUint(txNumOrHash, 10, 64)
	if err == nil {
		// Ok, we have tx number, not hash, let's try to get tx by number from service
		ctrl.txFromLatestBlockByIndex(c, int(txNumber64))
		return
	}

	// Here we try to parse transaction hash, because parsing of the transaction num failed
	txHash := common.HexToHash(txNumOrHash)
	ctrl.txFromLatestBlockByHash(c, txHash)
}

func (ctrl *ethBlockController) GetTxFromBlock(c *gin.Context) {
	// Parse block number from request path
	blockNumber64, err := strconv.ParseInt(c.Param("blockNumber"), 10, 64)
	if err != nil {
		// 400 Bad Request
		c.JSON(http.StatusBadRequest, util.WrapError(errors.New("error parsing the provided block number")))
		return
	}
	blockNumber := big.NewInt(blockNumber64)

	txNumOrHash := c.Param("txNumOrHash")

	// Try to parse transaction number from request path
	txNumber64, err := strconv.ParseUint(txNumOrHash, 10, 64)
	if err == nil {
		// Ok, we have tx number, not hash, let's try to get tx by number from service
		ctrl.txFromBlockByIndex(c, blockNumber, int(txNumber64))
		return
	}

	// Here we try to parse transaction hash, because parsing of the transaction num failed
	txHash := common.HexToHash(txNumOrHash)
	ctrl.txFromBlockByHash(c, blockNumber, txHash)
}

func (ctrl *ethBlockController) txFromBlockByIndex(c *gin.Context, blockNumber *big.Int, txNumber int) {
	tx, e := ctrl.services.BlockService.TxFromBlockByIndex(blockNumber, txNumber)
	if e != nil && e.Error() == ErrTxNotFound {
		// 404 Not Found
		c.JSON(http.StatusNotFound, util.WrapError(e))
		return
	}

	if e != nil {
		// 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, util.WrapError(e))
		return
	}

	// 200 OK
	c.JSON(http.StatusOK, tx)
}

func (ctrl *ethBlockController) txFromBlockByHash(c *gin.Context, blockNumber *big.Int, txHash common.Hash) {
	tx, e := ctrl.services.BlockService.TxFromBlockByHash(blockNumber, txHash)
	if e != nil && e.Error() == ErrTxNotFound {
		// 404 Not Found
		c.JSON(http.StatusNotFound, util.WrapError(e))
		return
	}

	if e != nil {
		// 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, util.WrapError(e))
		return
	}

	// 200 OK
	c.JSON(http.StatusOK, tx)
}

func (ctrl *ethBlockController) txFromLatestBlockByIndex(c *gin.Context, txNumber int) {
	tx, e := ctrl.services.BlockService.TxFromLatestBlockByIndex(txNumber)
	if e != nil && e.Error() == ErrTxNotFound {
		// 404 Not Found
		c.JSON(http.StatusNotFound, util.WrapError(e))
		return
	}

	if e != nil {
		// 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, util.WrapError(e))
		return
	}

	// 200 OK
	c.JSON(http.StatusOK, tx)
}

func (ctrl *ethBlockController) txFromLatestBlockByHash(c *gin.Context, txHash common.Hash) {
	// Get tx by hash from service
	tx, err := ctrl.services.BlockService.TxFromLatestBlockByHash(txHash)

	if err != nil && err.Error() == ErrTxNotFound {
		// 404 Not Found
		c.JSON(http.StatusNotFound, util.WrapError(err))
		return
	}

	if err != nil {
		// 500 Internal Server Error
		c.JSON(http.StatusInternalServerError, util.WrapError(err))
		return
	}

	// 200 OK
	c.JSON(http.StatusOK, tx)
}
