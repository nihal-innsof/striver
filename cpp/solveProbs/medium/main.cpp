#include <algorithm>
#include <cmath>
#include <iostream>
#include <map>
#include <ostream>
#include <vector>

using namespace std;

vector<int> twoSum(vector<int> &nums, int target);

void sortColors(vector<int> &nums);

int majorityElement(vector<int> &nums);

int main() {
  // vector<int> nums = {2, 7, 11, 15}; int target = 9;
  // vector<int> nums = {3, 2, 4}; int target = 6;
  // vector<int> nums = {3, 3}; int target = 6;
  /* vector<int> nums = {-1, -2, -3, -4, -5};
  int target = -8;
  vector<int> result = twoSum(nums, target);
  for (int i = 0; i < result.size(); i++) {
    cout << result[i] << " ";
  }
  cout << endl; */

  /* vector<int> nums = {2, 0, 2, 1, 1, 0};
  sortColors(nums);
  for (int i = 0; i < nums.size(); i++) {
    cout << nums[i] << " ";
  }
  cout << endl; */

  // vector<int> nums = {3, 2, 3};
  vector<int> nums = {8, 8, 7, 7, 7};
  cout << majorityElement(nums) << endl;
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

void sortColors(vector<int> &nums) {}

int majorityElement(vector<int> &nums) {
  int n = nums.size();
  for (int i = 0; i < n; i++) {
    int maxCount = floor(n / 2.0);
    for (int j = i + 1; j < n; j++) {
      if (nums[i] == nums[j]) {
        maxCount -= maxCount;
      }
      if (maxCount <= 0) {
        return nums[i];
      }
    }
  }
  return 0;
}
