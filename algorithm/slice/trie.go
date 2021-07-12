/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:43 PM  4/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

/*
class Trie {
public:
    Trie(): root_(new TrieNode()) {}

    void insert(const string& word) {
        TrieNode* p = root_.get();
        for (const char c : word) {
            if (!p->children[c - 'a'])
                p->children[c - 'a'] = new TrieNode();
            p = p->children[c - 'a'];
        }
        p->is_word = true;
    }

    bool hasAllPrefixes(const string& word) {
        const TrieNode* p = root_.get();
        for (const char c : word) {
            if (!p->children[c - 'a']) return false;
            p = p->children[c - 'a'];
            if (!p->is_word) return false;
        }
        return true;
    }
private:
    struct TrieNode {
        TrieNode():is_word(false), children(26, nullptr){}

        ~TrieNode() {
            for (auto node : children)
                delete node;
        }

        bool is_word;
        vector<TrieNode*> children;
    };

    std::unique_ptr<TrieNode> root_;
};
 */