package initializer

import (
	"fmt"

	"github.com/honeyscience/crystal-caves-golang/internal/handler"
)

func Initialize() {
	fmt.Println("Running initializer")

	/* create game objects */

	/* launch websocket listener
	   to wait for incoming connections
	*/
	handler.Initialize()
	/* start ECS and physics update loop */
}
