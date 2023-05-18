go test -bench . -tags base -count 10 | tee x_base.txt~
go test -bench . -tags cast -count 10 | tee x_cast.txt~
go test -bench . -tags base,nextslicecap -count 10 | tee x_base_next.txt~
go test -bench . -tags cast,nextslicecap -count 10 | tee x_cast_next.txt~

benchstat x_base.txt~ x_cast.txt~ x_base_next.txt~ x_cast_next.txt~