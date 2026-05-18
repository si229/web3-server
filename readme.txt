临时代币地址：0x8464135c8F25Da09e49BC8782676a84730C318bC
逻辑合约地址：0x71C95911E9a5D330f4D621842EC243EE1343292e

安装abigen:
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

签名验证：
https://etherscan.io/verifiedSignatures#

生成命令：
abigen -abi internal\contract\baccarat.abi --pkg baccarat --out internal\contract\baccarat.go
abigen -abi internal\contract\erc20.usdt.abi --pkg erc20usdt --out internal\contract\erc20usdt.go
