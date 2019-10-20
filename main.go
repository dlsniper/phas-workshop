//    Copyright 2019 Florin Pățan
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

// Import packages to ensure that we can compile the application
import (
	"log"

	_ "github.com/charithe/porcupine-go"
	_ "github.com/gordonklaus/portaudio"
)

func main() {
	log.Println("If you see this, then compilation worked and we can go forward.")
}
