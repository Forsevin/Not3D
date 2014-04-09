function Initialize() {

}

function Update() {
	if (input.KeyDown("w")) {
		object.SetY(object.GetY() - 3)
	}
	if (input.KeyDown("s")) {
		object.SetY(object.GetY() + 3)
	}
	if (input.KeyDown("a")) {
		object.SetX(object.GetX() - 3)
	}
	if (input.KeyDown("d")) {
		object.SetX(object.GetX() + 3)
	}
}