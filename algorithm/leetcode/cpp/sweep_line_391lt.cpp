//
// Created by 裴星鑫 on 2021/11/16.
//

#include <algorithm>
#include <map>
#include <vector>

using namespace std;

class Solution {
 public:
  bool isRectangleCover(vector<vector<int>>& rectangles) {
    return simulation(rectangles);
  }

  bool simulation(vector<vector<int>>& rectangles) {
    typedef pair<int, int> Point;
    int min_x = rectangles[0][0], min_y = rectangles[0][1];
    int max_x = rectangles[0][2], max_y = rectangles[0][3];
    long area = 0;
    map<Point, int> cache;

    for (int i = 0; i < rectangles.size(); i++) {
      min_x = min(min_x, rectangles[i][0]);
      min_y = min(min_y, rectangles[i][1]);
      max_x = max(max_x, rectangles[i][2]);
      max_y = max(max_y, rectangles[i][3]);

      area += (rectangles[i][2] - rectangles[i][0]) * (rectangles[i][3] - rectangles[i][1]);

      Point x = {rectangles[i][0], rectangles[i][1]};
      Point x1 = {rectangles[i][2], rectangles[i][1]};
      Point y = {rectangles[i][2], rectangles[i][3]};
      Point y1 = {rectangles[i][0], rectangles[i][3]};
      cache[x]++;
      cache[x1]++;
      cache[y]++;
      cache[y1]++;
    }

    Point minx_miny = {min_x, min_y};
    Point minx_maxy = {min_x, max_y};
    Point maxx_miny = {max_x, min_y};
    Point maxx_maxy = {max_x, max_y};
    if (area != long(max_x - min_x) * (max_y - min_y) || cache[minx_miny] != 1 || cache[minx_maxy] != 1 ||
    cache[maxx_miny] != 1 || cache[maxx_maxy] != 1) {
      return false;
    }

    cache.erase(minx_maxy);
    cache.erase(minx_miny);
    cache.erase(maxx_maxy);
    cache.erase(maxx_miny);

    for (auto it = cache.begin(); it != cache.end(); it++) {
      if (it->second != 2 && it->second != 4) return false;
    }
    return true;
  }
};