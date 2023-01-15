package uniswap_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"io/ioutil"
	"math/big"
	"mev-template-go/contracts/bindings/uniswap_v2/uniswap_v2_factory"
	"mev-template-go/contracts/bindings/uniswap_v2/uniswap_v2_router"
	"mev-template-go/types"
	"net/http"
	"os"
	"strings"
	"time"
)

type UniswapV2 struct {
	Router   UniV2Router.UniV2Router
	RouerAbi abi.ABI
	Factory  UniV2Factory.UniV2Factory
}

func New(config types.Config, routerAddress common.Address, factoryAddress common.Address, pathToAbi string) (*UniswapV2, error) {
	router, err := UniV2Router.NewUniV2Router(routerAddress, &config.Client)
	if err != nil {
		return nil, fmt.Errorf("Error creating new router: %v", err)
	}
	factory, err := UniV2Factory.NewUniV2Factory(factoryAddress, &config.Client)
	if err != nil {
		return nil, fmt.Errorf("Error creating new factory: %v", err)
	}
	routerAbi, err := abi.JSON(strings.NewReader(GetLocalABI(pathToAbi)))
	if err != nil {
		return nil, fmt.Errorf("Error creating new router abi: %v", err)
	}
	return &UniswapV2{Router: *router, RouerAbi: routerAbi, Factory: *factory}, nil
}

func (uniV2 UniswapV2) DecodeTransactionInputData(data []byte) {
	methodSigData := data[:4]
	method, err := uniV2.RouerAbi.MethodById(methodSigData)
	if err != nil {
		fmt.Println("Error getting method by ID: ", err)
		return
	}

	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		fmt.Println("Error unpacking inputs: ", err)
		return
	}

	fmt.Printf("Method Name: %s\n", method.Name)
	fmt.Printf("Method inputs: %v\n", inputsMap)
}

func GetLocalABI(path string) string {
	abiFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer abiFile.Close()

	result, err := io.ReadAll(abiFile)
	if err != nil {
		fmt.Println(err)
	}
	return string(result)
}

type UniV2TempStruct struct {
	Data struct {
		Pairs []struct {
			ID       string `json:"id"`
			Reserve0 string `json:"reserve0"`
			Reserve1 string `json:"reserve1"`
			Token0   struct {
				ID string `json:"id"`
			} `json:"token0"`
			Token1 struct {
				ID string `json:"id"`
			} `json:"token1"`
		} `json:"pairs"`
	} `json:"data"`
}

func GetUniV2Pools() ([]types.UniV2Pool, error) {
	poolsQuery := `
		{
			pairs(first: 1000 orderBy: trackedReserveETH, orderDirection: desc) 
			{ id, reserve0, reserve1, token0 {id}, token1{id} }
		}
	`
	var UniV2Pairs []types.UniV2Pool

	data, err := queryDataV2(poolsQuery)
	if err != nil {
		return nil, err
	}

	var final UniV2TempStruct
	if err := json.Unmarshal(data, &final); err != nil {
		return nil, err
	}

	for _, pair := range final.Data.Pairs {
		pool := types.UniV2Pool{
			Address:  common.HexToAddress(pair.ID),
			Token0:   common.HexToAddress(pair.Token0.ID),
			Token1:   common.HexToAddress(pair.Token1.ID),
			Reserve0: FloatStringToBigInt(pair.Reserve0),
			Reserve1: FloatStringToBigInt(pair.Reserve1),
			Fees:     997,
		}
		UniV2Pairs = append(UniV2Pairs, pool)
	}
	return UniV2Pairs, nil
}

func queryDataV2(query string) ([]byte, error) {
	data := map[string]string{
		"query": query,
	}
	queryJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2", bytes.NewBuffer(queryJSON))
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

// TODO - only works for 18 decimals
func FloatStringToBigInt(val string) big.Int {
	bigval := new(big.Float)
	bigval.SetString(val)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))
	bigval.Mul(bigval, coin)

	result := new(big.Int)
	f, _ := bigval.Uint64()
	result.SetUint64(f)

	return *result
}
