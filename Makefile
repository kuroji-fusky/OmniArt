.PHONY: install client-install server-install pnpm-install

install: pnpm-install client-install server-install

pnpm-install:
	@if ! command -v pnpm >/dev/null 2>&1; then \
		echo "PNPM not found, installing rn"; \
		npm install -g pnpm; \
	else \
		echo "PNPM is installed on your system"; \
	fi

client-install:
	cd client && pnpm install

server-install:
	cd server && go mod download