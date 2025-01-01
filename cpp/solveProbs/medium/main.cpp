#include <algorithm>
#include <iostream>
#include <map>
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
vector<int> twoSum(vector<int> &nums, int target) {
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
}

// Optimal solution (2 pointers)
bool twoSum2(vector<int> &nums, int target) {
  int n = nums.size();
  sort(nums.begin(), nums.end());
  int left = 0, right = n - 1;
  long long sum = nums[left];
  while (left <= right) {
    sum = nums[left] + nums[right];
    if (sum > target) {
      right--;
    }
    if (sum == target) {
      return true;
    }
    if (sum < target) {
      left++;
    }
  }
  return false;
}
