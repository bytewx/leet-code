#include <iostream>
#include <vector>
#include <cmath>

int main()
{
    std::vector<int> nums{ 1, 1, 1, 0, 1 };

    int oneCounter = 0;
    int maxOnes = 0;

    for (int i = 0; i < nums.size(); i++) {
        if (nums[i] == 1) {
            oneCounter++;
            maxOnes = std::max(maxOnes, oneCounter);
        }
        else {
            oneCounter = 0;
        }
    }

    return maxOnes;
}
