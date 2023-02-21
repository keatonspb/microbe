package assets

import "bacteria/helper/storage"

const (
	ImagePlayer = iota
	ImageMobA
	ImageMobB
	ImageMobC
	ImageCell
)

var ImagePaths = map[storage.AssetKey]string{
	ImagePlayer: "assets/player.png",
	ImageMobA:   "assets/ley1.png",
	ImageMobB:   "assets/ley2.png",
	ImageMobC:   "assets/ley3.png",
	ImageCell:   "assets/cell.png",
}
