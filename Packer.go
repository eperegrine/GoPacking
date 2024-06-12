package main

type PackedEntry struct {
	pack_volume int
	quantity    int
}

// / <summary>
// / Packs the items in the most efficient way
// / </summary>
// / <param name="availablePacks">The available packs, in descending order</param>
// / <param name="quantity">The quantity of the item</param>
func packItems(availablePacks []int, quantity int) map[int]PackedEntry {

	packedItems := make(map[int]PackedEntry)

	remainingItems := quantity
	lastPackIndex := len(availablePacks) - 1

	for i, pack := range availablePacks {

		for remainingItems >= pack {
			// pack the item
			remainingItems -= pack
			addPack(packedItems, pack)
		}

		morePacksAvailable := i < lastPackIndex
		isPenultimatePack := i == lastPackIndex-1
		packingIsIncomplete := remainingItems > 0

		//If we are the penultimate pack and the last pack is too small then add it
		if packingIsIncomplete && isPenultimatePack && availablePacks[i+1] < remainingItems {
			remainingItems -= pack
			addPack(packedItems, pack)
		}

		if packingIsIncomplete && !morePacksAvailable {
			remainingItems -= pack
			addPack(packedItems, pack)
		}

		if !packingIsIncomplete {
			break
		}
	}

	return packedItems

}

func addPack(packedItems map[int]PackedEntry, pack int) {
	activePack, ok := packedItems[pack]
	if ok {
		packedItems[pack] = PackedEntry{
			pack_volume: activePack.pack_volume,
			quantity:    activePack.quantity + 1,
		}
	} else {
		packedItems[pack] = PackedEntry{
			pack_volume: pack,
			quantity:    1,
		}
	}
}
