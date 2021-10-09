
class SummaryRanges {
public:
    SummaryRanges() {
        tree_.insert(make_pair(-5, -5));
        tree_.insert(make_pair(10005, 10005));
    }

    void addNum(int val) {
       auto ceil = tree_.lower_bound(val);  // val <= ceil->val
       auto floor = prev(ceil);  // val > floor->val

       // cout << val << ",, " << floor->first << ", " << floor->second << ",, " << ceil->first << "," << ceil->second << endl;
       if (val >= floor->first && val <= floor->second || val >= ceil->first && val <= ceil->second) {}
       else if (val == floor->second + 1 && val == ceil->first - 1) {
           floor->second = ceil ->second;
           tree_.erase(ceil);
       } else if (val == floor->second + 1) {
           floor->second += 1;
       } else if (val == ceil->first - 1) {
           auto p = make_pair(ceil->first - 1, ceil->second);
           tree_.erase(ceil);
           tree_.insert(p);
       } else {
           tree_.insert(make_pair(val, val));
       }
    }

    vector<vector<int>> getIntervals() {
        vector<vector<int>> res;
        int i = 0;
        for (auto it = tree_.begin(); it != tree_.end(); it++, i++) {
            if (i >= 1 && i < tree_.size() - 1) {
                res.push_back({it->first, it->second});
            }
        }
        return res;
    }

    private:
    map<int, int> tree_;
};

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * SummaryRanges* obj = new SummaryRanges();
 * obj->addNum(val);
 * vector<vector<int>> param_2 = obj->getIntervals();
 */