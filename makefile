# Hot Komutu (Windows için)
hot:
	docker compose -f docker-compose.hot.yaml up

# Frontend Hazırlık (Yükleme ve Derleme)
prepare:
	pnpm --dir ./frontend/ install >> /dev/null && pnpm --dir ./frontend/ build >> /dev/null && rsync -a ./frontend/dist/ ./backend/pkg/embed/dist

# Frontend Konteynerini Başlat
refresh-frontend:
	docker start vocapedia-frontend
	docker exec -it vocapedia-frontend sh -c "pnpm install"

# Mobil Uygulama Çalıştırma
mobile:
	pnpm --dir ./frontend/ exec cap run android && pnpm exec cap run android
