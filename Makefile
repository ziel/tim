# This flag is essential for vendoring with go < 1.6
# See: https://golang.org/cmd/go/#hdr-Vendor_Directories
#
export GO15VENDOREXPERIMENT=1

default: valid test install

test:
	go test $(glide novendor)

valid:
	go fmt
	go vet

install:
	go install
