class Solution {
public:
    int maximumWealth(vector<vector<int>>& accounts) 
    {
        int richest{ 0 };
        int accountSum{ 0 };

        for (int i = 0; i < accounts.size(); i++)
        {
            for (int j = 0; j < accounts[i].size(); j++)
            {
                accountSum += accounts[i][j];
            }

            if (accountSum > richest)
            {
                richest = accountSum;
            }

            accountSum = 0;
        }    

        return richest;
    }
};
