class Solution {
public:
    vector<vector<int>> getSkyline(vector<vector<int>>& buildings) {
        multiset<pair<int, int>> dots;
        for (auto& d : buildings) {
            dots.emplace(make_pair(d[0], -d[2]));
            dots.emplace(make_pair(d[1], d[2]));
        }

        vector<vector<int>> r;
        vector<int> last = {0, 0};
        multiset<int> heights({0});

        for (auto d : dots) {
            if (d.second < 0) heights.emplace(-d.second);
            else heights.erase(heights.find(d.second));

            auto max_height = *heights.rbegin();
            if (max_height != last[1]) {
                last[0] = d.first;
                last[1] = max_height;
                r.push_back(last);
            }
        }
        return r;
    }
};