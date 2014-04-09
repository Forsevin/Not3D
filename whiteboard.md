### Todos, design decisions etc. ###


* Currently components are held in the objects, this breaks the Entity Component System pattern whereas entities (objects) only are mere identifier.
I personally don't feel like this is a big problem though.

* The components in components.go could be moved to their respective 'main' system, e.g SpriteComponent in RenderSystem and so on.
