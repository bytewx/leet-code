class Solution {
public:
    int numberOfSteps(int num) {
        int steps{};

        while (num > 0)
        {
            for (int i = 0; i <= num; ++i)
            {
                if (num % 2 != 0)
                {
                    num -= 1;
                    steps += 1;
                }
                else
                {
                    num /= 2;
                    steps += 1;
                }

            }
        }

        return steps;
    }
};
