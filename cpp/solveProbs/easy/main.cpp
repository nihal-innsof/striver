#include <iostream>
#include <map>
#include <vector>

int getLongestSubArray(std::vector<int> arr, int k);

int main() {
  // std::cout << getLongestSubArray({2, 3, 5, 1, 9}, 10);
  std::cout << getLongestSubArray({-59, -25, 58, -59, -25, 58}, -85);
  // std::cout << getLongestSubArray( {94, -33, -13, 40, -82, 94, -33, -13, 40,
  // -82}, 52); std::cout << getLongestSubArray({783, 0, 615, 868, 783, 0, 615,
  // 868}, 2266);
  return 0;
}

// Better approach (Using hashing) (Suitable for both positive and negative)
int getLongestSubArray(std::vector<int> arr, int k) {
  int n = arr.size();

  std::map<long long, int> preSumMap;
  long long sum = 0;
  int maxLen = 0;
  for (int i = 0; i < n; i++) {
    sum += arr[i];

    if (sum == k) {
      maxLen = std::max(maxLen, i + 1);
    }

    long long rem = sum - k;

    if (preSumMap.find(rem) != preSumMap.end()) {
      int len = i - preSumMap[rem];
      maxLen = std::max(maxLen, len);
    }

    if (preSumMap.find(sum) == preSumMap.end()) {
      preSumMap[sum] = i;
    }
  }
  return maxLen;
}

// Better approach (Using two pointers) (Suitable for positive numbers only)
/* int getLongestSubArray(std::vector<int> arr, int k) {
  int n = arr.size();
  int result = 0;
  int right, left;
  right = 0;
  left = 0;
  long long sum = arr[left];
  while (right < n) {
    while (left <= right && sum > k) {
      sum -= arr[left];
      left++;
    }

    if (sum == k) {
      result = std::max(result, right - left + 1);
    }
    right++;
    if (right < n) {
      sum += arr[right];
    }
  }
  return result;
} */
