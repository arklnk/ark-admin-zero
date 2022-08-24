package utils

import (
	"fmt"
	"math/rand"
)

func AvatarUrl() string {
	url := fmt.Sprintf(
		"https://avataaars.io/?clotheColor=%s&accessoriesType=%s&avatarStyle=%s&clotheType=%s&eyeType=%s&eyebrowType=%s&facialHairColor=%s&facialHairType=%s&hairColor=%s&hatColor=%s&mouthType=%s&skinColor=%s&topType=%s",
		clotheColor(),
		accessoriesType(),
		avatarStyle(),
		clotheType(),
		eyeType(),
		eyebrowType(),
		facialHairColor(),
		facialHairType(),
		hairColor(),
		hatColor(),
		mouthType(),
		skinColor(),
		topType())

	return url
}

func clotheColor() string {
	clotheColor := make(map[int]string, 0)
	clotheColor[0] = "Black"
	clotheColor[1] = "Blue01"
	clotheColor[2] = "Blue02"
	clotheColor[3] = "Blue03"
	clotheColor[4] = "Gray01"
	clotheColor[5] = "Gray02"
	clotheColor[6] = "Heather"
	clotheColor[7] = "PastelBlue"
	clotheColor[8] = "PastelGreen"
	clotheColor[9] = "PastelOrange"
	clotheColor[10] = "PastelRed"
	clotheColor[11] = "PastelYellow"
	clotheColor[12] = "Pink"
	clotheColor[13] = "Red"
	clotheColor[14] = "White"
	return clotheColor[rand.Intn(15)]
}

func accessoriesType() string {
	accessoriesType := make(map[int]string, 0)
	accessoriesType[0] = "Blank"
	accessoriesType[1] = "Kurt"
	accessoriesType[2] = "Prescription01"
	accessoriesType[3] = "Prescription02"
	accessoriesType[4] = "Round"
	accessoriesType[5] = "Sunglasses"
	accessoriesType[6] = "Wayfarers"
	return accessoriesType[rand.Intn(7)]
}

func avatarStyle() string {
	avatarStyle := make(map[int]string, 0)
	avatarStyle[0] = "Circle"
	avatarStyle[1] = "Transparent"
	return avatarStyle[rand.Intn(1)]
}

func clotheType() string {
	clotheType := make(map[int]string, 0)
	clotheType[0] = "BlazerShirt"
	clotheType[1] = "BlazerSweater"
	clotheType[2] = "CollarSweater"
	clotheType[3] = "GraphicShirt"
	clotheType[4] = "Hoodie"
	clotheType[5] = "Overall"
	clotheType[6] = "ShirtCrewNeck"
	clotheType[7] = "ShirtScoopNeck"
	clotheType[8] = "ShirtVNeck"
	return clotheType[rand.Intn(9)]
}

func eyeType() string {
	eyeType := make(map[int]string, 0)
	eyeType[0] = "Close"
	eyeType[1] = "Cry"
	eyeType[2] = "Default"
	eyeType[3] = "Dizzy"
	eyeType[4] = "EyeRoll"
	eyeType[5] = "Happy"
	eyeType[6] = "Hearts"
	eyeType[7] = "Side"
	eyeType[8] = "Squint"
	eyeType[9] = "Surprised"
	eyeType[10] = "Wink"
	eyeType[11] = "WinkWacky"
	return eyeType[rand.Intn(12)]
}

func eyebrowType() string {
	eyebrowType := make(map[int]string, 0)
	eyebrowType[0] = "Angry"
	eyebrowType[1] = "AngryNatural"
	eyebrowType[2] = "Default"
	eyebrowType[3] = "DefaultNatural"
	eyebrowType[4] = "FlatNatural"
	eyebrowType[5] = "RaisedExcited"
	eyebrowType[6] = "RaisedExcitedNatural"
	eyebrowType[7] = "SadConcerned"
	eyebrowType[8] = "SadConcernedNatural"
	eyebrowType[9] = "UnibrowNatural"
	eyebrowType[10] = "UpDown"
	eyebrowType[11] = "UpDownNatural"
	return eyebrowType[rand.Intn(12)]
}

func facialHairColor() string {
	facialHairColor := make(map[int]string, 0)
	facialHairColor[0] = "Auburn"
	facialHairColor[1] = "Black"
	facialHairColor[2] = "Blonde"
	facialHairColor[3] = "BlondeGolden"
	facialHairColor[4] = "Brown"
	facialHairColor[5] = "BrownDark"
	facialHairColor[6] = "Platinum"
	facialHairColor[7] = "Red"
	return facialHairColor[rand.Intn(8)]
}

func facialHairType() string {
	facialHairType := make(map[int]string, 0)
	facialHairType[0] = "Blank"
	facialHairType[1] = "BeardMedium"
	facialHairType[2] = "BeardLight"
	facialHairType[3] = "BeardMajestic"
	facialHairType[4] = "MoustacheFancy"
	facialHairType[5] = "MoustacheMagnum"
	return facialHairType[rand.Intn(6)]
}

func hairColor() string {
	hairColor := make(map[int]string, 0)
	hairColor[0] = "Auburn"
	hairColor[1] = "Black"
	hairColor[2] = "Blonde"
	hairColor[3] = "BlondeGolden"
	hairColor[4] = "Brown"
	hairColor[5] = "BrownDark"
	hairColor[6] = "PastelPink"
	hairColor[7] = "Blue"
	hairColor[8] = "Platinum"
	hairColor[9] = "Red"
	hairColor[10] = "SilverGray"
	return hairColor[rand.Intn(11)]
}

func hatColor() string {
	hatColor := make(map[int]string, 0)
	hatColor[0] = "Black"
	hatColor[1] = "Blue01"
	hatColor[2] = "Blue02"
	hatColor[3] = "Blue03"
	hatColor[4] = "Gray01"
	hatColor[5] = "Gray02"
	hatColor[6] = "Heather"
	hatColor[7] = "PastelBlue"
	hatColor[8] = "PastelGreen"
	hatColor[9] = "PastelOrange"
	hatColor[10] = "PastelRed"
	hatColor[11] = "PastelYellow"
	hatColor[12] = "Pink"
	hatColor[13] = "Red"
	hatColor[14] = "White"
	return hatColor[rand.Intn(15)]
}

func mouthType() string {
	mouthType := make(map[int]string, 0)
	mouthType[0] = "Concerned"
	mouthType[1] = "Default"
	mouthType[2] = "Disbelief"
	mouthType[3] = "Eating"
	mouthType[4] = "Grimace"
	mouthType[5] = "Sad"
	mouthType[6] = "ScreamOpen"
	mouthType[7] = "Serious"
	mouthType[8] = "Smile"
	mouthType[9] = "Tongue"
	mouthType[10] = "Twinkle"
	mouthType[11] = "Vomit"
	return mouthType[rand.Intn(12)]
}

func skinColor() string {
	skinColor := make(map[int]string, 0)
	skinColor[0] = "Tanned"
	skinColor[1] = "Yellow"
	skinColor[2] = "Pale"
	skinColor[3] = "Light"
	skinColor[4] = "Brown"
	skinColor[5] = "DarkBrown"
	skinColor[6] = "Black"
	return skinColor[rand.Intn(7)]
}

func topType() string {
	topType := make(map[int]string, 0)
	topType[0] = "NoHair"
	topType[1] = "Eyepatch"
	topType[2] = "Hat"
	topType[3] = "Hijab"
	topType[4] = "Turban"
	topType[5] = "WinterHat1"
	topType[6] = "WinterHat2"
	topType[7] = "WinterHat3"
	topType[8] = "WinterHat4"
	topType[9] = "LongHairBigHair"
	topType[10] = "LongHairBob"
	topType[11] = "LongHairBun"
	topType[12] = "LongHairCurly"
	topType[13] = "LongHairCurvy"
	topType[14] = "LongHairDreads"
	topType[15] = "LongHairFrida"
	topType[16] = "LongHairFro"
	topType[17] = "LongHairFroBand"
	topType[18] = "LongHairNotTooLong"
	topType[19] = "LongHairShavedSides"
	topType[20] = "LongHairMiaWallace"
	topType[21] = "LongHairStraight"
	topType[22] = "LongHairStraight2"
	topType[23] = "LongHairStraightStrand"
	topType[24] = "ShortHairDreads01"
	topType[25] = "ShortHairDreads02"
	topType[26] = "ShortHairFrizzle"
	topType[27] = "ShortHairShaggyMullet"
	topType[28] = "ShortHairShortCurly"
	topType[29] = "ShortHairShortFlat"
	topType[30] = "ShortHairShortRound"
	topType[31] = "ShortHairShortWaved"
	topType[32] = "ShortHairSides"
	topType[33] = "ShortHairTheCaesar"
	topType[34] = "ShortHairTheCaesarSidePart"
	return topType[rand.Intn(35)]
}
