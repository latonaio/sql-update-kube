docker-build:
	echo "build for development"
	bash docker-build.sh

add-submodule:
	git submodule add git@github.com:latonaio/sqlboiler.git sqlboiler

# Foreign keyに関するエラーを修正したリポジトリ
clone-submodule:
	git clone git@bitbucket.org:latonaio/sqlboiler.git sqlboiler --recurse-submodules

sqlboiler-build:
	cd sqlboiler && go build -o ../database/sqlboiler main.go
