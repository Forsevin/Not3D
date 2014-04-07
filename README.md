Oden
====

Oden is a open source cross-platform game Engine written in Go (if you haven't guessed).

It's built on the Entity Component System design pattern, but does some things lazy (e.g storing Components in the objects/entities). The goal is a working game Engine useful for general purposes 2D game development and without having to write any additional systems or Components (though, for speed this might be useful).



###installation

`$ go get github.com/Forsevin/oden`

Requires SDL2

###notes
* SpriteBatch probably isn't implemented correctly (someone who actually understands SpriteBatch are very welcome to pick it up)
* The api (interface.go) requires a lot of work

###license
LGPLv3
