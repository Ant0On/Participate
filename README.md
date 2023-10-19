# Participate

## Prerequisite commands
`yarn init`\
`yarn add -D @babel/cli @babel/core @babel/preset-env babel-loader webpack webpack-cli webpack-dev-server html-webpack-plugin vue-loader vue-template-compiler css-loader vue-style-loader sass-loader`\
`yarn add babel-polyfill vue  node-sass axios`
## To run project
Run `setup_project.sh` script in `scripts` directory

## To run without docker
1. Start main with `air -c .air.conf` (air will need to be installed) command to have it automatically refreshed (make sure you have correctly setup postgres database in .env file)
2. Run `npm run build` from root directory
3. Run `npm run start` from root directory

## To run with docker
1. Run `docker-compose build` from root directory
2. Run `docker-compose up` from root directory

## Plan
1.	Struktura i organizacja projektu (Grzegorz, Antek)
2.	Postawienie Dockera (Grzegorz, Antek)
3.	Autoryzacja/logowanie/rejestracja użytkownika. Podział na role. (Grzegorz, Antek)
4.	Po kolei przypadki użycia. (Grzegorz, Antek)
5.	Stopniowo front wprowadzać. (Grzegorz, Antek)

## TODO
1. Ogarniecie pakietów z vue js (https://stackoverflow.com/questions/64868632/vuejs-3-webpack-problem-with-vue-template-compiler)
2. Dodanie proda bez wyrzuconych portów i innych tego typu rzeczy
3. Ogarniecie zmiennych srodowiskowych uzywanych do polaczenia sie z baza
4. Setup bazy danych w modelach? Chyba powinno to byc w innym miejscu
5. cos sie jeszcze znajdzie...

## Przydatne linki
### User login/authentication
https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817
### Go/Gin/Vue integration
https://dev.to/mizutani/how-to-build-web-app-with-go-gin-gonic-vue-3987