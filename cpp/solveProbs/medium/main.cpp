#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

vector<int> twoSum(vector<int> &nums, int target);

int main() {
  // vector<int> nums = {2, 7, 11, 15}; int target = 9;
  // vector<int> nums = {3, 2, 4}; int target = 6;
  // vector<int> nums = {3, 3}; int target = 6;
  vector<int> nums = {-1, -2, -3, -4, -5};
  int target = -8;
  vector<int> result = twoSum(nums, target);
  for (int i = 0; i < result.size(); i++) {
    cout << result[i] << " ";
  }
  cout << endl;
  return 0;
}

// Better approach (Hashing)
/* vector<int> twoSum(vector<int> &nums, int target) {
  int n = nums.size();
  map<int, int> x;
  vector<int> result;
  for (int i = 0; i < n; i++) {
    int reminder = target - nums[i];

    if (x.find(reminder) != x.end()) {
      return {x[reminder], i};
    }

    x[nums[i]] = i;
  }
  return {};
} */

// Optimal solution (2 pointers)
vector<int> twoSum(vector<int> &nums, int target) {
  sort(nums.begin(), nums.end());
  return nums;
}
