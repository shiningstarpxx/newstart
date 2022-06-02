//
// Created by 裴星鑫 on 2022/6/2.
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
    TreeNode* deleteNode(TreeNode* root, int key) {
        if (root == nullptr) return nullptr;
        if (root->val < key) {
            root->right = deleteNode(root->right, key);
            return root;
        }
        if (root->val > key) {
            root->left = deleteNode(root->left, key);
            return root;
        }
        if (root->val == key) {
            if (root->left == nullptr && root->right == nullptr) return nullptr;
            if (root->left == nullptr) return root->right;
            if (root->right == nullptr) return root->left;
            TreeNode* t = root->right;
            while (t->left) t = t->left;
            // passed
            // root->right = deleteNode(root->right, t->val);
            // t->right = root->right;
            // t->left = root->left;
            // return t;

            // try
            t->left = root->left;
            return root->right;
        }
        return root;
    }
};