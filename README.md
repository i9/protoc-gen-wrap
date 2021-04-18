# protoc-gen-wrap
generate helper wrap struct and sql scan/value

# Motivation
- bytes field in .proto often stores value of big.Int, Address or Hash.
- Save serialized pb msg to db BYTEA column and read back is common and repeated manual work
