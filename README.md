# vercel-faas
A set of serverless apis written in go(gin) and deployed in [vercel](https://vercel.com) 

all apis: [https://faas.kirito41dd.cn/api/1](https://faas.kirito41dd.cn/api/1)

## local run
```bash
# install cli
npm i vercel
# run
vercel dev
```
## notice
* `go.mod` must be placed in the root directory to take effect 
* `/api` just entry, can't declare modules in subdirectories 
* `/handle` stores your code
* `vercel.json` [document](https://vercel.com/docs/configuration)