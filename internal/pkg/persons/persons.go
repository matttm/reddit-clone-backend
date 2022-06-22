package persons

type Person struct {
    id float32
    username string
    createdAt string
    updatedAt string
    posts []*Posts
}
