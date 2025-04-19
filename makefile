hot:
	docker compose -f docker-compose.hot.yaml up
app:
	docker compose -f docker-compose.yaml up

docker-up:
	docker-compose build && docker-compose up -d

prepare:
	export VITE_BUILT=1 && pnpm --dir ./frontend install && pnpm --dir ./frontend build && rsync -av --delete ./frontend/dist/ ./backend/pkg/embed/dist/

refresh-frontend:
	docker start vocapedia-frontend
	docker exec -it vocapedia-frontend sh -c "pnpm install"

mobile:
	pnpm --dir ./frontend exec cap run android && pnpm exec cap run android

sync:
	rsync -av --delete --progress --exclude 'frontend' --exclude '.git' --exclude '.vscode' --exclude 'backend/tmp' ~/Desktop/projects/vocapedia/ root@93.115.20.7:/opt/vocapedia/
