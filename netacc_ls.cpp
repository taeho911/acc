#include "netacc_ls.h"
#include "utils.h"

char helpStr_ls[] = R"(
����:
  netacc ls [ options ] [ --help ]

options:
  -a, --all		ȸ������ ��ü ���
  -i, --index	�ش� ������ ȸ�������� �󼼸� ���
				ex) -i=13 or --index=13
  -h, --help	���� ���
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

	std::cout << "[ aflag = " << aflag << " ]\n";
	std::cout << "[ iflag = " << iflag << " ]\n";
	std::cout << "[ hflag = " << hflag << " ]\n";
	std::cout << "[ dumpflag = " << dumpflag << " ]\n";
	std::cout << "[ aindex = " << aindex << " ]\n";
	std::cout << "[ iindex = " << iindex << " ]\n";
	std::cout << "[ hindex = " << hindex << " ]\n";
	std::cout << "[ ival = " << ival << " ]\n";
	
	return 0;
}