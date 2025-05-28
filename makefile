hot:
	docker compose -f docker-compose.hot.yaml up

test:
	docker compose -f docker-compose.test.yaml up

app-prepare:
	git submodule update --init --recursive

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
	rsync -av --delete --progress --exclude 'frontend' --exclude '.git' --exclude '.vscode' --exclude 'backend/pkg/whisper' --exclude 'backend/tmp' ~/Desktop/vocapedia/ root@93.115.20.7:/opt/vocapedia/

whisper-build:
	cd backend/pkg/whisper && make

