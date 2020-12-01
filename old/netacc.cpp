#include "netacc.h"
#include "netacc_ls.h"
#include "utils.h"

#define TAEHO 1

#if TAEHO
static uint32_t sAllocCount = 0;

void* operator new(size_t size) {
	sAllocCount++;
	std::cout << "Heap alloc! " << size << " bytes\n";
	return malloc(size);
}
#endif

char helpStr[] = R"(
문법:
  netacc [ actions ] [ options ] [ --help ]

actions:
  ls			회원정보조회
  add			회원정보추가
  del			회원정보삭제
  mod			회원정보수정

options:
  -h, --help	도움말 출력
)";

int main(int argc, char* argv[]) {
    auto start = std::chrono::high_resolution_clock::now();
	
	if (argc < 2) {
		std::cout << helpStr << std::endl;
		return 0;
	}

	int index = 1;
	char* cmd = argv[index];
	if (!strcmp(cmd, "ls")) {
		lsentry(argc, argv, index);
	}
	else if (!strcmp(cmd, "add")) {

	}
	else if (!strcmp(cmd, "del")) {

	}
	else if (!strcmp(cmd, "mod")) {

	}
	else {
		std::cout << helpStr << std::endl;
	}
	
	auto end = std::chrono::high_resolution_clock::now();
	auto duration = std::chrono::duration_cast<std::chrono::microseconds>(end - start);
#if TAEHO
	std::cout << "----------------------------------------------------" << std::endl;
	std::cout << "Execution time: " << duration.count() << " microseconds" << std::endl;
	std::cout << "Heap alloc count: " << sAllocCount << std::endl;
#endif

	return 0;
}

// Run program: Ctrl + F5 or Debug > Start Without Debugging menu
// Debug program: F5 or Debug > Start Debugging menu

// Tips for Getting Started: 
//   1. Use the Solution Explorer window to add/manage files
//   2. Use the Team Explorer window to connect to source control
//   3. Use the Output window to see build output and other messages
//   4. Use the Error List window to view errors
//   5. Go to Project > Add New Item to create new code files, or Project > Add Existing Item to add existing code files to the project
//   6. In the future, to open this project again, go to File > Open > Project and select the .sln file
