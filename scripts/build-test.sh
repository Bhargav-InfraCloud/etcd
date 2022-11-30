#!/usr/bin/env bash

source ./scripts/test_lib.sh

# Test to check whether "make build" reads and prepends build flags from env GO_BUILD_FLAGS
function test_makefile_read_build_flags_from_env {
    # 3 occurrences, one for each: etcd, etcdctl, etcdutl 
    # Build command on output should print flags -tags (from env GO_BUILD_FLAGS) and -v (from Makefile)
    local occurrences=3

    # Build flag -tags without arguments is same as default build, handy for testing
    local extra_flag="-tags"

    # Build command is directed to stdout for counting occurrences
    # Order of flags should be same as that in env GO_BUILD_FLAGS (as prepended) + that in Makefile
    local count
    count=$(GO_BUILD_FLAGS="${extra_flag}" make build 2>&1 | grep -c "'go' 'build' '${extra_flag}' '-v'")

    if [ "$count" -ne "$occurrences" ]; then
        # Failure case: Number of occurrences doesn't match with expected
        log_error "Error: Build flags from env GO_BUILD_FLAGS not found in one or more logs (matches expected: ${occurrences}, got: ${count})"

        return 255
    fi

    log_success "SUCCESS: Build flags from env GO_BUILD_FLAGS read successfully"
    return 0
}

# Test to check whether "make build" reads and prepends build flags from env GO_BUILD_FLAGS
function test_makefile_read_build_flags_from_env_with_args {
    # 3 occurrences, one for each: etcd, etcdctl, etcdutl 
    # Build command on output should print flags -tags (from env GO_BUILD_FLAGS) and -v (from Makefile)
    local occurrences=3

    local extra_flag="-tags"
    local arg1='linux'

    # Build command is directed to stdout for counting occurrences
    # Order of flags should be same as that in env GO_BUILD_FLAGS (as prepended) + that in Makefile
    local count 
    count=$(GO_BUILD_FLAGS="${extra_flag} ${arg1}" make build 2>&1 | grep -c "'go' 'build' '${extra_flag}' '${arg1}' '-v'")

    if [ "$count" -ne "$occurrences" ]; then
        # Failure case: Number of occurrences doesn't match with expected
        log_error "Error: Build flags with args from env GO_BUILD_FLAGS not found in one or more logs (matches expected: ${occurrences}, got: ${count})"

        return 255
    fi

    log_success "SUCCESS: Build flags with args from env GO_BUILD_FLAGS read successfully"
    return 0
}

"$@"
