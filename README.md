<p align="center"><a href="https://saguru.dev" target="_blank" rel="noopener noreferrer"><img width="270" src="https://user-images.githubusercontent.com/46510874/144610986-29261b45-9b4a-4957-9973-54be60dc6c45.png" alt="saguru logo"></a></p>

## [saguru](https://saguru.dev)
saguru is Web application to help you search for GitHub issues and repositories with functional filters and beautiful UI.

saguru support your open-source contribution.

## How to spin-up

### Create .env
First, it's necessary to create `.env` in the root directory of this project.

As this project access to GitHub API, GitHub user name and token need to be provided in `.env`.

Replace `YOUR_GITHUB_USERNAME` and `YOUR_GITHUB_API_TOKEN` with your GitHub user name and token respectively.

```sh
cat <<EOF > .env
GITHUB_API_USER=YOUR_GITHUB_USERNAME
GITHUB_API_TOKEN=YOUR_GITHUB_API_TOKEN
MONGODB_HOST=mongo
MONGO_INITDB_ROOT_USERNAME=root
MONGO_INITDB_ROOT_PASSWORD=password
ME_CONFIG_MONGODB_ADMINUSERNAME=root
ME_CONFIG_MONGODB_ADMINPASSWORD=password
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
