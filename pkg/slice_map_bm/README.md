# Slice and map benchmarking

## Run

clear; go test -bench='.' -run='^#' -count=1 -benchtime 10s -gcflags "-N -m -l"

clear; go test -bench='.' -run='^#' -count=10 -benchtime 10s -timeout 30m -gcflags "-N -m -l" 2>result_10.deb | tee result_10.txt
benchstat -row '/size,.fullname' -col '.config' -format csv result_10.txt >result_10.csv
