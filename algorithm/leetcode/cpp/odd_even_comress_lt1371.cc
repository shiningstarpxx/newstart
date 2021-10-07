
class Solution {
public:
    int findTheLongestSubstring(string s) {
        vector<int> cache(1 << 5, -1);
        cache[0] = 0;
        int status = 0;
        int ans = 0;
        for (int i = 0; i < s.length(); i++) {
            char c = s[i];
            if (c == 'a') status ^= 1;
            else if (c == 'e') status ^= 1 << 1;
            else if (c == 'i') status ^= 1 << 2;
            else if (c == 'o') status ^= 1 << 3;
            else if (c == 'u') status ^= 1 << 4;

            if (cache[status] != -1) {
                ans = max(ans, i + 1 - cache[status]);
                // cout << status << "," << ans << endl;
            } else {
                cache[status] = i + 1;
            }
        }
        return ans;
    }
};