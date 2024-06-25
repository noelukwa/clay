// Simplified adaptation of wasm_exec.js for Cloudflare Workers

import { Go } from "./exec.js";



addEventListener('fetch', (event) => {
    event.respondWith(_do(event.request));
});


// function _fetchhandle(req) {
//     return new Promise((async (resolve, reject) => {
//         try {
//         //     const go = new Go();
//         //     const instance = await WebAssembly.instantiate(WORKER_WASM, go.importObject);
//         //     go.run(instance);
//         //     wrapfunc(req).then(async (rr)=>{
//         //     const res_1 = await gohandle(rr)
//         //     return resolve(res_1);
//         //    })
//         const go = new Go();
//         const instance = new WebAssembly.Instance(WORKER_WASM, go.importObject)
//         go._inst = instance
//         console.log(instance.exports)
//         r
//         } catch (e) {
//             console.log(e);
//             reject(new Response(e.message, { status: 500 }));
//         }
//     }));
// }


async function _do(req) {
    console.log(Go)
    return new Promise((async (resolve, reject) => { 
       try {
            const go = new Go();
            const instance = await WebAssembly.instantiate(WORKER_WASM, go.importObject);
            go.run(instance); 
            console.log(instance.instance)
            resolve(new Response("YO", { status: 200 }));
       } catch (error) {
            console.log(error)
            reject(new Response(e.message, { status: 500 }));
       }
    }))
}

