frontend:
	export VITE_BUILT=1 && pnpm --dir ../frontend install && pnpm --dir ../frontend build && rsync -av --delete ../frontend/dist/ ./pkg/embed/dist/

local:
	export VITE_BUILT=1 && air -c .air.toml
