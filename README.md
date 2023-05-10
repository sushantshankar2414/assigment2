Requirements
LOCAL SETUP
go >= 1.17.7
docker >= 20.10.7
git >= 2.25.1

LOCAL SETUP
Clone git repository: git clone https://github.com/sushantshankar2414/assigment2
COPY .env
RUN FOLLOWING COMMANDS
go get 			# to get all dependencies defined in the project
go build		# to build the project
docker-compose up --build #to make docker image

