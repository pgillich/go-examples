# Slice and map benchmarking

## Run

clear; go test -bench='.' -run='^#' -count=1 -benchtime 10s -gcflags "-N -m -l" 2>result_1.deb | tee result_1.txt

clear; go test -bench='.' -run='^#' -count=5 -benchtime 10s -gcflags "-N -m -l" 2>result_5.deb | tee result_5.txt
