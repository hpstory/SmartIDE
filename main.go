/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"smartide-cli/cmd"
)

func main() {
	fmt.Println(`
	 _____                      _     ___________ _____ 
	/  ___|                    | |   |_   _|  _  \  ___|
	\ ` + "`" + `--. _ __ ___   __ _ _ __| |_    | | | | | | |__
	 ` + "`" + `--. \ '_ ` + "`" + ` _ \ / _` + "`" + ` | '__| __|   | | | | | |  __|
	/\__/ / | | | | | (_| | |  | |_   _| |_| |/ /| |___
	\____/|_| |_| |_|\__,_|_|   \__|  \___/|___/ \____/ `)
	cmd.Execute()
}
