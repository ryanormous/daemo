
Json = configuration.json

LDFLAGS = -ldflags="-s -w"

# ROOT PATH
PREFIX = $(shell jq '.Root' $(Json) | sed 's/"//g')

# USER AND GROUP
User = www-data
Group = www-data

# DIRECTORIES
Dirs = $(addprefix $(PREFIX)/,bin conf log run systemd)

# CONFIGURATION FILE
Conf = $(PREFIX)/conf/$(Json)

# daemo EXECUTABLE
Exe = $(PREFIX)/bin/daemo

# SYSTEMD SERVICE FILE
Srv = $(PREFIX)/systemd/daemo.service

# SYSTEMD UNIT DIRECTORY
UnitDir = $(shell pkg-config --variable=systemdsystemunitdir systemd)

all: Makefile $(Dirs) $(Conf) $(Exe) $(Srv)
.PHONY: all

Makefile: ;

$(Dirs):
	mkdir -p $@
	chown $(User):$(Group) $@

$(Conf):
	cp $(Json) $@
	chown $(User):$(Group) $@

$(Exe):
	go build $(LDFLAGS) -o $@
	go clean
	chown $(User):$(Group) $@

$(Srv):
	cp daemo.service $@
	chown $(User):$(Group) $@
	ln -s $@ $(UnitDir)/daemo.service
	systemctl enable daemo.service
	systemctl start daemo.service

