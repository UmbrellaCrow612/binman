package global

// Contains a list of valid platform / operating systems that can be keys
// https://nodejs.org/api/process.html#processplatform
var ValidPlatforms = []string{
	"aix",
	"darwin",
	"freebsd",
	"linux",
	"openbsd",
	"sunos",
	"win32",
}

// Holds a list of valid architectures keys mapped from
// https://nodejs.org/api/process.html#processarch
var ValidArchitectures = []string{
	"arm",
	"arm64",
	"ia32",
	"loong64",
	"mips",
	"mipsel",
	"ppc64",
	"riscv64",
	"s390",
	"s390x",
	"x64",
}
