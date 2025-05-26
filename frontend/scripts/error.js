import { Init } from "./script.js";

export function RenderError(Str , status, message) {
    const body = document.body
    body.innerHTML = ""
    body.innerHTML= `<div class="error-container">
    <h1 id="err-status">${status}</h1>
    <h2 id="err-head">${Str}</h2>
    <p id="err-text">${message}</p>
    <button id="back-butt" type="button">Back</button>
    </div>
    `
    document.getElementById("back-butt").addEventListener('click', ()=>{
        body.innerHTML=`
        <div id="root"></div>
        `
        history.replaceState(null, null, "/")
        Init()
    });
}