class Solution {
public:
    vector<string> fizzBuzz(int n) {
        std::vector<std::string> result(n);

        for (int i = 1; i <= n; i++)
        {
            if (i % 5 != 0 && i % 3 != 0)
            {
                result[i - 1] = std::to_string(i);
            }
            
            if (i % 3 == 0)
            {
                result[i - 1] = "Fizz";
            }

            if (i % 5 == 0)
            {
                result[i - 1] = "Buzz";
            }

            if (i % 3 == 0 && i % 5 == 0)
            {
                result[i - 1] = "FizzBuzz";
            }
        }

        return result;
    }
};
