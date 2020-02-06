"use strict";

// source: https://blog.myitcv.io/2020/02/04/portable-ci-cd-with-pure-go-github-actions.html

const { spawnSync} = require('child_process');


(async function() {
  const a = spawnSync("go", ["build", "-o", "piper", "."], {cwd: __dirname})
  console.log('error', a.error);
console.log('stdout ', a.stdout);
console.log('stderr ', a.stderr);
  const b = spawnSync(__dirname + "/piper", ["mavenExecute", "--goals", "verify"], {cwd: "/home/runner/work/piper-go-actions-playground/piper-go-actions-playground"})
  console.log('error', b.error);
console.log('stdout ', b.stdout);
console.log('stderr ', b.stderr);
})();
