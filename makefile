hot: 
	sudo docker compose -f docker-compose.hot.yml up

prepare:
	pnpm --dir ./frontend/ install >> /dev/null && pnpm --dir ./frontend/ build >> /dev/null && rsync -a ./frontend/dist/ ./backend/pkg/embed/dist

refresh-frontend:
	sudo docker start vocapedia-frontend
	sudo docker exec -it vocapedia-frontend sh -c "pnpm install"

hot-mobile:
	pnpm --dir ./frontend/ exec cap run android