//
// Created by 裴星鑫 on 2021/11/2.
//

class BrowserHistory {
 public:
  BrowserHistory(string homepage) {
    back_.push_back(homepage);
  }

  void visit(string url) {
    forward_.clear();
    back_.push_back(url);
  }

  string back(int steps) {
    int count = 0;
    while (count++ < steps && back_.size() >= 2) {
      forward_.push_back(back_.back());
      back_.pop_back();
    }
    return back_.back();
  }

  string forward(int steps) {
    if (forward_.empty()) return back_.back();
    string current;
    int count = 0;
    while (count++ < steps && !forward_.empty()) {
      current = forward_.back();
      back_.push_back(current);
      forward_.pop_back();
    }
    return current;
  }

  vector<string> back_;
  vector<string> forward_;
};

