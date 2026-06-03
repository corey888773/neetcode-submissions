type Node struct {
    key int
    value int
    next *Node
    prev *Node
}

type LRUCache struct {
    capacity int
    count int
    head *Node
    tail *Node
    cache map[int]*Node
}

func Constructor(capacity int) LRUCache {
    head, tail := &Node{}, &Node{}
    head.prev = tail
    tail.next = head

    return LRUCache{
        capacity: capacity,
        count: 0,
        head: head,
        tail: tail,
        cache: map[int]*Node{},
    }
}

func (this *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) insertNode(node *Node) {
    node.next = this.head
    node.prev = this.head.prev
    this.head.prev.next = node
    this.head.prev = node
}

func (this *LRUCache) moveToFront(node *Node) {
    this.removeNode(node)
    this.insertNode(node)
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.moveToFront(node)
        return node.value
    }

    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
        node.value = value
        this.moveToFront(node)
        return
    }

    newNode := &Node{key:key, value:value}
    if this.count < this.capacity {
        this.count++
        this.cache[key] = newNode
        this.insertNode(newNode)
        return
    }

    // remove least freq used
    leastUsedNode := this.tail.next
    delete(this.cache, leastUsedNode.key)
    this.removeNode(leastUsedNode)

    this.cache[key] = newNode
    this.insertNode(newNode)
}