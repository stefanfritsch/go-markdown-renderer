.PHONY: sass sass_build docker

sass/sass:
	mkdir -p sass
	curl -L https://github.com/sass/dart-sass/releases/download/1.55.0/dart-sass-1.55.0-linux-x64.tar.gz > sass/sass.tar.gz
	tar xvzf sass/sass.tar.gz
	rm sass/sass.tar.gz
	mv sass/dart-sass/sass sass/
	rm -r sass/dart-sass

sass: sass/sass
	sass/sass -w scss:assets/css/

sass_build: sass/sass
	sass/sass scss:assets/css/

run: sass_build
	TITLE="testtitle" go run main.go

build: sass_build assets/css/bootstrap.min.css assets/js/bootstrap.bundle.min.js assets/js/polyfill.min.js assets/js/tex-mml-chtml.js
	go build main.go

assets/css/bootstrap.min.css:
	curl https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css > assets/css/bootstrap.min.css

assets/js/bootstrap.bundle.min.js:
	curl https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js > assets/js/bootstrap.bundle.min.js

assets/js/polyfill.min.js:
	curl https://polyfill.io/v3/polyfill.min.js?features=es6 > assets/js/polyfill.min.js

assets/mathjax/tex-mml-chtml.js:
	npm install mathjax@3
	rm -rf assets/mathjax
	mv node_modules/mathjax/es5 assets/mathjax
	rm -rf node_modules

docker: build
	docker build -t stefanfritsch/go-markdown-renderer .
	docker push stefanfritsch/go-markdown-renderer
