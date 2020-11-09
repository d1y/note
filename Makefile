dev_db := cache/dev.db
backup_sql := note.sql

prod := note.db

build_src := build/note.app

clean:
	@echo "clean build file"
	rm -rf build

dev-db:
	@echo "create dev db"
	rm -rf $(dev_db)
	cat $(backup_sql) | sqlite3 $(dev_db)

build_prod:
	@echo "try build prodution bin file"
	rm -rf build
	mkdir build
	cat $(backup_sql) | sqlite3 build/note.db
	cp -rf static build/static
	cp -rf template build/template
	go build -o $(build_src) -ldflags "-linkmode external -extldflags -static" .
	# @echo "use upx"
	# upx -9 $(build_src)
	