## Trees

A tree data structure can be defined recursively (locally) as a collection of nodes (starting at a root node), where each node is a data structure consisting of a value, together with a list of references to nodes (the "children"), with the constraints that no reference is duplicated, and none points to the root. A tree is a type of graph.

**Tree Representation**: Trees can be represented using linked list where each parent/root node points to it's children.

An **ordered tree** is a rooted tree in which all the children of each vertex are ordered. It is convenient to assume that in a tree's diagram, all the children are ordered left to right.

A **binary tree** can be defined as an ordered tree in which every vertex has no more than two children and each child is designated as either a left child or a right child of its parent; a binary tree may also be empty.

**Note**: The height `h` of binary tree with `n` nodes is **&lfloor;log<sub>2</sub>n&rfloor; ≤ h ≤ n − 1**.

#### Binary Trees

**Full Binary Tree**: A Binary Tree is full if every node has 0 or 2 children. Following are examples of a full binary tree. We can also say a full binary tree is a binary tree in which all nodes except leaves have two children.

**Complete Binary Tree**: A Binary Tree is complete Binary Tree if all levels are completely filled except possibly the last level and the last level has all keys as left as possible.

**Perfect Binary Tree**: A Binary tree is Perfect Binary Tree in which all internal nodes have two children and all leaves are at the same level.

**Balanced Binary Tree**: A binary tree is balanced if the height of the tree is O(Log n) where n is the number of nodes. For Example, AVL tree maintains O(Log n) height by making sure that the difference between heights of left and right subtrees is 1.

**A degenerate (or pathological) tree**: A Tree where every internal node has one child. Such trees are performance-wise same as linked list.

**Binary Search Tree** is a node-based binary tree data structure which has the following properties:

- The left subtree of a node contains only nodes with keys lesser than the node’s key.
- The right subtree of a node contains only nodes with keys greater than the node’s key.
- The left and right subtree each must also be a binary search tree.

#### AVL Tree

**AVL Tree** is a self-balancing Binary Search Tree (BST) where the difference between heights of left and right subtrees cannot be more than 1 for all nodes and less than -1.

#### Red Black Tree

A red–black tree is a kind of self-balancing binary search tree in computer science. Each node of the binary tree has an extra bit, and that bit is often interpreted as the color (red or black) of the node. These color bits are used to ensure the tree remains approximately balanced during insertions and deletions.

#### Operations

1. Insert - Inserts an element in a tree/create a tree.
2. Search - Searches an element in a tree.
3. Preorder Traversal − Traverses a tree in a pre-order manner.
4. Inorder Traversal − Traverses a tree in an in-order manner.
5. Postorder Traversal − Traverses a tree in a post-order manner.

Traversals:

- BFS - Breadth First Search
- DFS - Depth First Search
  - Pre-order traversal
  - In-order traversal
  - Post-order traversal

These are the common operations that can be performed on a tree. Different types of trees has different operations which can be performed.

#### Applications

1. Heap is a tree data structure which is implemented using arrays and used to implement priority queues.
2. Syntax Tree: Used in Compilers.
3. B-Tree and B+ Tree : They are used to implement indexing in databases.
4. Manipulate hierarchical data.
5. Make information easy to search (see tree traversal).
6. Manipulate sorted lists of data.
7. Router algorithms

_References_:

- [https://www.geeksforgeeks.org/binary-tree-set-3-types-of-binary-tree/](https://www.geeksforgeeks.org/binary-tree-set-3-types-of-binary-tree/)
- [https://www.geeksforgeeks.org/avl-tree-set-1-insertion/](https://www.geeksforgeeks.org/avl-tree-set-1-insertion/)
- [https://en.wikipedia.org/wiki/Tree\_(data_structure)](<https://en.wikipedia.org/wiki/Tree_(data_structure)>)
- [https://en.wikipedia.org/wiki/AVL_tree](https://en.wikipedia.org/wiki/AVL_tree)

**Refer** [here](https://en.wikipedia.org/wiki/List_of_data_structures) for more information on types of trees and terminologies used.
