// Copyright 2022 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// clang-format off
// Contains types needed by later headers.
#include <linux/types.h>
// clang-format on
#include <bpf/bpf_helpers.h>
#include <linux/bpf.h>

#define section(secname) __attribute__((section(secname), used))

char __license[] section("license") = "Apache-2.0";

// Note: bpf_helpers.h includes a struct definition for bpf_map_def in some, but
// not all, environments. Define our own equivalent struct to avoid issues with
// multiple declarations.
struct gvisor_bpf_map_def {
  unsigned int type;
  unsigned int key_size;
  unsigned int value_size;
  unsigned int max_entries;
  unsigned int map_flags;
};

// A map of RX queue number to AF_XDP socket. We only ever use one key: 0.
struct gvisor_bpf_map_def section("maps") sock_map = {
    .type = BPF_MAP_TYPE_XSKMAP,  // Note: "XSK" means AF_XDP socket.
    .key_size = sizeof(int),
    .value_size = sizeof(int),
    .max_entries = 1,
};

section("xdp") int xdp_prog(struct xdp_md *ctx) {
  // Lookup the socket for the current RX queue. Veth devices by default have
  // only one RX queue. If one is found, redirect the packet to that socket.
  // Otherwise pass it on to the kernel network stack.
  return bpf_redirect_map(&sock_map, ctx->rx_queue_index, XDP_PASS);
}
