//
// Created by 裴星鑫 on 2021/11/21.
//

/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val) {
        val = _val;
    }

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
 public:
  int maxDepth(Node* root) {
    if (root == nullptr) return 0;
    dfs(root, 1);
    return height_;
  }

  void dfs(Node* root, int deep) {
    if (root == nullptr || root->children.empty()) {
      // cout << deep << endl;
      height_ = max(height_, deep);
      return;
    }

    for (int i = 0; i < root->children.size(); i++) {
      dfs(root->children[i], deep + 1);
    }
  }

  int height_ = 0;
};