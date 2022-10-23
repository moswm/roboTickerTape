/*
 * baev.one / Personal website
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
	"strings"
	"time"
	"strconv"
)

func start_index(w http.ResponseWriter, r *http.Request) {
	
	ruri := r.URL.Path
	ruriGo := strings.Split(ruri, "/")
	ruriGo_len := len(ruriGo)
	if (ruriGo[1]=="") { ruriGo[1]="/" }

	title := ""
	if ruriGo_len>3 && ruriGo[3]!="" {
		title = pg_title[ruriGo[1]+"/"+ruriGo[2]+"/"+ruriGo[3]]
	} else if ruriGo_len>2 && ruriGo[2]!="" {
		title = pg_title[ruriGo[1]+"/"+ruriGo[2]]
		ruriGo = append(ruriGo, "")
	} else {
		title = pg_title[ruriGo[1]]
		ruriGo = append(ruriGo, "")
		ruriGo = append(ruriGo, "")
	}

	if (title=="") {
		hdl_redirect_home(w, r)
	} else {
		tpl_data := map[string]string{
			"pfxstdir":		proj_dir,
			"title":		title,
			"myurl":		myurl,
			"author":		author,
			"copyright":	copyright,
			"ruriGo_1":		ruriGo[1],
			"ruriGo_2":		ruriGo[2],
			"ruriGo_3":		ruriGo[3],
			"time":			strconv.FormatInt(time.Now().Sub(time.Unix(0,0)).Nanoseconds(),10)}
		
		DispHtml(w, "main", tpl_data)
	}
}
