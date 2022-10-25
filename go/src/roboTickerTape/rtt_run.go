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
	"net/http"
	"time"
	"fmt"
)

type rChEx struct {
	ex string
	resp *http.Response
}

func getWr(ex string,ul string,r chan rChEx) {
	t_out:=time.After(10*time.Second)
	select {
	default:
		if resp,err:=http.Get(ul); err==nil {
			r<-rChEx{ex,resp}
		} else {
			fmt.Println("getWr - request error", err.Error())
		}
	case <-t_out:
		return
	}
}

func getEx(ex string,tm int,ul string) {
	ch:=make(chan rChEx)
	for {
		go getWr(ex,ul,ch)
		time.Sleep(time.Duration(tm)*time.Second)
		rTTparse(ch)
	}
}

func rTTmain() {
	//chPbl=make(chan string,len(rtt_vExNm))
	for ex,_:=range rtt_vExUrl {
		go getEx(ex,rtt_vExTm[ex],rtt_vExUrl[ex])
	}
}
