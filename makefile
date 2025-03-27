hot:
	docker compose -f docker-compose.hot.yaml up

prepare:
	set VITE_BUILT=1 && pnpm --dir ./frontend install && pnpm --dir ./frontend build && robocopy "./frontend/dist" "./backend/pkg/embed/dist" /E

refresh-frontend:
	docker start vocapedia-frontend
	docker exec -it vocapedia-frontend sh -c "pnpm install"

mobile:
	pnpm --dir ./frontend/ exec cap run android && pnpm exec cap run android
