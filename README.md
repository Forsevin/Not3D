Oden
====

Oden is a open source cross-platform game Engine written in Go (if you haven't guessed).

It built on the Entity Component System design pattern, but does some things lazy (e.g storing Components in the objects/entities). The goal is a working game Engine useful for general purposes and without having to write any additional systems or Components (though, for speed this might be useful). To achive this Oden will use the embedable language Gel which hopefully meet the needs of a scripting language for a game Engine, both by usability and speed.



###installation

`$ go get github.com/Forsevin/oden`

Requires SDL2
