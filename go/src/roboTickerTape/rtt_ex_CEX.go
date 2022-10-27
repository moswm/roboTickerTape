/*
 * roboTickerTape / Aggregates and caches ticker tapes from exchanges
 * Copyright (C) 2022 Baev
 *
 * MIT License
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 * 
 * GNU General Public License, version 2
 * This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later version.
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 * You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 * 
*/

package main

import (
	"fmt"
	"strings"
	"log"
	"encoding/json"
	"strconv"
)

type vCEX_json_dt_tk struct {
	Vol string
	Low string
	Open string
	High string
	Last string
	Buy string
	Buy_amount string
	Sell string
	Sell_amount string
}
type vCEX_json_dt struct {
	Date int64
	Ticker map[string]vCEX_json_dt_tk
}
type vCEX_json struct {
	Code int
	Data vCEX_json_dt
}

func rTTparse_CEX(bdy []byte) {

	var jInf vCEX_json
	json.Unmarshal(bdy, &jInf)

	var pairArr=[]string{"BTC","USDT"}
	var rslTkr []string
	for _,v:=range rtt_tickers {
		for _,p:=range pairArr {
			if tckrLn:=rTT_CEX_getticker(v,p,jInf.Data.Ticker);tckrLn!="" {
				rslTkr=append(rslTkr,tckrLn)
			}
		}
	}

	if err:=fl_wrtLines(rslTkr,www_path+"rTT/ex_CEX_tickers");err!=nil {
		log.Fatalf("error, CEX tickers write: %s",err)
	}
	for _,v:=range rslTkr {
		fmt.Println(v)
	}

}

func rTT_CEX_getticker(name string,pair string,ticker map[string]vCEX_json_dt_tk) string {
	if name==pair { return "" }
	rslt:=""
	if tckr,err:=ticker[name+pair];err {
		tckrVol,_:=strconv.ParseFloat(tckr.Vol,64)
		tckrLast,_:=strconv.ParseFloat(tckr.Last,64)
		var rsltArr = []string{
			pair+"_"+name,
			"trading",
			tckr.Last,
			tckr.Sell,
			tckr.Buy,
			tckr.Low,
			tckr.High,
			strconv.FormatFloat(tckrVol*tckrLast,'f',8,64),
			tckr.Vol}
		rslt=strings.Join(rsltArr,"	")
	}
	return rslt
}
