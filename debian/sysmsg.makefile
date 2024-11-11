# Ubuntu unconditionally add ELF_PACKAGE_METADATA in dpkg-buildpacakge,
# which causes ld to add extra .note.package sections.
unexport ELF_PACKAGE_METADATA

ARCH := amd64

# sync with pkg/sentry/platform/systrap/sysmsg/build.bzl

CFLAGS := -Wall -Wno-unused-command-line-argument
CFLAGS += -fpie

ifeq ($(ARCH),amd64)
	CFLAGS += -O2
endif
ifeq ($(ARCH),arm64)
	CFLAGS += -O1 -mno-outline-atomics
endif
CFLAGS += -fno-builtin
CFLAGS += -ffreestanding
CFLAGS += -mgeneral-regs-only
CFLAGS += -g
CFLAGS += -Wa,--noexecstack
CFLAGS += -fno-asynchronous-unwind-tables
CFLAGS += -fno-stack-protector

%.o: %.S
	$(CC) $(CFLAGS) -c $^ -o $@

%.o: %.c
	$(CC) $(CFLAGS) -c $^ -o $@

sighandler.built-in.bin.o: sighandler_$(ARCH).o syshandler_$(ARCH).o sigrestorer_$(ARCH).o sysmsg_lib.o
	$(LD) -pie -z noexecstack -T pie.lds.S $^ -o $@

sighandler_$(ARCH).go: sighandler.built-in.bin.o
	bash gen_offsets_go.sh sighandler Sighandler $^ > $@

sighandler.built-in.$(ARCH).bin: sighandler.built-in.bin.o
	 objcopy -O binary $^ $@

all: sighandler.built-in.$(ARCH).bin sighandler_$(ARCH).go

clean:
	$(RM) *.o

distclean: clean
	$(RM) *.bin sighandler_*.go
