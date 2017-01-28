// Copyright (C) 2017 Jack Maloney. All Rights Reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {

	// read JSON file
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	cards := make([]interface{}, 0)
	json.Unmarshal(file, &cards)

	names := make([]string, 0)

	for _, c := range cards {
		m := c.(map[string]interface{})
    n := m["name"]
    id := m["id"]
    if n != nil && id != nil {
      fmt.Println(n.(string) + id.(string))
  		names = append(names, n.(string) + id.(string))
    }
	}

	sort.Strings(names)

	out := make([]interface{}, 0)

	for _, name := range names {
		for _, c := range cards {
      c := c.(map[string]interface{})
      n := c["name"]
      id := c["id"]
      if (n != nil && id != nil) && (n.(string) + id.(string)) == name {
        out = append(out, c)
      }
		}
	}

  o, err := json.MarshalIndent(out, "", "    ")
  fmt.Println(o)

}
