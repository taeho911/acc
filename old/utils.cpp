#include "utils.h"

namespace utils {
	std::vector<std::string> optsplit(char* str, char delim) {
		std::vector<std::string> strarr;
		strarr.reserve(2);
		int delimIndex = -1;
		int strlen = -1;

		while (str[++strlen] != '\0') {
			if (str[strlen] == delim) {
				delimIndex = strlen;
				break;
			}
		}
		
		if (delimIndex == -1) {
			strarr.push_back(std::string(str));
			return strarr;
		}

		std::string str1(str, delimIndex);
		std::string str2 = "";
		int index = delimIndex + 1;
		
		while (str[index] != '\0') {
			str2 += str[index++];
		}

		strarr.push_back(str1);
		strarr.push_back(str2);

		return strarr;
	}

	int indexof(char* str, char token) {
		int index = -1;
		int i = -1;

		while (str[++i] != '\0') {
			if (str[i] == token) {
				index = i;
				break;
			}
		}

		return index;
	}

	bool isin(char* str, const char* token) {
		int strindex = 0;
		int tokenindex = 0;

		while (str[strindex] != '\0' && token[tokenindex] != '\0') {
			if (str[strindex] == token[tokenindex]) {
				strindex++;
				tokenindex++;
			}
			else {
				strindex++;
				tokenindex = 0;
			}
		}

		if (token[tokenindex] == '\0') return true;
		else return false;
	}
}