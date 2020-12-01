#include "netacc_ls.h"
#include "utils.h"

char helpStr_ls[] = R"(
문법:
  netacc ls [ options ] [ --help ]

options:
  -a, --all		회원정보 전체 출력
  -i, --index	해당 색인의 회원정보의 상세를 출력
				ex) -i=13 or --index=13
  -h, --help	도움말 출력
)";

int lsentry(int argc, char* argv[], int index) {
	index++;
	char* cmd = argv[index];

	bool aflag = false;
	bool iflag = false;
	bool hflag = false;
	bool dumpflag = false;
	int aindex = 0;
	int iindex = 0;
	int hindex = 0;
	std::string ival;

	for (int i = index; i < argc; i++) {
		if (!strcmp(argv[i], "-a") || !strcmp(argv[i], "--all")) {
			aflag = true;
			aindex = i;
		}
		else if (utils::isin(argv[i], "-i") || utils::isin(argv[i], "--index")) {
			iflag = true;
			iindex = i;
			if (utils::indexof(argv[i], '=') != -1) {
				ival = utils::optsplit(argv[i], '=')[1];
			}
		}
		else if (!strcmp(argv[i], "-h") || !strcmp(argv[i], "--help")) {
			hflag = true;
			hindex = i;
		}
		else if (iflag && i == iindex + 1 && ival.empty()) {
			ival = std::string(argv[i]);
		}
		else {
			dumpflag = true;
		}
	}

	if (hflag) {
		std::cout << helpStr_ls << std::endl;
		return 0;
	}

	if (aflag) {
		return 0;
	}

	if (iflag) {
		return 0;
	}

	return 0;
}