buildserver:
	cd ./server/cmd/server/ && go build --tags json1

runserver: buildserver
	sudo ./server/cmd/server/server


runui: 
	cd ./client && npm run dev


runuinetwork: 
	cd ./client && npm run devnetwork


