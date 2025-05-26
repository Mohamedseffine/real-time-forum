rm -rf database/database.db-shm
rm -rf database/database.db-wal
git add .
git commit -m "$1" 
git push