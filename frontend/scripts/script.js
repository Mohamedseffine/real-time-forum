import { showAuthFormSignup, createBaseLayout } from "./render.js";
import {updateUserlist}from "./user.js"

// Initialize WebSocket connection
export let conn
 export function initializeWebSocket() {
    if (window["WebSocket"]) {
         conn = new WebSocket("ws://" + document.location.host + "/chat");
        conn.onopen = () => {
            conn.send(JSON.stringify({
                type:"message",
                message:"hi"
            }))
            console.log("WebSocket connection established");
        };
        
        conn.onmessage = (evt) => {
            console.log(evt);
            
            try {
                const data = JSON.parse(evt.data);
                console.log("Received data:", data);
                if (data.type === 'all_users'){
                    console.log(data.users);
                    updateUserlist(data.users)
                }
            } catch (err) {
                console.error("Error parsing WebSocket message:", err);
            }
        };
        
        conn.onerror = (err) => {
            console.error("WebSocket error:", err);
        };
        
        conn.onclose = () => {
            console.log("WebSocket connection closed");
        };
        
        return conn;
    }
    return null;
}
// Initialize the application
document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    if (token) {
        createBaseLayout();
    } else {
        showAuthFormSignup();
    }
});

// Utility function for date formatting
export function formatDateFromTimestamp(ms) {
    const date = new Date(ms);
    return date.toISOString();
}