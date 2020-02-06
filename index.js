"use strict";

// source: https://blog.myitcv.io/2020/02/04/portable-ci-cd-with-pure-go-github-actions.html

const { spawnSync} = require('child_process');


(async function() {
  const a = spawnSync("go", ["build", "-o", "piper", "."], {cwd: __dirname})
console.log('stdout ', a.stdout.toString());
console.log('stderr ', a.stderr.toString());
  const b = spawnSync(__dirname + "/piper", ["mavenExecute", "--goals", "verify"], {cwd: "/home/runner/work/piper-go-actions-playground/piper-go-actions-playground"})
console.log('stdout ', b.stdout.toString());
console.log('stderr ', b.stderr.toString());
})();
