class Solution {
public:
    int findNumbers(vector<int>& nums) {
        int counter{ 0 };

        for (int i = 0; i < nums.size(); i++)
        {
            std::string numberString{ std::to_string(nums[i]) };

            if (numberString.length() % 2 == 0)
            {
                counter++;
            }
        }

        return counter;
    }
};
