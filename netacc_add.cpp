#include "netacc_add.h"

char helpStr_add[] = R"(
문법:
  netacc add [ options ] [ --help ]

options:
  -a, --all		회원정보 전체 출력
  -i, --index	해당 색인의 회원정보의 상세를 출력
				ex) -i=13 or --index=13
  -h, --help	도움말 출력
)";