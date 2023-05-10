// // import * as rust from "../src-rust/pkg/src_rust.js"
//   

// import init, * as rust from "../src-rust/pkg/src_rust.js"

// // importScripts("../src-rust/pkg/src_rust.js")
// console.log("Foo and BAR")

import init, * as rust from "../src-rust/pkg/src_rust.js"

console.log('foo')

init().then(() => {
  rust.MagicSquare.run()
})

