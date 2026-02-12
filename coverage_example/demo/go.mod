module hello.test

go 1.25.0

replace ext => ./hack/ext
replace ext/sub => ./hack/extsub

require ext v0.0.0-00010101000000-000000000000 // indirect
require ext/sub v0.0.0-00010101000000-000000000000 // indirect
