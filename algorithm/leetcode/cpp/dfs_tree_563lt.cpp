//
// Created by 裴星鑫 on 2021/11/18.
//

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
 public:
  int findTilt(TreeNode* root) {
    dfs(root);
    return sum_;
  }

  int dfs(TreeNode* root) {
    if (root == nullptr) return 0;
    int left = dfs(root->left);
    int right = dfs(root->right);
    sum_ += abs(left - right);
    return root->val + left + right;
  }

  int sum_ = 0;
};
