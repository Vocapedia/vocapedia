hot:
	docker compose -f docker-compose.hot.yaml up

prepare:
	pnpm --dir ./frontend/ install >> /dev/null && pnpm --dir ./frontend/ build >> /dev/null && rsync -a ./frontend/dist/ ./backend/pkg/embed/dist

refresh-frontend:
	docker start vocapedia-frontend
	docker exec -it vocapedia-frontend sh -c "pnpm install"

mobile:
	pnpm --dir ./frontend/ exec cap run android && pnpm exec cap run android
