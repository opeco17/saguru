<p align="center"><a href="https://gitnavi.dev/en" target="_blank" rel="noopener noreferrer"><img width="270" src="https://user-images.githubusercontent.com/46510874/144610986-29261b45-9b4a-4957-9973-54be60dc6c45.png" alt="gitnavi logo"></a></p>

## [gitnavi](https://gitnavi.dev/en)
gitnavi is Web application to help you search for GitHub issues and repositories with flexible queries and beautiful UI.

gitnavi support your open-source contribution.

## How to spin-up

### Create .env
First, it's necessary to create `.env` in the root directory of this project.

As this project access to GitHub API, GitHub user name and token need to be provided in `.env`.

Replace `YOUR_GITHUB_USERNAME` and `YOUR_GITHUB_API_TOKEN` with your GitHub user name and token respectively.

```sh
cat <<EOF > .env
GITHUB_API_USER=YOUR_GITHUB_USERNAME
GITHUB_API_TOKEN=YOUR_GITHUB_API_TOKEN
DB_USER=root
DB_PASSWORD=root
DB_HOST=database
MYSQL_ROOT_PASSWORD=root
EOF
```

### Run backend

Following command spin-up API server, Web server, and DB server.

```sh
docker-compose up -d --build
```

To update database, you can run the job with following command. It fetches GitHub issues and repositories through GitHub API.

```sh
docker-compose exec job bash -c "go run *.go all"
```

### Run frontend
Following commands spin-up Node.js server and serve frontend application.

```sh
cd frontend
npm run dev
```

You can check frontend page from http://localhost:3000.
