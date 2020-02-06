"use strict";

const { spawn } = require('child_process');

(async function() {
    await spawn("go", ["build", "-o", "piper", "."], { cwd: __dirname, stdio: 'inherit' })
    await spawn(__dirname + "/piper", ["mavenExecute", "--goals", "verify"], { stdio: 'inherit', cwd: "/home/runner/work/piper-go-actions-playground/piper-go-actions-playground" })
})();
