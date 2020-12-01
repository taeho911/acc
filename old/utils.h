#pragma once

#include<iostream>
#include<vector>

namespace utils {
	std::vector<std::string> optsplit(char* str, char delim);
	int indexof(char* str, char token);
	bool isin(char* str, const char* token);
}