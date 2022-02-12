package solution

import "github.com/kyokomi/emoji/v2"


func GetMessage() string {
	helloWorld := emoji.Sprint("Hello :world_map:!")
	return helloWorld
}
