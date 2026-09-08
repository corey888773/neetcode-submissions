# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        prev = None
        curr = head

        while curr is not None:
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next

        dummy.next = prev

        return dummy.next

#  None -> head -> next -> next.next

# temp = head.next


# head.next = prev
#  # P None <- head -> t next -> next.next


# prev = head
# head = temp

#  # None <- P head -> h next -> next.next

# temp = head.next

#  # None <- P head -> h next -> t None
#  # None <- P head <- h next -> t None
#  # None <-  head <- P next -> h None






