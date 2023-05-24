buildserver:
	cd ./server/cmd/server/ && go build

runserver: buildserver
	sudo ./server/cmd/server/server

runui: 
	cd ./client && npm run dev


runuinetwork: 
	cd ./client && npm run devnetwork

plantuml: 
	goplantuml -aggregate-private-members -show-aggregations -show-connection-labels -recursive server/internal/ > serverClassDiagram.puml
	plantuml serverClassDiagram.puml 

generatecoveragereport:
	cd ./server && go test ./... -coverprofile cover.out