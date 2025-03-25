pnpm-check:
    if ! command -v pnpm >/dev/null 2>&1; then
        echo "PNPM not found, installing rn"
        npm install -g pnpm
    else
        echo "PNPM is installed on your system"
    fi

# Install
client-install:
    cd client && pnpm install

server-install:
    cd server && go mod download

# Build
client-build:
  cd client && pnpm build

server-build:
  cd server && go build cmd/server

alias i := install
alias b := build

install:
    pnpm-check client-install server-install

build:
    pnpm-check client-build server-build

# Test stuff WIP