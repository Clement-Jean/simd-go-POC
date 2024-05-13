SHELL=bash

record:
	@pushd go &&									\
	git diff HEAD -O ../.gitorderfile						\
				  src/cmd/compile/internal/ssagen/ssa.go 		\
				  src/cmd/compile/internal/ssa/opGen.go  		\
				  src/cmd/compile/internal/ssa/_gen/genericOps.go 	\
		>> ../patches/00_common_test.diff &&					\
	git diff HEAD src/cmd/compile/internal/arm64/ssa.go 				\
				  src/cmd/compile/internal/ssa/rewriteARM64.go 		\
				  src/cmd/compile/internal/ssa/_gen/ARM64.rules 	\
				  src/cmd/compile/internal/ssa/_gen/ARM64Ops.go 	\
		>> ../patches/01_arm64_test.diff &&					\
	git diff HEAD src/cmd/compile/internal/amd64/ssa.go 				\
				  src/cmd/compile/internal/ssa/rewriteAMD64.go 		\
				  src/cmd/compile/internal/ssa/_gen/AMD64.rules 	\
				  src/cmd/compile/internal/ssa/_gen/AMD64Ops.go 	\
		>> ../patches/02_amd64_test.diff &&					\
	popd
