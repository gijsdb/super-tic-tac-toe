buildserver:
	cd ./server/cmd/server/ && go build

runserver: buildserver
	sudo ./server/cmd/server/server


runui: 
	cd ./client && npm run dev


runuinetwork: 
	cd ./client && npm run devnetwork

plantuml: 
	goplantuml -aggregate-private-members -show-aggregations -show-connection-labels -recursive server/internal/ > classDiagram.puml
	plantuml classDiagram.puml 