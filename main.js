const go = new Go();

let mod, inst;

// Loads Go compiled wasm file

WebAssembly.instantiateStreaming(fetch("go.wasm"), go.importObject).then(async (result) => {
  mod = result.module;
  inst = result.instance;
  await go.run(inst);
});

// Loads C compiled wasm file

fetch('main.wasm').then(response =>
  response.arrayBuffer()
).then(bytes => WebAssembly.instantiate(bytes)).then(results => {
  instance = results.instance;
  document.getElementById("container").textContent = instance.exports.main();
}).catch(console.error);
