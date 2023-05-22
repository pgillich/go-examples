# Slice and map benchmarking

## Run

clear; go test -bench='.' -run='^#' -count=1 -benchtime 10s -gcflags "-N -m -l" 2>result_1.deb | tee result_1.txt

clear; go test -bench='.' -run='^#' -count=5 -benchtime 10s -gcflags "-N -m -l" 2>result_5.deb | tee result_5.txt

benchstat -filter '.name:/Fill/' -col /size result_5.txt > result_5_Fill.txt
benchstat -filter '.name:/Seek/' -col /size result_5.txt > result_5_Seek.txt

clear; go test -bench='.' -run='^#' -count=10 -benchtime 10s -timeout 30m -gcflags "-N -m -l" 2>result_10.deb | tee result_10.txt
