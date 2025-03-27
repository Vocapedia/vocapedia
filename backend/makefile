frontend:
	set VITE_BUILT=1 && pnpm --dir ../frontend install && pnpm --dir ../frontend build && robocopy "../frontend/dist" "./pkg/embed/dist" /E

local:
	set VITE_BUILT=1 && air -c .air.toml
