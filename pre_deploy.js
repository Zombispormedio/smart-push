/// CMD PROD

var fs=require("fs");

const DEV_MAIN="main.go";
const PROD_MAIN="cmd/smart-push/main.go";

//GO Deps
const GODEP="Godeps/Godeps.json";
const INTERNAL="github.com/Zombispormedio/smart-push/";
const PACKAGES=[
    "config", "router", "controllers", "lib/response", "middleware", "lib/request", "lib/store", "lib/redis", "lib/utils",
     "lib/rabbit", "lib/mosquito"
];

fs.writeFileSync(PROD_MAIN, fs.readFileSync(DEV_MAIN));



var pre_godep=JSON.parse(fs.readFileSync(GODEP));

pre_godep.Deps=pre_godep.Deps.concat(
    PACKAGES.map(function(pkg){return {"ImportPath":INTERNAL+pkg};})
);

fs.writeFileSync(GODEP, JSON.stringify(pre_godep));