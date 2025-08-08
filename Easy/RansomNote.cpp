class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        std::unordered_map<char, int> dictionary;

        for (char character : magazine)
        {
            if (dictionary.find(character) == dictionary.end())
            {
                dictionary[character] = 1;
            } 
            else
            {
                dictionary[character]++;
            }
        }

        for (char character : ransomNote)
        {
            if (dictionary.find(character) != dictionary.end() && dictionary[character] > 0)
            {
                dictionary[character]--;
            }
            else
            {
                return false;
            }
        }

        return true;
    }
};
