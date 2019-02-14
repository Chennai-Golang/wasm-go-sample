const go = new Go();

let mod, inst;

WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject).then(async (result) => {
  mod = result.module;
  inst = result.instance;
  await go.run(inst);
});

fetch('main.wasm').then(response =>
  response.arrayBuffer()
).then(bytes => WebAssembly.instantiate(bytes)).then(results => {
  instance = results.instance;
  document.getElementById("container").textContent = instance.exports.main();
}).catch(console.error);
