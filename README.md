# Participate

## Prerequisite commands
`yarn init`
`yarn add -D @babel/cli @babel/core @babel/preset-env babel-loader webpack webpack-cli webpack-dev-server html-webpack-plugin vue-loader vue-template-compiler css-loader vue-style-loader sass-loader`
`yarn add babel-polyfill vue  node-sass axios`

## To run
1. Start main with `air -c .air.conf` command to have it automatically refreshed (make sure you have correctly setup postgres database in .env file)
2. Run `npm run build` from root directory
3. Run `npm run start` from root directory

## Plan
1.	Struktura i organizacja projektu (Grzegorz, Antek)
2.	Postawienie Dockera (Grzegorz, Antek)
3.	Autoryzacja/logowanie/rejestracja użytkownika. Podział na role. (Grzegorz, Antek)
4.	Po kolei przypadki użycia. (Grzegorz, Antek)
5.	Stopniowo front wprowadzać. (Grzegorz, Antek)

## Przydatne linki
### User login/authentication
https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817
### Go/Gin/Vue integration
https://dev.to/mizutani/how-to-build-web-app-with-go-gin-gonic-vue-3987