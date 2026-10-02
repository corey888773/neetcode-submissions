from dataclasses import dataclass, field
from collections import defaultdict

@dataclass
class TrieNode:
    nodes: dict[str, "TrieNode"] = field(default_factory=dict)
    is_end: bool = False

class PrefixTree:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        curr = self.root
        for char in word:
            if char not in curr.nodes:
                curr.nodes[char] = TrieNode()

            curr = curr.nodes[char]

        curr.is_end = True

    def search(self, word: str) -> bool:
        curr = self.root
        for char in word:
            if char not in curr.nodes:
                return False

            curr = curr.nodes[char]

        return curr.is_end

    def startsWith(self, prefix: str) -> bool:
        curr = self.root
        for char in prefix:
            if char not in curr.nodes:
                return False

            curr = curr.nodes[char]

        return True
        
        