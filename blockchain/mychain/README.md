Built using "ignite scaffold list name:string description:string"

# 1. run the Dockerfile

docker build -t mychain .

docker run -p 26657:26657 -p 9090:9090 -p 1317:1317 mychain

# 2. send CRUD requests

ports: 26657 (tendermint rpc), 9090 (gRPC), 1317 (REST)

Create -> mychaind tx resource create "My Resource" "Very Cool Resource" --from mywallet --keyring-backend test

Read -> 
    All ->  curl "http://localhost:26657/abci_query?path=/store/resource/key"
            grpcurl -plaintext localhost:9090 mychain.resource.Query/AllResources
            curl http://localhost:1317/mychain/resource

    Specific -> curl "http://localhost:26657/abci_query?path=/store/resource/key&data=0x31"
                grpcurl -plaintext -d '{"id":"1"}' localhost:9090 mychain.resource.Query/GetResource
                curl http://localhost:1317/mychain/resource/1

Update -> mychaind tx resource update 1 "Updated Name" "Updated Description" --from mywallet --keyring-backend test

Delete -> mychaind tx resource delete 1 --from mywallet --keyring-backend test