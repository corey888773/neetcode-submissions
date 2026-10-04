from dataclasses import dataclass, field

@dataclass
class TreeNode:
    nodes: dict[str, "TreeNode"] = field(default_factory=dict)
    is_end: bool = False

class WordDictionary:
    def __init__(self):
        self.root = TreeNode()
        self.end = set()

    def addWord(self, word: str) -> None:
        curr = self.root
        for letter in word:
            curr = curr.nodes.setdefault(letter, TreeNode())

        curr.is_end = True

    def search(self, word: str) -> bool:
        return self._search(self.root, word, 0)

    def _search(self, node: TreeNode, word: str, idx: str) -> bool:
        curr = node
        for i in range(idx, len(word)):
            letter = word[i]
            if letter == ".":
                for option in curr.nodes.keys():
                    if self._search(curr.nodes[option], word, i+1):
                        return True
                return False
            else:
                if letter not in curr.nodes:
                    return False

                curr = curr.nodes[letter]

        return curr.is_end

# Your WordDictionary object will be instantiated and called as such:
# obj = WordDictionary()
# obj.addWord(word)
# param_2 = obj.search(word)